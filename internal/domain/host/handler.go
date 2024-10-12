package host

import (
	"net/http"

	"github.com/ferdiebergado/fullstack-go/internal/ui"
	myhttp "github.com/ferdiebergado/fullstack-go/pkg/http"
	"github.com/ferdiebergado/fullstack-go/pkg/validator"
)

type HostHandler struct {
	hostService HostService
}

func NewHostHandler(s HostService) *HostHandler {
	return &HostHandler{hostService: s}
}

func (h *HostHandler) SaveHost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	type createHostParams struct {
		Name string `json:"name"`
	}

	data, err := ui.DecodeJson[createHostParams](r)

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusBadRequest, "decode json", err)
		return
	}

	host, err := h.hostService.CreateHost(r.Context(), data.Name)

	if err != nil {
		errorBag, ok := err.(*validator.ValidationErrorBag)

		if !ok {
			myhttp.ErrorHandler(w, r, http.StatusBadRequest, "save host", err)
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
		Message: "Host created successfully!",
		Data:    host,
	}

	err = ui.RenderJson(w, r, http.StatusCreated, response)

	if err != nil {
		myhttp.ErrorHandler(w, r, http.StatusBadRequest, "unable to render json", err)
		return
	}
}
