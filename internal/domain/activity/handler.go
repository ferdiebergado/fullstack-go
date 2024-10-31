package activity

import (
	"database/sql"
	"encoding/json"
	"errors"
	"html/template"
	"net/http"
	"strconv"

	"github.com/ferdiebergado/fullstack-go/internal/db"
	"github.com/ferdiebergado/fullstack-go/internal/domain/division"
	"github.com/ferdiebergado/fullstack-go/internal/domain/host"
	"github.com/ferdiebergado/fullstack-go/internal/domain/venue"
	"github.com/ferdiebergado/fullstack-go/internal/ui"
	myhttp "github.com/ferdiebergado/fullstack-go/pkg/http"
	"github.com/ferdiebergado/fullstack-go/pkg/validator"
)

type ActivityHandler struct {
	activityService ActivityService
	venueService    venue.VenueService
	hostService     host.HostService
	divisionService division.DivisionService
}

func NewActivityHandler(as ActivityService, vs venue.VenueService, hs host.HostService, ds division.DivisionService) *ActivityHandler {
	return &ActivityHandler{activityService: as, venueService: vs, hostService: hs, divisionService: ds}
}

func (h *ActivityHandler) getPaginatedData(r *http.Request) (*myhttp.PaginatedData[db.ListActiveActivitiesRow], error) {

	paginatedData, err := h.activityService.ListActivities(r.Context(), r.URL.Query())

	if err != nil {
		return nil, err
	}

	return paginatedData, nil
}

func (h *ActivityHandler) ListActiveActivitiesJson(w http.ResponseWriter, r *http.Request) {

	paginatedData, err := h.getPaginatedData(r)

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusInternalServerError, "list activities", err)
		return
	}

	response := &myhttp.ApiResponse[[]db.ListActiveActivitiesRow]{
		Meta: myhttp.ResponseMeta{
			Pagination: paginatedData.Pagination,
		},
		Data: paginatedData.Data,
	}

	err = ui.RenderJson(w, http.StatusOK, response)

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusInternalServerError, "unable to render json", err)
		return
	}

}

func (h *ActivityHandler) ListActiveActivities(w http.ResponseWriter, r *http.Request) {
	tableHeaders := []myhttp.TableHeader{
		{Field: "title", Label: "Title"},
		{Field: "start_date", Label: "Start Date"},
		{Field: "end_date", Label: "End Date"},
		{Field: "venue", Label: "Venue"},
		{Field: "host", Label: "Host"},
	}

	jsonData, err := json.Marshal(tableHeaders)

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusInternalServerError, "unable to marshal table headers", err)
		return
	}

	data := &myhttp.TableData{
		ApiUrl:       ApiRoute,
		TableHeaders: template.JS(jsonData),
	}

	err = ui.RenderHTML(w, "pages/activities/index.html", data)

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

	divisions, err := h.divisionService.GetDivisions(r.Context())

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusInternalServerError, "get regions", err)
		return
	}

	hosts, err := h.hostService.GetHosts(r.Context())

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

	err = ui.RenderHTML(w, "pages/activities/create.html", data)

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusBadRequest, "unable to render template", err)
		return
	}
}

func (h *ActivityHandler) ParseId(idStr string) (int64, error) {
	return strconv.ParseInt(idStr, 10, 64)
}

func (h *ActivityHandler) getActivity(w http.ResponseWriter, r *http.Request) *db.ActiveActivityDetail {
	id, err := h.ParseId(r.PathValue("id"))

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusNotFound, "parse id", err)
		return nil
	}

	activity, err := h.activityService.FindActiveActivityDetails(r.Context(), id)

	if err != nil {

		status := http.StatusInternalServerError

		if errors.Is(err, sql.ErrNoRows) {

			status = http.StatusNotFound

		}

		myhttp.ErrorHandler(w, r, status, "find active activity", err)
		return nil
	}

	return activity
}

func (h *ActivityHandler) GetActivity(w http.ResponseWriter, r *http.Request) {

	activity := h.getActivity(w, r)

	err := ui.RenderJson(w, http.StatusOK, activity)

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusInternalServerError, "unable to render json", err)
		return
	}
}

func (h *ActivityHandler) ShowActivity(w http.ResponseWriter, r *http.Request) {
	activity := h.getActivity(w, r)

	err := ui.RenderHTML(w, "pages/activities/view.html", activity)

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusInternalServerError, "unable to render template", err)
		return
	}
}

func (h *ActivityHandler) ShowEditActivityForm(w http.ResponseWriter, r *http.Request) {
	activity := h.getActivity(w, r)

	divisions, err := h.divisionService.GetDivisions(r.Context())

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
		Activity  db.ActiveActivityDetail
		Divisions []db.GetDivisionWithRegionRow
		Venues    []db.GetVenuesRow
		Hosts     []db.Host
	}{
		Activity:  *activity,
		Divisions: divisions,
		Venues:    venues,
		Hosts:     hosts,
	}

	err = ui.RenderHTML(w, "pages/activities/edit.html", data)

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

		response := &myhttp.ApiResponse[any]{
			Meta: myhttp.ResponseMeta{
				Message: errorBag.Message,
				Errors:  errorBag.ValidationErrors,
			},
		}

		err = ui.RenderJson(w, http.StatusBadRequest, response)

		if err != nil {
			myhttp.ErrorHandler(w, r, http.StatusBadRequest, "unable to render json", err)
			return
		}

		return
	}

	response := &myhttp.ApiResponse[any]{
		Meta: myhttp.ResponseMeta{
			Message: "Activity updated.",
		},
	}

	err = ui.RenderJson(w, http.StatusOK, response)

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

		response := &myhttp.ApiResponse[any]{
			Meta: myhttp.ResponseMeta{
				Message: errorBag.Message,
				Errors:  errorBag.ValidationErrors,
			},
		}

		err = ui.RenderJson(w, http.StatusBadRequest, response)

		if err != nil {
			myhttp.ErrorHandler(w, r, http.StatusBadRequest, "unable to render json", err)
			return
		}

		return
	}

	response := &myhttp.ApiResponse[[]db.Activity]{
		Meta: myhttp.ResponseMeta{
			Message: "Activity created.",
		},
		Data: []db.Activity{*activity},
	}

	err = ui.RenderJson(w, http.StatusCreated, response)

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

		status := http.StatusInternalServerError

		if errors.Is(err, sql.ErrNoRows) {
			status = http.StatusNotFound
		}

		myhttp.ErrorHandler(w, r, status, "delete activity", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
