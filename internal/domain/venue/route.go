package venue

import (
	"github.com/ferdiebergado/fullstack-go/internal/config"
	router "github.com/ferdiebergado/go-express"
)

const ResourceUrl string = "/venues"
const ApiRoute = config.ApiPrefix + ResourceUrl

func AddRoutes(r *router.Router, h VenueHandler) {
	// Venues routes.
	r.Post(ApiRoute, h.SaveVenue)
}
