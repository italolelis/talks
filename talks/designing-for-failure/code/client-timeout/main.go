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

	if err := requestSimpleTimeout(); err != nil {
		log.Println(err)
	}

	fmt.Println("requests finished")
}

// START SIMPLE_TIMEOUT OMIT
func requestSimpleTimeout() error {
	c := &http.Client{
		Timeout: 2 * time.Second,
	}

	r, err := http.NewRequestWithContext(
		context.Background(),
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

// END SIMPLE_TIMEOUT OMIT
