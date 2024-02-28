package handlers

import (
	"net/http"
	"text/template"

	"github.com/hculpan/travtools/cmd/web/embed"
	"github.com/hculpan/travtools/pkg/generators"
)

func PlanetGeneratorPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFS(embed.GetEmbeddedFS(), "templates/planetgenerator.html")
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

		system, _ := generators.GenerateNewPlanet(name, sector, hex, &generators.GeneratePlanetConfig{
			GenerateName:        false,
			VerifyUniqueName:    false,
			EnforceEnvTechLevel: true,
		})

		tmpl, err := template.ParseFS(embed.GetEmbeddedFS(), "templates/planetgenerator-post.html")
		if err != nil {
			http.Error(w, "Failed to load template", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, system)
	}
}
