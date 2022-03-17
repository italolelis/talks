package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	fmt.Println("sending requests")

	if err := requestCompleteTimeout(); err != nil {
		log.Println(err)
	}

	fmt.Println("requests finished")
}

func requestCompleteTimeout() error {
	// START COMPLETE_TIMEOUT OMIT
	c := &http.Client{
		Timeout: 3 * time.Second,
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	r, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"https://httpstat.us/200?sleep=2000",
		nil,
	)
	// ...
	// END COMPLETE_TIMEOUT OMIT
	if err != nil {
		return err
	}

	if _, err := c.Do(r); err != nil {
		return err
	}

	return nil
}
