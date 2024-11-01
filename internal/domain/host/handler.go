package host

import (
	"net/http"

	"github.com/ferdiebergado/fullstack-go/internal/db"
	"github.com/ferdiebergado/fullstack-go/internal/ui"
	"github.com/ferdiebergado/fullstack-go/pkg/http/response"
	"github.com/ferdiebergado/fullstack-go/pkg/validator"
)

type HostHandler struct {
	hostService HostService
}

func NewHostHandler(s HostService) *HostHandler {
	return &HostHandler{hostService: s}
}

func (h *HostHandler) SaveHost(w http.ResponseWriter, r *http.Request) {
	type createHostParams struct {
		Name string `json:"name"`
	}

	data, err := ui.DecodeJson[createHostParams](r)

	if err != nil {
		response.ErrorHandler(w, r, http.StatusBadRequest, "decode json", err)
		return
	}

	host, err := h.hostService.CreateHost(r.Context(), data.Name)

	if err != nil {
		errorBag, ok := err.(*validator.ValidationErrorBag)

		if !ok {
			response.ErrorHandler(w, r, http.StatusBadRequest, "save host", err)
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

	res := &response.ApiResponse[[]db.Host]{
		Meta: response.ResponseMeta{
			Message: "Host created successfully!",
		},
		Data: []db.Host{host},
	}

	err = ui.RenderJson(w, http.StatusCreated, res)

	if err != nil {
		response.ErrorHandler(w, r, http.StatusBadRequest, "unable to render json", err)
		return
	}
}
