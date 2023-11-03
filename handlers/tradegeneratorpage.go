package handlers

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/hculpan/travtools/internal/generators"
)

func TradeGeneratorPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFS(content, "templates/tradegenerator.html")
		if err != nil {
			http.Error(w, "Failed to load template", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, generators.NewTradeParams())
	} else if r.Method == "POST" {
		r.ParseForm()

		tradeInfo := generators.NewTradeParams()

		tradeInfo.StartPopulation = r.FormValue("startPopulation")
		tradeInfo.StartStarport = r.FormValue("startStarport")
		tradeInfo.EndPopulation = r.FormValue("endPopulation")
		tradeInfo.EndStarport = r.FormValue("endStarport")

		tradeInfo.StartAmberZone = r.Form["startAlertLevel"][0] == "Amber Alert"
		tradeInfo.StartRedZone = r.Form["startAlertLevel"][0] == "Red Alert"
		tradeInfo.EndAmberZone = r.Form["endAlertLevel"][0] == "Amber Alert"
		tradeInfo.EndRedZone = r.Form["endAlertLevel"][0] == "Red Alert"

		brokerEffect, _ := strconv.ParseInt(r.FormValue("brokerEffect"), 10, 32)
		stewardSkill, _ := strconv.ParseInt(r.FormValue("stewardSkill"), 10, 32)
		jumps, _ := strconv.ParseInt(r.FormValue("jumps"), 10, 32)
		// distance, _ := strconv.ParseInt(r.FormValue("distance"), 10, 32)

		tradeInfo.BrokerEffect = int(brokerEffect)
		tradeInfo.StewardSkill = int(stewardSkill)
		tradeInfo.Jumps = int(jumps)

		trades, err := generators.GeneratePassengerTrade(tradeInfo)
		if err == nil {
			tradeInfo.PassengerTrades = trades
		}

		tmpl, err := template.ParseFS(content, "templates/tradegenerator.html")
		if err != nil {
			http.Error(w, "Failed to load template", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, tradeInfo)
	}
}
