package response

import (
	"html/template"

	"github.com/ferdiebergado/fullstack-go/pkg/validator"
)

type ApiResponse[T any] struct {
	Meta ResponseMeta `json:"meta"`
	Data T            `json:"data,omitempty"`
}

type ResponseMeta struct {
	Message    string                      `json:"message,omitempty"`
	Errors     []validator.ValidationError `json:"errors,omitempty"`
	Pagination *PaginationMeta             `json:"pagination,omitempty"`
}

type PaginationMeta struct {
	TotalItems int64 `json:"total_items"`
	TotalPages int64 `json:"total_pages"`
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
}

type PaginatedData[T any] struct {
	Pagination *PaginationMeta
	Data       []T
}

type TableHeader struct {
	Field string `json:"field"`
	Label string `json:"label"`
}

type TableData struct {
	ApiUrl       string
	TableHeaders template.JS
}
