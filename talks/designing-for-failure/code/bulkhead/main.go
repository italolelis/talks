package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	// START BULKHEAD OMIT
	err = hystrix.Do("titanic", func() error {
		resp, err = c.Do(request)
		if err != nil {
			return err
		}

		if resp.StatusCode >= http.StatusInternalServerError {
			return errors.New("failed to block the water. Water is coming in!!")
		}
		return nil
	}, func(err error) error {
		return fmt.Errorf("This ship is going to sink because of: %w", err)
	})
	// END BULKHEAD OMIT
}
