package main

import (
	"database/sql"
	"net/http"

	"github.com/ferdiebergado/fullstack-go/db"
	"github.com/ferdiebergado/fullstack-go/internal/activity/handlers"
	myhttp "github.com/ferdiebergado/fullstack-go/pkg/http"
	"github.com/ferdiebergado/fullstack-go/view"
)

func NewApp(conn *sql.DB) *myhttp.Router {

	queries := db.New(conn)

	activityHandler := &handlers.ActivityHandler{Queries: queries}

	// Create the router.
	router := myhttp.NewRouter()

	// Use logging and error handling middleware.
	router.Use(myhttp.LoggingMiddleware)
	router.Use(myhttp.ErrorHandlerMiddleware)

	// Register routes.
	router.Handle("GET /activities", http.HandlerFunc(activityHandler.ActivityIndex))
	router.Handle("GET /activities/create", http.HandlerFunc(activityHandler.CreateActivity))
	router.Handle("GET /activities/{id}", http.HandlerFunc(activityHandler.ViewActivity))
	router.Handle("GET /activities/{id}/edit", http.HandlerFunc(activityHandler.EditActivity))
	router.Handle("POST /activities/{id}/edit", http.HandlerFunc(activityHandler.UpdateActivity))
	router.Handle("POST /activities", http.HandlerFunc(activityHandler.SaveActivity))

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
			http.NotFound(w, r)
			return
		}

		view.RenderTemplate(w, "index.html", nil)
	}))

	router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	return router
}
