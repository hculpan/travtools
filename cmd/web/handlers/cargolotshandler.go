package handlers

import (
	"log"
	"net/http"
	"sort"
	"strconv"
	"text/template"

	"github.com/hculpan/travtools/cmd/web/embed"
	"github.com/hculpan/travtools/pkg/cache"
	"github.com/hculpan/travtools/pkg/generators"
)

type cargoLot struct {
	Description string
	Tons        int
}

func CargoLotsHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	cargoType := r.FormValue("type")
	count64, err := strconv.ParseInt(r.FormValue("count"), 10, 64)
	if err != nil {
		log.Default().Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	count := int(count64)

	tmpl, err := template.ParseFS(embed.GetEmbeddedFS(), "templates/cargolots.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, &struct {
		CargoLots []cargoLot
	}{
		CargoLots: generateCargoLot(cargoType, count),
	})
}

func generateCargoLot(cargoType string, count int) []cargoLot {
	if cache.HasCache(cargoType) {
		pi, ok := cache.GetCache(cargoType).([]cargoLot)
		if ok {
			return pi
		}
	}

	result := make([]cargoLot, count)
	for i := range result {
		result[i].Description = generators.GenerateCargoName(cargoType)
		result[i].Tons = generators.GenerateCargoTons(cargoType)
	}

	cache.SetCache(cargoType, result)

	sort.Slice(result, func(i, j int) bool {
		return result[i].Tons > result[j].Tons
	})

	return result
}
