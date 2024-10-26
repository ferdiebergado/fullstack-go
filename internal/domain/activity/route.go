package activity

import (
	router "github.com/ferdiebergado/go-express"
)

func AddRoutes(router *router.Router, handler ActivityHandler) {

	// html pages
	router.Get("/activities", handler.ListActiveActivities)
	router.Get("/activities/create", handler.ShowCreateActivityForm)
	router.Get("/activities/{id}", handler.ShowActivity)
	router.Get("/activities/{id}/edit", handler.ShowEditActivityForm)

	// api routes
	router.Get("/api/activities", handler.ListActiveActivitiesJson)
	router.Get("/api/activities/{id}", handler.GetActivity)
	router.Post("/api/activities", handler.SaveActivity)
	router.Put("/api/activities/{id}", handler.UpdateActivity)
	router.Delete("/api/activities/{id}", handler.DeleteActivity)

}
