package handlers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/hculpan/travtools/cmd/web/templates"
	"github.com/hculpan/travtools/pkg/entities"
)

func SpecTradeSellHandler(w http.ResponseWriter, r *http.Request) {
	step := r.URL.Query().Get("step")
	if step == "" {
		component := templates.FindBuyer()
		component.Render(context.Background(), w)
		return
	}

	stepNum, err := strconv.Atoi(step)
	if err != nil {
		component := templates.FindBuyer()
		component.Render(context.Background(), w)
		return
	}

	switch stepNum {
	case 2:
		component := templates.DetermineSellOptions()
		component.Render(context.Background(), w)
	case 3:
		pop := r.URL.Query().Get("population")
		tradeCodes := r.URL.Query().Get("trade")
		broker := r.URL.Query().Get("broker")

		component := templates.PickToSell(entities.GetAllTradeGoods(), pop, tradeCodes, broker)
		component.Render(context.Background(), w)
	case 4:
		tradeGoods, err := getSelectedTradeGoods(r)
		if err != nil {
			component := templates.Message(err.Error(), true)
			component.Render(context.Background(), w)
			return
		}

		component := templates.DetermineSalePrice(tradeGoods)
		component.Render(context.Background(), w)
	default: // also case 1
		component := templates.FindBuyer()
		component.Render(context.Background(), w)
	}
}

func getSelectedTradeGoods(r *http.Request) (entities.TradeGoods, error) {
	pop := r.URL.Query().Get("population")
	popMatch, _ := regexp.MatchString("[0-9a-fA-F]", pop)
	if !popMatch {
		return entities.TradeGoods{}, errors.New("population must be a single-character code between 0-9 or A-F")
	}
	popSize, err := strconv.ParseInt(pop, 16, 8)
	if err != nil {
		return entities.TradeGoods{}, fmt.Errorf("invalid population %q: %w", pop, err)
	}

	broker := r.URL.Query().Get("broker")
	brokerSkill, err := strconv.Atoi(broker)
	if err != nil {
		return entities.TradeGoods{}, fmt.Errorf("invalid broker skill %q", broker)
	}

	tradeCodes := r.URL.Query().Get("trade")
	tCodes := strings.Split(tradeCodes, ",")

	tradeGoodIds := []string{}

	values := r.URL.Query()
	for k := range values {
		_, err := strconv.Atoi(k)
		if err == nil {
			tradeGoodIds = append(tradeGoodIds, k)
		}
	}

	tradeGoods := entities.GetTradeGoodsByIds(tradeGoodIds)
	for i := range tradeGoods {
		tradeGoods[i].SetSalePrice(int(popSize), tCodes, brokerSkill)
	}

	return tradeGoods, nil
}
