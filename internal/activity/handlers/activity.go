package handlers

import (
	"log"
	"net/http"

	"github.com/ferdiebergado/fullstack-go/db"
	"github.com/ferdiebergado/fullstack-go/view"
)

type ActivityHandler struct {
	Queries *db.Queries
}

func (a *ActivityHandler) ActivityIndex(w http.ResponseWriter, r *http.Request) {
	activities, err := a.Queries.ListActivities(r.Context())

	if err != nil {
		log.Printf("list activities: %v\n", err)
		http.Error(w, "failed to get activities", http.StatusInternalServerError)
		return
	}

	view.RenderTemplate(w, "activities/index.html", activities)
}
