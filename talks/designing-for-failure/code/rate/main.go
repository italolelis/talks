package main

import (
	"log"
	"net/http"

	libredis "github.com/go-redis/redis/v7"
	limiter "github.com/ulule/limiter/v3"
	mhttp "github.com/ulule/limiter/v3/drivers/middleware/stdlib"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

func main() {
	// START LIMITER OMIT
	// Define a limit rate to 2 requests per minute.
	rate, err := limiter.NewRateFromFormatted("2-M")
	if err != nil {
		log.Fatal(err)
	}

	limiter := limiter.New(memory.NewStore(), rate)
	// END LIMITER OMIT

	// START MIDDLEWARE OMIT
	// Create a new middleware with the limiter instance.
	middleware := mhttp.NewMiddleware(limiter)

	// Launch a simple server.
	http.Handle("/", middleware.Handler(http.HandlerFunc(index)))
	log.Fatal(http.ListenAndServe(":8080", nil))
	// END MIDDLEWARE OMIT

}

// START REDIS_LIMITER OMIT
func redisLimiter(rate limiter.Rate) (limiter.Limiter, error) {
	// Create a redis client.
	option, err := libredis.ParseURL("redis://localhost:6379/0")
	if err != nil {
		return nil, err
	}
	client := libredis.NewClient(option)

	// Create a store with the redis client.
	store, err := sredis.NewStoreWithOptions(client, limiter.StoreOptions{
		Prefix:   "limiter_http_example",
		MaxRetry: 3,
	})
	if err != nil {
		return nil, err
	}

	return limiter.New(store, rate), nil
}

// END REDIS_LIMITER OMIT

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	_, err := w.Write([]byte(`{"message": "ok"}`))
	if err != nil {
		log.Fatal(err)
	}
}
