package handlers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/hculpan/travtools/cmd/web/embed"
	"github.com/hculpan/travtools/pkg/cache"
	"github.com/hculpan/travtools/pkg/generators"
)

type passengerInfo struct {
	Name   string
	Aspect string
}

func PassengersHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	passage := r.FormValue("passage")
	count64, err := strconv.ParseInt(r.FormValue("count"), 10, 64)
	if err != nil {
		log.Default().Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	count := int(count64)

	tmpl, err := template.ParseFS(embed.GetEmbeddedFS(), "templates/passengers.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, struct {
		Passage    string
		Count      int
		Passengers []passengerInfo
	}{
		Passage:    passage,
		Count:      int(count),
		Passengers: generatePassengerInfo(passage, count),
	})
}

func generatePassengerInfo(passage string, count int) []passengerInfo {
	if cache.HasCache(passage) {
		pi, ok := cache.GetCache(passage).([]passengerInfo)
		if ok {
			return pi
		}
	}

	result := make([]passengerInfo, count)
	for i := range result {
		result[i].Name = generators.GenerateName(generators.EITHER)
		result[i].Aspect = generators.GeneratePassengerAspect()
	}

	cache.SetCache(passage, result)

	return result
}
