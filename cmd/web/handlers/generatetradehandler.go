package handlers

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/hculpan/travtools/cmd/web/embed"
	"github.com/hculpan/travtools/pkg/cache"
	"github.com/hculpan/travtools/pkg/generators"
	"github.com/hculpan/travtools/pkg/util"
)

func GenerateTradePage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	tradeInfo := generators.NewTradeParams()

	startUwp := util.CleanUwp(r.FormValue("startUwp"))
	endUwp := util.CleanUwp(r.FormValue("endUwp"))

	tradeInfo.StartPopulation = string(startUwp[4])
	tradeInfo.StartStarport = string(startUwp[0])
	tradeInfo.StartTechLevel = string(startUwp[7])
	tradeInfo.EndPopulation = string(endUwp[4])
	tradeInfo.EndStarport = string(endUwp[0])
	tradeInfo.EndTechLevel = string(endUwp[7])

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

	cache.ResetCache()

	{
		trades, err := generators.GeneratePassengerTrade(tradeInfo)
		if err == nil {
			tradeInfo.PassengerTrades = trades
		}
	}

	{
		trades, err := generators.GenerateCargoLots(tradeInfo)
		if err == nil {
			tradeInfo.CargoLots = trades
		}
	}

	tmpl, err := template.ParseFS(embed.GetEmbeddedFS(), "templates/generatetrade.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, tradeInfo)
}
