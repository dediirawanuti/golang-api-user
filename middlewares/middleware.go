package middlewares

import (
	"github.com/golang-api-user/auth"
	"github.com/golang-api-user/user/common"
	"net/http"
	"strings"
)

// SetMiddlewareJSON use for output of the json format
func SetMiddlewareJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var urls = r.URL.Path
		var escapeExtension = []string{".jpeg", ".jpg", ".png"}

		for _, v := range escapeExtension {
			if !strings.Contains(urls, v) {
				w.Header().Set("Content-Type", "application/json")
			}
		}
		next.ServeHTTP(w, r)
	})
}

// SetMiddlewareAuthentication use for verify token when user access the api
func SetMiddlewareAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := auth.TokenValid(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			common.Response(w, common.Message(false, "Unauthorized", nil))
			return
		}
		next.ServeHTTP(w, r)
	})
}

// FileSizeLimiter use for is used to limit file size when uploaded
func FileSizeLimiter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var maxUploadSize int64 = 5 * 1024 * 1024 // limiter 2mb / file
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			common.Response(w, common.Message(false, "Request failed, because file too big to upload", nil))
			return
		}
		next.ServeHTTP(w, r)
	})
}
