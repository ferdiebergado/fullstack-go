package host

import (
	"github.com/ferdiebergado/fullstack-go/internal/config"
	router "github.com/ferdiebergado/go-express"
)

const ResourceUrl = "/hosts"
const ApiRoute = config.ApiPrefix + ResourceUrl

func AddRoutes(r *router.Router, h HostHandler) {
	// Hosts routes.
	r.Post(ApiRoute, h.SaveHost)
}
