package middlewares

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()

			next.ServeHTTP(w, r)

			log.Printf(
					"%s %s %s %s",
					r.Method,
					r.RequestURI,
					r.RemoteAddr,
					time.Since(startTime),
			)
	})
}