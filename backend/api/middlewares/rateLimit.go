package middlewares

import (
	"net/http"
	"os"
	"scarlet/api/errors"

	"golang.org/x/time/rate"
)

const (
	LIMIT  = 10
	HEADER = "X-Scarlet-Session"
)

// This rate limit is temporary
// TODO: Rate Limit based in Token;
func RateLimit(next http.Handler) http.Handler {

	limit := rate.NewLimiter(LIMIT, 1)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get(HEADER) == os.Getenv("TOKEN") {
			next.ServeHTTP(w, r)
			return
		}

		if !limit.Allow() {
			response(w, http.StatusTooManyRequests, errors.TOO_MANY_REQUESTS)
			return
		}

		next.ServeHTTP(w, r)
	})
}
