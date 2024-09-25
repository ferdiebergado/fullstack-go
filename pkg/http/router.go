package http

import (
	"net/http"
)

// Router is a custom HTTP router built on top of ServeMux with middleware support.
type Router struct {
	mux         *http.ServeMux
	middlewares []Middleware
}

// Middleware defines the signature for middleware functions.
type Middleware func(http.Handler) http.Handler

// NewRouter creates a new Router.
func NewRouter() *Router {
	return &Router{
		mux:         http.NewServeMux(),
		middlewares: []Middleware{},
	}
}

// Use adds a middleware to the router.
func (r *Router) Use(mw Middleware) {
	r.middlewares = append(r.middlewares, mw)
}

// Handle registers a new route with the router.
func (r *Router) Handle(pattern string, handler http.Handler) {
	// Wrap the handler with all the middlewares.
	finalHandler := handler
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		finalHandler = r.middlewares[i](finalHandler)
	}
	r.mux.Handle(pattern, finalHandler)
}

// ServeHTTP allows Router to satisfy http.Handler.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}

// Custom function to handle specific errors like 400, 403.
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int) {
	switch status {
	case http.StatusBadRequest:
		http.Error(w, "Bad Request", http.StatusBadRequest)
	case http.StatusForbidden:
		http.Error(w, "Forbidden", http.StatusForbidden)
	case http.StatusNotFound:
		http.Error(w, "Not Found", http.StatusNotFound)
	case http.StatusInternalServerError:
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	default:
		http.Error(w, "An Error Occurred", status)
	}
}
