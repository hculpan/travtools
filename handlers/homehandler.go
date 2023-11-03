package handlers

import (
	"net/http"
	"text/template"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	//	tmpl, err := template.ParseFiles("templates/index.html")
	tmpl, err := template.ParseFS(content, "templates/index.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func MessagePage(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Message": "Welcome to Traveller Tools!",
	}

	tmpl, err := template.ParseFS(content, "templates/message.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}
