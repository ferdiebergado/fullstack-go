package main

import (
	"errors"
	"net/http"

	"github.com/ferdiebergado/fullstack-go/internal/config"
	"github.com/ferdiebergado/fullstack-go/internal/db"
	"github.com/ferdiebergado/fullstack-go/internal/domain/activity"
	"github.com/ferdiebergado/fullstack-go/internal/domain/division"
	"github.com/ferdiebergado/fullstack-go/internal/domain/host"
	"github.com/ferdiebergado/fullstack-go/internal/domain/venue"
	"github.com/ferdiebergado/fullstack-go/internal/ui"
	"github.com/ferdiebergado/fullstack-go/pkg/http/response"
	router "github.com/ferdiebergado/go-express"
	"github.com/ferdiebergado/go-express/middleware"
)

func NewApp(database *db.Database) *router.Router {

	// Host Handler
	hostService := host.NewHostService(database)
	hostHandler := host.NewHostHandler(hostService)

	// Create the Venue Handler.
	venueService := venue.NewVenueService(database)
	venueHandler := venue.NewVenueHandler(venueService)

	// Division Service
	divisionService := division.NewDivisionService(database)

	// Create the Activity Handler.
	activityService := activity.NewActivityService(database)
	activityHandler := activity.NewActivityHandler(activityService, venueService, hostService, divisionService)

	// Create the router.
	router := router.NewRouter()

	// Register global middlewares.
	router.Use(middleware.RequestLogger)
	router.Use(middleware.StripTrailingSlashes)
	router.Use(middleware.PanicRecovery)

	// Serve static files.
	router.ServeStatic(config.StaticDir)

	// Activities routes.
	activity.AddRoutes(router, *activityHandler)

	// Venue routes.
	venue.AddRoutes(router, *venueHandler)

	// Hosts routes.
	host.AddRoutes(router, *hostHandler)

	// Home page
	router.Get("/{$}", func(w http.ResponseWriter, r *http.Request) {
		err := ui.RenderHTML(w, "index.html", nil)

		if err != nil {
			response.ErrorHandler(w, r, http.StatusBadRequest, "unable to render template", err)
			return
		}
	})

	// Not found handler
	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		status := http.StatusNotFound
		response.ErrorHandler(w, r, status, http.StatusText(status), errors.New("page not found"))
	})

	return router
}
