package main

import (
	"net/http"

	"github.com/hculpan/travtools/cmd/web/embed"
	"github.com/hculpan/travtools/cmd/web/handlers"
)

func routes() {
	// Serve static files from "assets" directory
	http.Handle("/assets/", http.FileServer(http.FS(embed.GetEmbeddedFS())))

	http.HandleFunc("/", handlers.HomePage)
	http.HandleFunc("/message", handlers.MessagePage)
	http.HandleFunc("/planet-generator", handlers.PlanetGeneratorPage)

	http.HandleFunc("/trade-generator", handlers.TradeGeneratorPage)
	http.HandleFunc("/search-systems", handlers.SearchSystemsHandler)

	http.HandleFunc("/passengers", handlers.PassengersHandler)
	http.HandleFunc("/cargolots", handlers.CargoLotsHandler)

	http.HandleFunc("/generate-trade", handlers.GenerateTradePage)

	http.HandleFunc("/spec-trade-purchase", handlers.SpecTradePurchaseHandler)
	http.HandleFunc("/spec-trade-sell", handlers.SpecTradeSellHandler)
}
