package middleware

import (
	"log"
	"net/http"
	"time"
)

// LoggerMiddleware logs incoming requests
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		
		// Serve the request
		next.ServeHTTP(w, r)
		
		// Log the request details
		log.Printf(
			"%s %s %s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}