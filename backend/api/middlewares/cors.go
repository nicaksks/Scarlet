package middlewares

import "net/http"

func CorsAndCache(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := w.Header()
		header.Add("Access-Control-Allow-Origin", "*")
		header.Add("Access-Control-Request-Method", "GET")
		header.Add("Access-Control-Request-Headers", "Content-Type")
		header.Add("Access-Control-Max-Ages", "1800") // 30 minutes in cache
		header.Add("Content-Type", "application/json; charset=utf-8")
		header.Add("Cache-Control", "public, max-age=1800, must-revalidate") // 30 minutes in cache
		next.ServeHTTP(w, r)
	})
}
