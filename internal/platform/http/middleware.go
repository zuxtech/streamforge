package http

import (
	"net/http"
	"time"

	"github.com/google/uuid"
)

const poweredBy = "StreamForge"

func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")

		if requestID == "" {
			requestID = "req_" + uuid.NewString()
		}

		w.Header().Set("X-Request-ID", requestID)

		next.ServeHTTP(w, r)
	})
}

func poweredByMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Powered-By", poweredBy)

		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start)

		// Add structured logging later.
		// Example:
		// method, path, request ID, duration, status, etc.
		_ = duration
	})
}