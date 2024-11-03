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
	templateDir      = "templates"
	layoutFile       = "base.html"
	partialTemplates = "partials/*.html"
)

//go:embed templates/*
var templatesFS embed.FS

func RenderHTML(w http.ResponseWriter, templateFile string, data any) error {
	layoutTemplate := filepath.Join(templateDir, layoutFile)
	partialTemplates := filepath.Join(templateDir, partialTemplates)
	targetTemplate := filepath.Join(templateDir, templateFile)

	funcMap := template.FuncMap{
		"attr": func(s string) template.HTMLAttr {
			return template.HTMLAttr(s)
		},
		"safe": func(s string) template.HTML {
			return template.HTML(s)
		},
		"url": func(s string) template.URL {
			return template.URL(s)
		},
		"js": func(s string) template.JS {
			return template.JS(s)
		},
		"jsstr": func(s string) template.JSStr {
			return template.JSStr(s)
		},
		"css": func(s string) template.CSS {
			return template.CSS(s)
		},
	}

	templates, err := template.New("template").Funcs(funcMap).ParseFS(templatesFS, layoutTemplate, partialTemplates, targetTemplate)

	if err != nil {
		return fmt.Errorf("parse html files: %w", err)
	}

	var buf bytes.Buffer

	if err := templates.ExecuteTemplate(&buf, layoutFile, data); err != nil {
		return fmt.Errorf("execute template: %w", err)
	}

	_, err = buf.WriteTo(w)

	if err != nil {
		return fmt.Errorf("write to buffer: %w", err)
	}

	return nil
}

func RenderJson[T any](w http.ResponseWriter, status int, v T) error {
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

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, err
	}

	return v, nil
}
