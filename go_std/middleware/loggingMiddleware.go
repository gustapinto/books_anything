package middleware

import (
	"log"
	"net/http"
)

func Logging(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			logger.Println(r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
		}()

		next.ServeHTTP(w, r)
	})
}
