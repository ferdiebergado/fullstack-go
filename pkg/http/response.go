package http

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

const (
	templateDir = "templates"
	layoutFile  = "layout.html"
)

func HTMLResponse(w http.ResponseWriter, templateFile string) {
	templates, err := template.ParseFiles(filepath.Join(templateDir, layoutFile), filepath.Join(templateDir, templateFile))

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
