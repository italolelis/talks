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

	if err := requestSimpleTimeout(); err != nil {
		log.Println(err)
	}

	if err := requestCompleteTimeout(); err != nil {
		log.Println(err)
	}

	if err := requestWithContext(); err != nil {
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

func requestCompleteTimeout() error {
	// START COMPLETE_TIMEOUT OMIT
	c := &http.Client{
		Timeout: 2 * time.Second,
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
	// END COMPLETE_TIMEOUT OMIT

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

// START CONTEXT_TIMEOUT OMIT
func requestWithContext() error {
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
