package activity

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ferdiebergado/fullstack-go/internal/db"
	"github.com/ferdiebergado/fullstack-go/internal/domain/host"
	"github.com/ferdiebergado/fullstack-go/internal/domain/venue"
	"github.com/ferdiebergado/fullstack-go/internal/ui"
	myhttp "github.com/ferdiebergado/fullstack-go/pkg/http"
	"github.com/ferdiebergado/fullstack-go/pkg/validator"
)

type Data struct {
	Activities []db.ListActivitiesRow
	Regions    []db.Region
}

type ActivityHandler struct {
	activityService ActivityService
	venueService    venue.VenueService
	hostService     host.HostService
}

func NewActivityHandler(activityService ActivityService, venueService venue.VenueService, hostService host.HostService) *ActivityHandler {
	return &ActivityHandler{activityService: activityService, venueService: venueService, hostService: hostService}
}

func (h *ActivityHandler) paginatedActiveActivities(w http.ResponseWriter, r *http.Request) interface{} {
	// TODO: validate query parameters
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")
	sortCol := r.URL.Query().Get("sort")
	sortDir := r.URL.Query().Get("sortDir")

	page, err := strconv.ParseInt(pageStr, 0, 64)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.ParseInt(limitStr, 0, 64)
	if err != nil || limit < 1 {
		limit = 10
	}

	offset := (page - 1) * limit
	totalItems, err := h.activityService.CountActivities(r.Context())
	totalPages := (totalItems + limit - 1) / limit

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusInternalServerError, "count activities", err)
		return nil
	}

	var activities any

	if sortDir == "1" {

		args := &db.ListActivitiesParams{
			Limit:   limit,
			Offset:  offset,
			Column1: &sortCol,
		}

		activities, err = h.activityService.ListActivities(r.Context(), *args)

	} else {
		args := &db.ListActivitiesOrderedDescParams{
			Limit:   limit,
			Offset:  offset,
			Column1: &sortCol,
		}

		activities, err = h.activityService.ListActivitiesOrderedDesc(r.Context(), *args)

	}

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusInternalServerError, "list activities", err)
		return nil
	}

	return &myhttp.PaginatedData{

		TotalItems: totalItems,
		TotalPages: totalPages,
		Page:       page,
		Limit:      limit,
		Data:       activities,
	}
}

func (h *ActivityHandler) ListActiveActivitiesJson(w http.ResponseWriter, r *http.Request) {
	// TODO: check type assertion
	data := h.paginatedActiveActivities(w, r)

	response := &myhttp.ApiResponse{
		Success: true,
		Data:    data,
	}

	err := ui.RenderJson(w, r, http.StatusOK, response)

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusInternalServerError, "unable to render json", err)
		return
	}
}

func (h *ActivityHandler) ListActiveActivities(w http.ResponseWriter, r *http.Request) {
	// TODO: check type assertion
	data := h.paginatedActiveActivities(w, r)

	err := ui.RenderTemplate(w, "activities/index.html", data)

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusInternalServerError, "unable to render template", err)
		return
	}
}

func (h *ActivityHandler) ShowCreateActivityForm(w http.ResponseWriter, r *http.Request) {
	venues, err := h.venueService.GetVenues(r.Context())

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusInternalServerError, "get venues", err)
		return
	}

	divisions, err := h.activityService.GetDivisions(r.Context())

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusInternalServerError, "get regions", err)
		return
	}

	hosts, err := h.activityService.GetHosts(r.Context())

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusInternalServerError, "get hosts", err)
		return
	}

	data := struct {
		Venues    []db.GetVenuesRow
		Divisions []db.GetDivisionWithRegionRow
		Hosts     []db.Host
	}{
		Venues:    venues,
		Divisions: divisions,
		Hosts:     hosts,
	}

	err = ui.RenderTemplate(w, "activities/create.html", data)

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusBadRequest, "unable to render template", err)
		return
	}
}

func (h *ActivityHandler) ParseId(idStr string) (int64, error) {
	id, err := strconv.ParseInt(idStr, 10, 64)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (h *ActivityHandler) GetActivity(w http.ResponseWriter, r *http.Request) {

	id, err := h.ParseId(r.PathValue("id"))

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusNotFound, "parse id", err)
		return
	}

	activity, err := h.activityService.FindActiveActivity(r.Context(), id)

	if err != nil || activity == nil {
		myhttp.ErrorHandler(w, r, http.StatusNotFound, "find active activity", err)
		return
	}

	err = ui.RenderJson(w, r, http.StatusOK, activity)

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusBadRequest, "unable to render json", err)
		return
	}
}

func (h *ActivityHandler) ShowActivity(w http.ResponseWriter, r *http.Request) {
	id, err := h.ParseId(r.PathValue("id"))

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusNotFound, "parse id", err)
		return
	}

	activity, err := h.activityService.FindActiveActivity(r.Context(), id)

	if err != nil || activity == nil {
		myhttp.ErrorHandler(w, r, http.StatusNotFound, "find active activity", err)
		return
	}

	err = ui.RenderTemplate(w, "activities/view.html", activity)

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusBadRequest, "unable to render template", err)
		return
	}
}

func (h *ActivityHandler) ShowEditActivityForm(w http.ResponseWriter, r *http.Request) {
	id, err := h.ParseId(r.PathValue("id"))

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusNotFound, "parse id", err)
		return
	}

	activity, err := h.activityService.FindActiveActivity(r.Context(), id)

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusNotFound, "find active activity", err)
		return
	}

	divisions, err := h.activityService.GetDivisions(r.Context())

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusInternalServerError, "get regions", err)
		return
	}

	venues, err := h.venueService.GetVenues(r.Context())

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusInternalServerError, "get venues", err)
		return
	}

	hosts, err := h.hostService.GetHosts(r.Context())

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusInternalServerError, "get hosts", err)
		return
	}

	data := struct {
		Activity  db.FindActivityRow
		Divisions []db.GetDivisionWithRegionRow
		Venues    []db.GetVenuesRow
		Hosts     []db.Host
	}{
		Activity:  *activity,
		Divisions: divisions,
		Venues:    venues,
		Hosts:     hosts,
	}

	err = ui.RenderTemplate(w, "activities/edit.html", data)

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusBadRequest, "unable to render template", err)
		return
	}
}

func (h *ActivityHandler) UpdateActivity(w http.ResponseWriter, r *http.Request) {
	id, err := h.ParseId(r.PathValue("id"))

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusNotFound, "parse id", err)
		return
	}

	data, err := ui.DecodeJson[db.UpdateActivityParams](r)

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusBadRequest, "decode json body", err)
		return
	}

	data.ID = id

	if data.Metadata == nil {
		data.Metadata = json.RawMessage(`{}`) // Set to an empty JSON object if nil
	}

	err = h.activityService.UpdateActivity(r.Context(), data)

	if err != nil {
		errorBag, ok := err.(*validator.ValidationErrorBag)

		if !ok {
			myhttp.ErrorHandler(w, r, http.StatusBadRequest, "update activity", err)
			return
		}

		response := &myhttp.ApiResponse{
			Success: false,
			Message: errorBag.Message,
			Errors:  errorBag.ValidationErrors,
		}

		err = ui.RenderJson(w, r, http.StatusBadRequest, response)

		if err != nil {
			myhttp.ErrorHandler(w, r, http.StatusBadRequest, "unable to render json", err)
			return
		}

		return
	}

	response := &myhttp.ApiResponse{
		Success: true,
		Message: "Activity updated.",
	}

	err = ui.RenderJson(w, r, http.StatusOK, response)

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusBadRequest, "unable to render json", err)
		return
	}
}

func (h *ActivityHandler) SaveActivity(w http.ResponseWriter, r *http.Request) {
	data, err := ui.DecodeJson[db.CreateActivityParams](r)

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusBadRequest, "decode json", err)
		return
	}

	if data.Metadata == nil {
		data.Metadata = json.RawMessage(`{}`) // Set to an empty JSON object if nil
	}

	activity, err := h.activityService.CreateActivity(r.Context(), data)

	if err != nil {
		errorBag, ok := err.(*validator.ValidationErrorBag)

		if !ok {
			myhttp.ErrorHandler(w, r, http.StatusBadRequest, "save activity", err)
			return
		}

		response := &myhttp.ApiResponse{
			Success: false,
			Message: errorBag.Message,
			Errors:  errorBag.ValidationErrors,
		}

		err = ui.RenderJson(w, r, http.StatusBadRequest, response)

		if err != nil {
			myhttp.ErrorHandler(w, r, http.StatusBadRequest, "unable to render json", err)
			return
		}

		return
	}

	response := &myhttp.ApiResponse{
		Success: true,
		Message: "Activity created.",
		Data:    activity,
	}

	err = ui.RenderJson(w, r, http.StatusCreated, response)

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusBadRequest, "unable to render json", err)
		return
	}
}

func (h *ActivityHandler) DeleteActivity(w http.ResponseWriter, r *http.Request) {

	id, err := h.ParseId(r.PathValue("id"))

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusNotFound, "parse id", err)
		return
	}

	err = h.activityService.DeleteActivity(r.Context(), id)

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusBadRequest, "delete activity", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
