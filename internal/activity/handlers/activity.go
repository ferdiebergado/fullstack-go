package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ferdiebergado/fullstack-go/db"
	myhttp "github.com/ferdiebergado/fullstack-go/pkg/http"
	"github.com/ferdiebergado/fullstack-go/view"
)

type ActivityHandler struct {
	Db      *sql.DB
	Queries *db.Queries
}

type Data struct {
	Activities []db.ActiveActivity
}

type FormDates struct {
	startDate db.Date
	endDate   db.Date
}

func NewActivityHandler(database *sql.DB, queries *db.Queries) *ActivityHandler {
	return &ActivityHandler{Db: database, Queries: queries}
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

func (a *ActivityHandler) ListActiveActivities(w http.ResponseWriter, r *http.Request) {

	activities, err := a.Queries.ListActivities(r.Context())

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusNotFound, "list activities", err)
		return
	}

	data := &Data{Activities: activities}

	acceptHeader := r.Header.Get("Accept")
	acceptedTypes := strings.Split(acceptHeader, ",")

	// Trim spaces and check each accepted media type
	for _, mediaType := range acceptedTypes {
		mediaType = strings.TrimSpace(mediaType)

		if mediaType == "application/json" {

			err = view.RenderJson(w, r, http.StatusOK, data)
			if err != nil {
				myhttp.ErrorHandler(w, r, http.StatusNotFound, "render json activities", err)
			}
			return
		} else if mediaType == "text/html" {
			view.RenderTemplate(w, "activities/index.html", data)
			return
		}
	}

	// Default fallback if no match
	// w.Header().Set("Content-Type", "text/plain")
	// w.Write([]byte("Default response in plain text"))
	w.WriteHeader(http.StatusBadRequest)
}

func (a *ActivityHandler) CreateActivity(w http.ResponseWriter, r *http.Request) {
	view.RenderTemplate(w, "activities/create.html", nil)
}

func (a *ActivityHandler) FindActiveActivity(ctx context.Context, idStr string) (*db.ActiveActivity, error) {
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

	activity, err := a.FindActiveActivity(r.Context(), r.PathValue("id"))

	if err != nil || activity == nil {
		myhttp.ErrorHandler(w, r, http.StatusNotFound, "find active activity", err)
		return
	}

	// view.RenderTemplate(w, "activities/view.html", activity)
	err = view.RenderJson(w, r, http.StatusOK, activity)

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusNotFound, "render json activity", err)
	}
}

func (a *ActivityHandler) EditActivity(w http.ResponseWriter, r *http.Request) {
	activity, err := a.FindActiveActivity(r.Context(), r.PathValue("id"))

	if err != nil || activity == nil {
		myhttp.ErrorHandler(w, r, http.StatusNotFound, "find active activity", err)
		return
	}

	view.RenderTemplate(w, "activities/edit.html", activity)
}

func (a *ActivityHandler) ParseFormDates(w http.ResponseWriter, r *http.Request) *FormDates {
	var startDate, endDate time.Time

	// TODO: validate form values
	startDate, err := time.Parse(time.DateOnly, r.FormValue("start_date"))

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusBadRequest, "parse start date", err)
		return nil
	}

	endDate, err = time.Parse(time.DateOnly, r.FormValue("end_date"))

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusBadRequest, "parse end date", err)
		return nil
	}

	return &FormDates{
		startDate: db.NewDate(startDate),
		endDate:   db.NewDate(endDate),
	}
}

func (a *ActivityHandler) UpdateActivityJson(w http.ResponseWriter, r *http.Request) {
	activity, err := a.FindActiveActivity(r.Context(), r.PathValue("id"))

	if err != nil || activity == nil {
		myhttp.ErrorHandler(w, r, http.StatusNotFound, "find active activity", err)
		return
	}

	defer r.Body.Close()

	var data db.UpdateActivityParams

	err = json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusBadRequest, "decode json body", err)
		return
	}

	data.ID = activity.ID

	if data.Metadata == nil {
		data.Metadata = json.RawMessage(`{}`) // Set to an empty JSON object if nil
	}

	err = a.Queries.UpdateActivity(r.Context(), data)

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusBadRequest, "update activity", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *ActivityHandler) UpdateActivity(w http.ResponseWriter, r *http.Request) {

	activity, err := a.FindActiveActivity(r.Context(), r.PathValue("id"))

	if err != nil || activity == nil {
		myhttp.ErrorHandler(w, r, http.StatusNotFound, "find active activity", err)
		return
	}

	err = r.ParseForm()

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusBadRequest, "parse form", err)
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
		myhttp.ErrorHandler(w, r, http.StatusBadRequest, "update activity", err)
		return
	}

	http.Redirect(w, r, "/activities", http.StatusSeeOther)
}

func (a *ActivityHandler) SaveActivityJson(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var data db.CreateActivityParams

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusBadRequest, "decode json", err)
		return
	}

	if data.Metadata == nil {
		data.Metadata = json.RawMessage(`{}`) // Set to an empty JSON object if nil
	}

	activity, err := a.Queries.CreateActivity(r.Context(), data)

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusBadRequest, "save activity", err)
		return
	}

	// w.WriteHeader(http.StatusCreated)
	view.RenderJson(w, r, http.StatusCreated, activity)
}

func (a *ActivityHandler) SaveActivity(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusBadRequest, "parse form", err)
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
		myhttp.ErrorHandler(w, r, http.StatusBadRequest, "save activity", err)
		return
	}

	http.Redirect(w, r, "/activities", http.StatusSeeOther)
}

func (a *ActivityHandler) DeleteActivity(w http.ResponseWriter, r *http.Request) {

	activity, err := a.FindActiveActivity(r.Context(), r.PathValue("id"))

	if err != nil || activity == nil {
		myhttp.ErrorHandler(w, r, http.StatusNotFound, "find active activity", err)
		return
	}

	err = a.Queries.DeleteActivity(r.Context(), activity.ID)

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusBadRequest, "delete activity", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
