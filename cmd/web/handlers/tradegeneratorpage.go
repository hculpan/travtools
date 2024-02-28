package handlers

import (
	"net/http"
	"text/template"

	"github.com/hculpan/travtools/cmd/web/embed"
)

func TradeGeneratorPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(embed.GetEmbeddedFS(), "templates/tradegenerator.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}
