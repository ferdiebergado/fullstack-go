package venue

import (
	"net/http"

	"github.com/ferdiebergado/fullstack-go/internal/db"
	"github.com/ferdiebergado/fullstack-go/internal/ui"
	"github.com/ferdiebergado/fullstack-go/pkg/http/response"
	"github.com/ferdiebergado/fullstack-go/pkg/validator"
)

type VenueHandler struct {
	venueService VenueService
}

func NewVenueHandler(s VenueService) *VenueHandler {
	return &VenueHandler{venueService: s}
}

func (h *VenueHandler) SaveVenue(w http.ResponseWriter, r *http.Request) {
	data, err := ui.DecodeJson[db.CreateVenueParams](r)

	if err != nil {
		response.ErrorHandler(w, r, http.StatusBadRequest, "decode json", err)
		return
	}

	venue, err := h.venueService.CreateVenue(r.Context(), data)

	if err != nil {
		errorBag, ok := err.(*validator.ValidationErrorBag)

		if !ok {
			response.ErrorHandler(w, r, http.StatusBadRequest, "save venue", err)
			return
		}

		res := &response.ApiResponse[any]{
			Meta: response.ResponseMeta{
				Message: errorBag.Message,
				Errors:  errorBag.ValidationErrors,
			},
		}

		err = ui.RenderJson(w, http.StatusBadRequest, res)

		if err != nil {
			response.ErrorHandler(w, r, http.StatusBadRequest, "unable to render json", err)
			return
		}

		return
	}

	res := &response.ApiResponse[[]*db.Venue]{
		Meta: response.ResponseMeta{
			Message: "Venue created successfully!",
		},
		Data: []*db.Venue{venue},
	}

	err = ui.RenderJson(w, http.StatusCreated, res)

	if err != nil {
		response.ErrorHandler(w, r, http.StatusBadRequest, "unable to render json", err)
		return
	}
}
