package http

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware logs each request.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := &statusWriter{ResponseWriter: w}
		next.ServeHTTP(sw, r)
		duration := time.Since(start)
		statusCode := sw.status
		log.Printf("%s %s %s %d %s %s", r.Method, r.URL.Path, r.Proto, statusCode, http.StatusText(statusCode), duration)
	})
}

// ErrorHandlerMiddleware catches and handles errors.
func ErrorHandlerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Internal error: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		// Catch 404 errors by checking if no handler is found.
		originalWriter := w
		w = &statusWriter{ResponseWriter: originalWriter}
		next.ServeHTTP(w, r)

		if w.(*statusWriter).status == 0 {
			http.Error(originalWriter, "Not Found", http.StatusNotFound)
		}
	})
}
