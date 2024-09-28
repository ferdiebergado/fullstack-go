package http

import (
	"log"
	"net/http"
	"runtime/debug"
	"strings"
	"time"
)

// LoggingMiddleware logs each request.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := &statusWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(sw, r)
		duration := time.Since(start)
		statusCode := sw.status
		log.Printf("%s %s %s %d %s %s", r.Method, r.URL.Path, r.Proto, statusCode, http.StatusText(statusCode), duration)
	})
}

// StripTrailingSlash is a middleware that removes the trailing slash from the URL path.
func StripTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" && strings.HasSuffix(r.URL.Path, "/") {
			// Remove the trailing slash and redirect to the new URL.
			http.Redirect(w, r, strings.TrimSuffix(r.URL.Path, "/"), http.StatusMovedPermanently)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// ErrorHandlerMiddleware catches and handles errors.
func ErrorHandlerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Internal error: %v", err)
				log.Println(debug.Stack())
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		// Catch 404 errors by checking if no handler is found.
		originalWriter := w
		w = &statusWriter{ResponseWriter: originalWriter, status: http.StatusOK}
		next.ServeHTTP(w, r)

		if w.(*statusWriter).status == 0 {
			http.Error(originalWriter, "Not Found", http.StatusNotFound)
		}
	})
}
