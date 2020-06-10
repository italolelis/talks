package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/avast/retry-go"
)

func main() {
	url := "http://example.com"
	var body []byte

	// START RETRY OMIT
	err := retry.Do(
		func() error {
			resp, err := http.Get(url)
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			body, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				return err
			}

			return nil
		},
		retry.Attempts(3),
		retry.DelayType(retry.DefaultDelayType),
	)
	// END RETRY OMIT
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(body)
}
