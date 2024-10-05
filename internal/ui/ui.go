package ui

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

const (
	templateDir = "templates"
	layoutFile  = "layout.html"
)

//go:embed templates/*
var templatesFS embed.FS

func RenderTemplate(w http.ResponseWriter, templateFile string, data interface{}) error {
	layoutPath := filepath.Join(templateDir, layoutFile)
	templatePath := filepath.Join(templateDir, templateFile)

	t, err := template.ParseFS(templatesFS, layoutPath, templatePath)

	if err != nil {
		return fmt.Errorf("parse html files: %w", err)
	}

	var buf bytes.Buffer

	if err := t.ExecuteTemplate(&buf, layoutFile, data); err != nil {
		return fmt.Errorf("execute template: %w", err)
	}

	_, err = buf.WriteTo(w)

	if err != nil {
		return fmt.Errorf("write to buffer: %w", err)
	}

	return nil
}

func RenderJson[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return EncodeJson(w, v)
}

func EncodeJson[T any](w http.ResponseWriter, v T) error {
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
