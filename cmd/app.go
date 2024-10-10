package main

import (
	"errors"
	"net/http"

	"github.com/ferdiebergado/fullstack-go/internal/db"
	"github.com/ferdiebergado/fullstack-go/internal/domain/activity"
	"github.com/ferdiebergado/fullstack-go/internal/domain/host"
	"github.com/ferdiebergado/fullstack-go/internal/domain/venue"
	"github.com/ferdiebergado/fullstack-go/internal/ui"
	myhttp "github.com/ferdiebergado/fullstack-go/pkg/http"
)

func NewApp(database *db.Database) *myhttp.Router {

	// Host Handler
	hostService := host.NewHostService(database)
	hostHandler := host.NewHostHandler(hostService)

	// Create the Venue Handler.
	venueService := venue.NewVenueService(database)
	venueHandler := venue.NewVenueHandler(venueService)

	// Create the Activity Handler.
	activityService := activity.NewActivityService(database)
	activityHandler := activity.NewActivityHandler(activityService, venueService, hostService)

	// Create the router.
	router := myhttp.NewRouter()

	// Register global middlewares.
	router.Use(myhttp.LoggingMiddleware)
	router.Use(myhttp.StripTrailingSlash)
	router.Use(myhttp.ErrorRecoveryMiddleware)

	// Serve static files.
	router.Handle("GET /assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Activities routes.
	activity.AddRoutes(router, *activityHandler)

	// Venues routes.
	router.Handle("POST /api/venues", http.HandlerFunc(venueHandler.SaveVenue))

	// Hosts routes.
	router.Handle("POST /api/hosts", http.HandlerFunc(hostHandler.SaveHost))

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
