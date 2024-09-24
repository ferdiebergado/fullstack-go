package main

import (
	"net/http"
	"path/filepath"

	myhttp "github.com/ferdiebergado/fullstack-go/pkg/http"
	"github.com/ferdiebergado/fullstack-go/templates"
)

func NewApp() *myhttp.Router {
	// Create the router.
	router := myhttp.NewRouter()

	// Use logging and error handling middleware.
	router.Use(myhttp.LoggingMiddleware)
	router.Use(myhttp.ErrorHandlerMiddleware)

	// Register routes.
	router.Handle("/activities", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		templates.Render(w, "activities/index.html")
	}))

	// Register a route that triggers a 403 Forbidden error.
	router.Handle("/forbidden", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		myhttp.ErrorHandler(w, r, http.StatusForbidden)
	}))

	// Register a route that triggers a 400 Bad Request error.
	router.Handle("/badrequest", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		myhttp.ErrorHandler(w, r, http.StatusBadRequest)
	}))

	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			myhttp.ErrorHandler(w, r, http.StatusNotFound)
			return
		}

		templates.Render(w, "index.html")
	}))

	router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	return router
}
