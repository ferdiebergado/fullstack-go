package view

import (
	"bytes"
	"embed"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

const (
	templateDir = "templates"
	layoutFile  = "layout.html"
)

//go:embed templates/*
var templatesFS embed.FS

func RenderTemplate(w http.ResponseWriter, templateFile string, data interface{}) {
	layoutPath := filepath.Join(templateDir, layoutFile)
	templatePath := filepath.Join(templateDir, templateFile)

	t, err := template.ParseFS(templatesFS, layoutPath, templatePath)

	if err != nil {
		log.Printf("Parse html files: %v\n", err)
		http.Error(w, "Html parsing error", http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer

	if err := t.ExecuteTemplate(&buf, layoutFile, data); err != nil {
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
