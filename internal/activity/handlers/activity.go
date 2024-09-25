package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ferdiebergado/fullstack-go/db"
	"github.com/ferdiebergado/fullstack-go/view"
)

type ActivityHandler struct {
	Queries *db.Queries
}

type Data struct {
	Activities []db.Activity
}

func (a *ActivityHandler) ActivityIndex(w http.ResponseWriter, r *http.Request) {
	activities, err := a.Queries.ListActivities(r.Context())

	if err != nil {
		log.Printf("list activities: %v\n", err)
		http.Error(w, "failed to get activities", http.StatusInternalServerError)
		return
	}

	data := &Data{Activities: activities}

	view.RenderTemplate(w, "activities/index.html", data)
}

func (a *ActivityHandler) CreateActivity(w http.ResponseWriter, r *http.Request) {
	view.RenderTemplate(w, "activities/create.html", nil)
}

func (a *ActivityHandler) FindActivity(w http.ResponseWriter, r *http.Request) *db.Activity {
	val := r.PathValue("id")

	id, err := strconv.ParseInt(val, 10, 32)

	if err != nil {
		log.Printf("parse path value: %v\n", err)
		http.Error(w, "Activity not found", http.StatusNotFound)
		return nil
	}

	log.Printf("val: %s, id: %d\n", val, id)

	activity, err := a.Queries.FindActivity(r.Context(), int32(id))

	if err != nil {
		log.Printf("find activity: %v\n", err)
		http.Error(w, "Activity not found", http.StatusNotFound)
		return nil
	}

	return &activity
}

func (a *ActivityHandler) ViewActivity(w http.ResponseWriter, r *http.Request) {

	activity := a.FindActivity(w, r)
	view.RenderTemplate(w, "activities/view.html", activity)
}

func (a *ActivityHandler) EditActivity(w http.ResponseWriter, r *http.Request) {
	activity := a.FindActivity(w, r)

	view.RenderTemplate(w, "activities/edit.html", activity)
}

func (a *ActivityHandler) UpdateActivity(w http.ResponseWriter, r *http.Request) {

	activity := a.FindActivity(w, r)

	r.ParseForm()

	// TODO: validate form values
	startDate, err := time.Parse(time.DateOnly, r.FormValue("start_date"))

	if err != nil {
		log.Printf("parse start date: %v\n", err)
		http.Error(w, "invalid start date", http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse(time.DateOnly, r.FormValue("end_date"))

	if err != nil {
		log.Printf("parse end date: %v\n", err)
		http.Error(w, "invalid end date", http.StatusBadRequest)
		return
	}

	params := db.UpdateActivityParams{
		Title:     r.FormValue("title"),
		StartDate: startDate,
		EndDate:   endDate,
		Venue:     db.StringToNullString(r.FormValue("venue")),
		Host:      db.StringToNullString(r.FormValue("host")),
		Metadata:  json.RawMessage(`{}`),
		ID:        activity.ID,
	}

	err = a.Queries.UpdateActivity(r.Context(), params)

	if err != nil {
		log.Printf("update activity: %v\n", err)
		http.Error(w, "failed to update activity", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/activities", http.StatusSeeOther)
}

func (a *ActivityHandler) SaveActivity(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// TODO: validate form values
	startDate, err := time.Parse(time.DateOnly, r.FormValue("start_date"))

	if err != nil {
		log.Printf("parse start date: %v\n", err)
		http.Error(w, "invalid start date", http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse(time.DateOnly, r.FormValue("end_date"))

	if err != nil {
		log.Printf("parse end date: %v\n", err)
		http.Error(w, "invalid end date", http.StatusBadRequest)
		return
	}

	params := db.CreateActivityParams{
		Title:     r.FormValue("title"),
		StartDate: startDate,
		EndDate:   endDate,
		Venue:     db.StringToNullString(r.FormValue("venue")),
		Host:      db.StringToNullString(r.FormValue("host")),
		Metadata:  json.RawMessage(`{}`),
	}

	_, err = a.Queries.CreateActivity(r.Context(), params)

	if err != nil {
		log.Printf("save activity: %v\n", err)
		http.Error(w, "failed to create activity", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/activities", http.StatusSeeOther)
}
