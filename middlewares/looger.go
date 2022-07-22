package middlewares

import (
	"log"
	"net/http"
)

// Logging use for log a http request
func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%v] : %v\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
