package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/sony/gobreaker"
)

func main() {
	// START CB_STRATEGY OMIT
	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name: "HTTP GET Example",
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 3 && failureRatio >= 0.6
		},
	})
	// END CB_STRATEGY OMIT

	// START CB_REQ OMIT
	_, err := cb.Execute(func() (interface{}, error) {
		resp, err := http.Get("https://google.com")
		if err != nil {
			return nil, err
		}

		if resp.StatusCode >= http.StatusBadRequest {
			return nil, errors.New("request failed")
		}

		return resp, nil
	})
	// END CB_REQ OMIT

	if err != nil {
		log.Print(err)
	}
}
