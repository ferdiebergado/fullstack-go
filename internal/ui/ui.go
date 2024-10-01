package ui

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
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

func RenderJson[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

func DecodeJson[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}
