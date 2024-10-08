package http

import (
	"log"
	"net/http"

	"github.com/ferdiebergado/fullstack-go/internal/ui"
)

// Custom function to handle specific errors like 400, 403.
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int, msg string, err error) {
	log.Printf("%s: %v\n", msg, err)

	switch status {
	case http.StatusBadRequest:
		http.Error(w, "Bad Request", http.StatusBadRequest)
	case http.StatusForbidden:
		http.Error(w, "Forbidden", http.StatusForbidden)
	case http.StatusNotFound:
		// http.Error(w, "Not Found", http.StatusNotFound)
		err := ui.RenderTemplate(w, "404.html", nil)

		if err != nil {
			http.Error(w, "Unable to render template", http.StatusInternalServerError)
			return
		}
	case http.StatusInternalServerError:
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	default:
		http.Error(w, "An Error Occurred", status)
	}
}
