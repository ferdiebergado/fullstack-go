package main

import (
	"net/http"

	myhttp "github.com/ferdiebergado/fullstack-go/pkg/http"
)

func NewApp() *myhttp.Router {
	// Create the router.
	router := myhttp.NewRouter()

	// Use logging and error handling middleware.
	router.Use(myhttp.LoggingMiddleware)
	router.Use(myhttp.ErrorHandlerMiddleware)

	// Register routes.
	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	}))

	// Register a route that triggers a 403 Forbidden error.
	router.Handle("/forbidden", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		myhttp.ErrorHandler(w, r, http.StatusForbidden)
	}))

	// Register a route that triggers a 400 Bad Request error.
	router.Handle("/badrequest", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		myhttp.ErrorHandler(w, r, http.StatusBadRequest)
	}))

	return router
}
