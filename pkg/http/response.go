package http

import "github.com/ferdiebergado/fullstack-go/pkg/validator"

type ApiResponse[T any] struct {
	Success bool                        `json:"success"`
	Message string                      `json:"message"`
	Errors  []validator.ValidationError `json:"errors,omitempty"`
	Data    T                           `json:"data,omitempty"`
}

type PaginatedData[T any] struct {
	TotalItems int64 `json:"total_items"`
	TotalPages int64 `json:"total_pages"`
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
	Data       []T   `json:"data"`
}
