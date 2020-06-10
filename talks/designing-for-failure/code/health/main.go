package main

import (
	"net/http"
	"time"

	"github.com/heptiolabs/healthcheck"
)

// START HEALTH OMIT
func main() {
	// ...
	h := healthcheck.NewHandler()
	h.AddLivenessCheck("goroutine-threshold", healthcheck.GoroutineCountCheck(100))

	h.AddReadinessCheck("database", healthcheck.DatabasePingCheck(db, 1*time.Second))
	h.AddReadinessCheck("payments-srv", healthcheck.HTTPGetCheck("https://example.io", 1*time.Second))
	h.AddReadinessCheck("google", healthcheck.HTTPGetCheck("https://www.google.com", 1*time.Second))

	http.ListenAndServe(":9090", h)
}

// END HEALTH OMIT
