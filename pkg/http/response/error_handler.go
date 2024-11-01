package response

import (
	"log"
	"net/http"
	"strings"

	"github.com/ferdiebergado/fullstack-go/internal/ui"
)

// Custom function to handle specific errors like 400, 403.
func ErrorHandler(w http.ResponseWriter, r *http.Request, status int, msg string, err error) {
	log.Printf("%s: %v\n", msg, err)

	// Check the "Accept" header for preferred content type
	acceptHeader := r.Header.Get("Accept")

	switch status {
	case http.StatusBadRequest:
		http.Error(w, "Bad Request", http.StatusBadRequest)
	case http.StatusForbidden:
		http.Error(w, "Forbidden", http.StatusForbidden)
	case http.StatusNotFound:
		if prefersJSON(acceptHeader) {
			http.NotFound(w, r)
			return
		}

		err := ui.RenderHTML(w, "404.html", nil)

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

// prefersJSON checks if "application/json" is present and prioritized in the Accept header
func prefersJSON(acceptHeader string) bool {
	mediaTypes := strings.Split(acceptHeader, ",")
	for _, mediaType := range mediaTypes {
		// Trim spaces and check if the media type is "application/json"
		if strings.TrimSpace(mediaType) == "application/json" {
			return true
		}
	}
	return false
}
