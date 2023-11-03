package handlers

import (
	"embed"
	"net/http"
	"text/template"

	"github.com/hculpan/travtools/internal/generators"
)

func SetContent(value *embed.FS) {
	content = value
}

var content *embed.FS

func PlanetGeneratorPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFS(content, "templates/planetgenerator.html")
		if err != nil {
			http.Error(w, "Failed to load template", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseForm()

		name := r.FormValue("planetName")
		sector := r.FormValue("sector")
		hex := r.FormValue("hex")

		system, err := generators.GenerateNewPlanet(name, sector, hex, &generators.GeneratePlanetConfig{
			GenerateName:        false,
			VerifyUniqueName:    false,
			EnforceEnvTechLevel: true,
		})

		tmpl, err := template.ParseFS(content, "templates/planetgenerator-post.html")
		if err != nil {
			http.Error(w, "Failed to load template", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, system)
	}
}
