package utils

import (
	"fmt"
	"net/http"
	"text/template"
)

type PageData struct {
	Title   string
	Message string
	Links   []Link
	Names   []string
}

type Link struct {
	Text string
	URL  string
}

// テンプレートをレンダリング
func RenderTemplate(w http.ResponseWriter, tmpl string, data PageData) {
	t, err := template.ParseFiles("templates/" + tmpl)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
