package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("sending requests")

	if err := requestWithContextFails(); err != nil {
		log.Println(err)
	}

	fmt.Println("requests finished")
}

// START CONTEXT_TIMEOUT OMIT
func requestWithContextFails() error {
	c := &http.Client{}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	r, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://httpstat.us/200?sleep=3000",
		nil,
	)
	if err != nil {
		return err
	}

	if _, err := c.Do(r); err != nil {
		return err
	}

	return nil
}

// END CONTEXT_TIMEOUT OMIT
