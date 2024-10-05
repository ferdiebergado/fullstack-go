package http

type ApiResponse struct {
	Success bool              `json:"success"`
	Errors  []ValidationError `json:"errors,omitempty"`
	Data    interface{}       `json:"data,omitempty"`
}

type ValidationError struct {
	Field string `json:"field"`
	Error string `json:"error"`
}

type ValidationErrorBag struct {
	error
	Message          string
	ValidationErrors []ValidationError
}

func (e *ValidationErrorBag) Error() string {
	return e.Message
}
