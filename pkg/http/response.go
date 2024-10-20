package http

import "github.com/ferdiebergado/fullstack-go/pkg/validator"

type ApiResponse struct {
	Success bool                        `json:"success"`
	Message string                      `json:"message"`
	Errors  []validator.ValidationError `json:"errors,omitempty"`
	Data    any                         `json:"data,omitempty"`
}

type PaginatedData struct {
	Columns    string `json:"columns"`
	Url        string `json:"url"`
	TotalItems int64  `json:"total_items"`
	TotalPages int64  `json:"total_pages"`
	Page       int64  `json:"page"`
	Limit      int64  `json:"limit"`
	Data       any    `json:"data"`
}
