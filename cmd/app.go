package main

import (
	"errors"
	"net/http"

	"github.com/ferdiebergado/fullstack-go/internal/db"
	"github.com/ferdiebergado/fullstack-go/internal/domain/activity"
	"github.com/ferdiebergado/fullstack-go/internal/ui"
	myhttp "github.com/ferdiebergado/fullstack-go/pkg/http"
)

func NewApp(database *db.Database) *myhttp.Router {

	// Create the Activity Handler.
	activityService := activity.NewActivityService(database)
	activityHandler := activity.NewActivityHandler(activityService)

	// Create the router.
	router := myhttp.NewRouter()

	// Register global middlewares.
	router.Use(myhttp.LoggingMiddleware)
	router.Use(myhttp.StripTrailingSlash)
	router.Use(myhttp.ErrorRecoveryMiddleware)

	// Serve static files.
	router.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Register activities routes.
	router.Handle("GET /activities", http.HandlerFunc(activityHandler.ListActiveActivities))
	router.Handle("GET /activities/create", http.HandlerFunc(activityHandler.CreateActivity))
	router.Handle("GET /activities/{id}", http.HandlerFunc(activityHandler.ViewActivity))
	router.Handle("GET /activities/{id}/edit", http.HandlerFunc(activityHandler.EditActivity))
	router.Handle("GET /api/activities/{id}", http.HandlerFunc(activityHandler.GetActivity))
	router.Handle("POST /api/activities", http.HandlerFunc(activityHandler.SaveActivity))
	router.Handle("PUT /api/activities/{id}", http.HandlerFunc(activityHandler.UpdateActivity))
	router.Handle("DELETE /api/activities/{id}", http.HandlerFunc(activityHandler.DeleteActivity))

	// Register venues routes.
	router.Handle("POST /api/venues", http.HandlerFunc(activityHandler.SaveVenue))

	// Home page
	router.Handle("GET /{$}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := ui.RenderTemplate(w, "index.html", nil)

		if err != nil {
			myhttp.ErrorHandler(w, r, http.StatusBadRequest, "unable to render template", err)
			return
		}
	}))

	// Not found handler
	router.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		status := http.StatusNotFound
		myhttp.ErrorHandler(w, r, status, http.StatusText(status), errors.New("page not found"))
	}))

	return router
}
