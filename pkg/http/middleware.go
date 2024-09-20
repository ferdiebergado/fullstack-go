package http

import (
	"log"
	"net/http"
)

// LoggingMiddleware logs each request.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s from %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

// ResponseMiddleware logs each response.
func ResponseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sw := &statusWriter{ResponseWriter: w}
		next.ServeHTTP(sw, r)
		status:=sw.status
		log.Printf("Response: %d %s\n", status, http.StatusText(status))
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