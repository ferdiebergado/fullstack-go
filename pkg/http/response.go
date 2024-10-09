package http

import "github.com/ferdiebergado/fullstack-go/pkg/validator"

type ApiResponse struct {
	Success bool                        `json:"success"`
	Message string                      `json:"message"`
	Errors  []validator.ValidationError `json:"errors,omitempty"`
	Data    interface{}                 `json:"data,omitempty"`
}
