package handlers

import (
	"net/http"
	"text/template"

	"github.com/hculpan/travtools/cmd/web/embed"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(embed.GetEmbeddedFS(), "templates/index.html")
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

	tmpl, err := template.ParseFS(embed.GetEmbeddedFS(), "templates/message.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}
