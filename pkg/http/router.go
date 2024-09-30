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

// Handle registers a new route with the router and applies route-specific middleware.
func (r *Router) Handle(pattern string, handler http.Handler, middlewares ...Middleware) {
	// Wrap the handler with route-specific middlewares first.
	finalHandler := handler
	for i := len(middlewares) - 1; i >= 0; i-- {
		finalHandler = middlewares[i](finalHandler)
	}

	// Wrap with global middlewares.
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		finalHandler = r.middlewares[i](finalHandler)
	}

	r.mux.Handle(pattern, finalHandler)
}

// ServeHTTP allows Router to satisfy http.Handler.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.mux.ServeHTTP(w, req)
}
