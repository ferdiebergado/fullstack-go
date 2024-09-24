package templates

import (
	"bytes"
	"embed"
	"html/template"
	"log"
	"net/http"
)

const (
	templateDir = "templates"
	layoutFile  = "layout.html"
)

//go:embed **/*.html
var templatesFS embed.FS

func Render(w http.ResponseWriter, templateFile string) {

	templates, err := template.ParseFS(templatesFS, layoutFile, templateFile)

	if err != nil {
		log.Printf("Parse html files: %v\n", err)
		http.Error(w, "Html parsing error", http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer

	if err := templates.ExecuteTemplate(&buf, layoutFile, nil); err != nil {
		log.Printf("Execute template: %v\n", err)
		http.Error(w, "Template execution error", http.StatusInternalServerError)
		return
	}

	_, err = buf.WriteTo(w)

	if err != nil {
		log.Printf("Write to buffer: %v\n", err)
		http.Error(w, "Buffer write error", http.StatusInternalServerError)
		return
	}
}