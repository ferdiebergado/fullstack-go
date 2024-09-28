package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/ferdiebergado/fullstack-go/db"
	"github.com/ferdiebergado/fullstack-go/pkg/debug"
	"github.com/ferdiebergado/fullstack-go/view"
)

type ActivityHandler struct {
	Queries *db.Queries
}

type Data struct {
	Activities []db.Activity
}

type FormDates struct {
	startDate db.Date
	endDate   db.Date
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

func (a *ActivityHandler) FindActivity(ctx context.Context, idStr string) (*db.Activity, error) {
	id, err := strconv.ParseInt(idStr, 10, 32)
	if err != nil {
		log.Printf("parse path value: %v\n", err)
		return nil, fmt.Errorf("invalid activity ID: %v", err)
	}

	activity, err := a.Queries.FindActivity(ctx, int32(id))
	if err != nil {
		log.Printf("find activity: %v\n", err)
		return nil, fmt.Errorf("activity not found: %v", err)
	}

	return &activity, nil
}

func (a *ActivityHandler) GetActivity(w http.ResponseWriter, r *http.Request) {

	// Extract ID from the path
	activity, err := a.FindActivity(r.Context(), r.PathValue("id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// view.RenderTemplate(w, "activities/view.html", activity)
	view.RenderJson(w, r, http.StatusOK, activity)
}

func (a *ActivityHandler) EditActivity(w http.ResponseWriter, r *http.Request) {
	// Extract ID from the path
	activity, err := a.FindActivity(r.Context(), r.PathValue("id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	view.RenderTemplate(w, "activities/edit.html", activity)
}

func (a *ActivityHandler) ParseFormDates(w http.ResponseWriter, r *http.Request) *FormDates {
	var startDate, endDate time.Time

	// TODO: validate form values
	startDate, err := time.Parse(time.DateOnly, r.FormValue("start_date"))

	if err != nil {
		log.Printf("parse start date: %v\n", err)
		http.Error(w, "invalid start date", http.StatusBadRequest)
		return nil
	}

	endDate, err = time.Parse(time.DateOnly, r.FormValue("end_date"))

	if err != nil {
		log.Printf("parse end date: %v\n", err)
		http.Error(w, "invalid end date", http.StatusBadRequest)
		return nil
	}

	return &FormDates{
		startDate: db.NewDate(startDate),
		endDate:   db.NewDate(endDate),
	}
}

func (a *ActivityHandler) UpdateActivityJson(w http.ResponseWriter, r *http.Request) {
	activity, err := a.FindActivity(r.Context(), r.PathValue("id"))

	if err != nil || activity == nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	defer r.Body.Close()

	var data db.UpdateActivityParams

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("decode json body: %v\n", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	data.ID = activity.ID

	// Ensure Metadata is set
	if data.Metadata == nil {
		data.Metadata = json.RawMessage(`{}`) // Set to an empty JSON object if nil
	}

	err = a.Queries.UpdateActivity(r.Context(), data)

	if err != nil {
		log.Printf("update activity: %v\n", err)
		http.Error(w, "failed to update activity", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *ActivityHandler) UpdateActivity(w http.ResponseWriter, r *http.Request) {

	// Extract ID from the path
	activity, err := a.FindActivity(r.Context(), r.PathValue("id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = r.ParseForm()

	if err != nil {
		log.Printf("parse form: %v\n", err)
		http.Error(w, "failed to parse the form", http.StatusBadRequest)
		return
	}

	formDates := a.ParseFormDates(w, r)

	venue := r.FormValue("venue")
	host := r.FormValue("host")

	params := db.UpdateActivityParams{
		Title:     r.FormValue("title"),
		StartDate: formDates.startDate,
		EndDate:   formDates.endDate,
		Venue:     &venue,
		Host:      &host,
		Metadata:  json.RawMessage(`{}`),
		ID:        activity.ID,
	}

	err = a.Queries.UpdateActivity(r.Context(), params)

	if err != nil {
		log.Printf("update activity: %v\n", err)
		http.Error(w, "failed to update activity", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/activities", http.StatusSeeOther)
}

func (a *ActivityHandler) SaveActivityJson(w http.ResponseWriter, r *http.Request) {
	// Read the body of the request
	defer r.Body.Close()
	var data db.CreateActivityParams

	// Decode the JSON body into the struct
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Printf("decode json body: %v\n", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	debug.DumpStruct(data)

	// Ensure Metadata is set
	if data.Metadata == nil {
		data.Metadata = json.RawMessage(`{}`) // Set to an empty JSON object if nil
	}

	_, err = a.Queries.CreateActivity(r.Context(), data)

	if err != nil {
		log.Printf("save activity: %v\n", err)
		http.Error(w, "failed to create activity", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (a *ActivityHandler) SaveActivity(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Printf("parse form: %v\n", err)
		http.Error(w, "failed to parse the form", http.StatusBadRequest)
		return
	}

	// TODO: validate form values
	formDates := a.ParseFormDates(w, r)

	venue := r.FormValue("venue")
	host := r.FormValue("host")

	params := db.CreateActivityParams{
		Title:     r.FormValue("title"),
		StartDate: formDates.startDate,
		EndDate:   formDates.endDate,
		Venue:     &venue,
		Host:      &host,
		Metadata:  json.RawMessage(`{}`),
	}

	_, err = a.Queries.CreateActivity(r.Context(), params)

	if err != nil {
		log.Printf("save activity: %v\n", err)
		http.Error(w, "failed to create activity", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/activities", http.StatusSeeOther)
}

func (a *ActivityHandler) DeleteActivity(w http.ResponseWriter, r *http.Request) {

	activity, err := a.FindActivity(r.Context(), r.PathValue("id"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = a.Queries.DeleteActivity(r.Context(), activity.ID)

	if err != nil {
		log.Printf("delete activity: %v\n", err)
		http.Error(w, "failed to delete activity", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
