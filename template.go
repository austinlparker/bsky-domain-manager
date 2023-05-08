package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type PageData struct {
	Handle     string
	Value      string
	GetURL     string
	PostURL    string
	PutURL     string
	ResolveURL string
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmpl = filepath.Join("templates", tmpl+".html")
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
