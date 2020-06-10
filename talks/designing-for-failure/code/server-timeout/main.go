package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"time"
)

// START SERVER_TIMEOUT OMIT
func main() {
	h := timeoutHandler{}

	srv := &http.Server{
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       10 * time.Second,
		ReadHeaderTimeout: 20 * time.Second,
		Handler:           h,
	}

	log.Println("serving requests")
	log.Println(srv.ListenAndServe())
}

// END SERVER_TIMEOUT OMIT

type timeoutHandler struct{}

func (h timeoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	timer := time.AfterFunc(5*time.Second, func() {
		r.Body.Close()
	})

	bodyBytes := make([]byte, 0)
	for {
		timer.Reset(1 * time.Second)

		_, err := io.CopyN(bytes.NewBuffer(bodyBytes), r.Body, 256)
		if err != nil {
			break
		}
	}
}
