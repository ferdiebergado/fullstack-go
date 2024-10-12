package http

import (
	"log"
	"net/http"
	"runtime/debug"
	"strings"
	"time"
)

// RequestLogger logs each request.
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := &statusWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(sw, r)
		duration := time.Since(start)
		statusCode := sw.status
		log.Printf("%s %s %s %d %s %s", r.Method, r.URL.Path, r.Proto, statusCode, http.StatusText(statusCode), duration)
	})
}

// StripTrailingSlashes is a middleware that removes the trailing slash from the URL path.
func StripTrailingSlashes(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" && strings.HasSuffix(r.URL.Path, "/") {
			// Remove the trailing slash and redirect to the new URL.
			log.Println("Removing trailing slash and redirecting...")
			http.Redirect(w, r, strings.TrimSuffix(r.URL.Path, "/"), http.StatusMovedPermanently)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// PanicRecovery catches and handles errors.
func PanicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Internal error: %v", err)
				log.Println(string(debug.Stack()))
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
