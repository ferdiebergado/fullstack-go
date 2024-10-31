package activity

import (
	"github.com/ferdiebergado/fullstack-go/internal/config"
	router "github.com/ferdiebergado/go-express"
)

const ResourceUrl string = "/activities"
const ApiRoute = config.ApiPrefix + ResourceUrl

func AddRoutes(router *router.Router, handler ActivityHandler) {

	const routeWithId = ResourceUrl + config.ResourceIdPath
	const apiRouteWithId = ApiRoute + config.ResourceIdPath

	// html pages
	router.Get(ResourceUrl, handler.ListActiveActivities)
	router.Get(ResourceUrl+"/create", handler.ShowCreateActivityForm)
	router.Get(routeWithId, handler.ShowActivity)
	router.Get(routeWithId+"/edit", handler.ShowEditActivityForm)

	// api routes
	router.Get(ApiRoute, handler.ListActiveActivitiesJson)
	router.Get(apiRouteWithId, handler.GetActivity)
	router.Post(ApiRoute, handler.SaveActivity)
	router.Put(apiRouteWithId, handler.UpdateActivity)
	router.Delete(apiRouteWithId, handler.DeleteActivity)

}
