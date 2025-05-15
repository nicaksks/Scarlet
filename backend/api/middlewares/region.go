package middlewares

import (
	"encoding/json"
	"net/http"
	"net/url"
	"scarlet/api/errors"
	"slices"
	"strings"
)

var REGIONS = []string{"NA", "EU", "ASIA"}

func response(w http.ResponseWriter, statusCode int, message errors.Message) {
	serialize, _ := json.Marshal(errors.ResponseError(statusCode, message))
	w.WriteHeader(statusCode)
	w.Write(serialize)
}

func updateQuery(r *http.Request, query string) url.Values {
	q := r.URL.Query()
	q.Set("r", query)
	return q
}

func ValidateRegion(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		query := strings.ToUpper(r.URL.Query().Get("r"))

		if !r.URL.Query().Has("r") {
			r.URL.RawQuery = updateQuery(r, REGIONS[0]).Encode()
			next.ServeHTTP(w, r)
			return
		}

		if !slices.Contains(REGIONS, query) {
			response(w, http.StatusForbidden, errors.UNSUPPORTED_REGION)
			return
		}

		r.URL.RawQuery = updateQuery(r, query).Encode()

		next.ServeHTTP(w, r)
	})
}
