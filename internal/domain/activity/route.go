package activity

import (
	"net/http"

	myhttp "github.com/ferdiebergado/fullstack-go/pkg/http"
)

func AddRoutes(router *myhttp.Router, handler ActivityHandler) {

	router.Handle("GET /activities", http.HandlerFunc(handler.ListActiveActivities))
	router.Handle("GET /activities/create", http.HandlerFunc(handler.CreateActivity))
	router.Handle("GET /activities/{id}", http.HandlerFunc(handler.ViewActivity))
	router.Handle("GET /activities/{id}/edit", http.HandlerFunc(handler.EditActivity))
	router.Handle("GET /api/activities/{id}", http.HandlerFunc(handler.GetActivity))
	router.Handle("POST /api/activities", http.HandlerFunc(handler.SaveActivity))
	router.Handle("PUT /api/activities/{id}", http.HandlerFunc(handler.UpdateActivity))
	router.Handle("DELETE /api/activities/{id}", http.HandlerFunc(handler.DeleteActivity))

}
