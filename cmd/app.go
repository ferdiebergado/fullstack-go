package main

import (
	"net/http"

	"github.com/ferdiebergado/fullstack-go/internal/config"
	"github.com/ferdiebergado/fullstack-go/internal/db"
	"github.com/ferdiebergado/fullstack-go/internal/domain/activity"
	"github.com/ferdiebergado/fullstack-go/internal/domain/host"
	"github.com/ferdiebergado/fullstack-go/internal/domain/venue"
	"github.com/ferdiebergado/fullstack-go/internal/ui"
	myhttp "github.com/ferdiebergado/fullstack-go/pkg/http"
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

	// Create the Activity Handler.
	activityService := activity.NewActivityService(database)
	activityHandler := activity.NewActivityHandler(activityService, venueService, hostService)

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

	// Venues routes.
	router.Post("/api/venues", venueHandler.SaveVenue)

	// Hosts routes.
	router.Post("/api/hosts", hostHandler.SaveHost)

	// Home page
	router.Get("/{$}", func(w http.ResponseWriter, r *http.Request) {
		err := ui.RenderTemplate(w, "index.html", nil)

		if err != nil {
			myhttp.ErrorHandler(w, r, http.StatusBadRequest, "unable to render template", err)
			return
		}
	})

	// Not found handler
	// router.Get("/", func(w http.ResponseWriter, r *http.Request) {
	// 	status := http.StatusNotFound
	// 	myhttp.ErrorHandler(w, r, status, http.StatusText(status), errors.New("page not found"))
	// })

	return router
}
