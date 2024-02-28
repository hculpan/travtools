package handlers

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand/v2"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/a-h/templ"
	"github.com/hculpan/travtools/cmd/web/templates"
	"github.com/hculpan/travtools/pkg/entities"
)

type PurchaseOptions struct {
	PopulationSize int
	TradeCodes     []string
	IllegalGoods   bool
	BrokerSkill    int
}

func SpecTradePurchaseHandler(w http.ResponseWriter, r *http.Request) {
	step := r.URL.Query().Get("step")
	if step == "" {
		component := templates.FindSupplier()
		component.Render(context.Background(), w)
		return
	}

	stepNum, err := strconv.Atoi(step)
	if err != nil {
		component := templates.FindSupplier()
		component.Render(context.Background(), w)
		return
	}

	switch stepNum {
	case 1:
		component := templates.FindSupplier()
		component.Render(context.Background(), w)
	case 2:
		component := templates.DetermineGoods()
		component.Render(context.Background(), w)
	case 3:
		po, err := NewPurchaseOptions(r)
		var component templ.Component
		if err != nil {
			component = templates.Message(err.Error(), true)
		} else {
			availMod := 0
			if po.PopulationSize <= 3 {
				availMod = -3
			} else if po.PopulationSize >= 9 {
				availMod = 3
			}
			tradeGoods := entities.GetAvailableForCodes(po.IllegalGoods, po.PopulationSize, availMod, po.TradeCodes...)
			tradeGoods, err = calculatePurchasePrice(po, tradeGoods)
			if err != nil {
				component = templates.Message(err.Error(), true)
			} else {
				component = templates.DeterminePurchasePrice(tradeGoods)
			}
		}
		component.Render(context.Background(), w)
	default:
		component := templates.FindSupplier()
		component.Render(context.Background(), w)
	}
}

func rollD6(num int) int {
	result := 0
	for i := 0; i < num; i++ {
		result += rand.IntN(6) + 1
	}
	return result
}

func calculatePurchasePrice(po PurchaseOptions, tradeGoods entities.TradeGoods) (entities.TradeGoods, error) {
	for i := range tradeGoods {
		result := rollD6(3)
		result += po.BrokerSkill
		result += entities.GetPurchaseResult(tradeGoods[i], po.TradeCodes)
		result -= entities.GetSaleResult(tradeGoods[i], po.TradeCodes)
		result -= 2 // default seller's broker skill

		priceMod, err := entities.GetPriceModifier(result)
		if err != nil {
			log.Default().Printf("Price mod = %d, %f", priceMod.Result, priceMod.PurchaseModifier)
			return tradeGoods, fmt.Errorf("error finding price modifier: %w", err)
		}
		tradeGoods[i].PurchasePrice = int(float32(tradeGoods[i].BasePrice) * priceMod.PurchaseModifier)
	}

	return tradeGoods, nil
}

func NewPurchaseOptions(r *http.Request) (PurchaseOptions, error) {
	pop := r.URL.Query().Get("population")
	tradeCodes := r.URL.Query().Get("trade")
	legalGoods := strings.Trim(r.URL.Query().Get("goodstype"), " ") == "legal-goods"
	broker := r.URL.Query().Get("broker")

	brokerSkill, err := strconv.Atoi(broker)
	if err != nil {
		brokerSkill = 0
	}

	popMatch, _ := regexp.MatchString("[0-9a-fA-F]", pop)
	if !popMatch {
		return PurchaseOptions{}, errors.New("population must be a single-character code between 0-9 or A-F")
	}
	population, err := strconv.ParseInt(pop, 16, 8)
	if err != nil {
		return PurchaseOptions{}, fmt.Errorf("invalid population %q: %w", pop, err)
	}

	codes := strings.Split(tradeCodes, ",")
	for i := range codes {
		codes[i] = strings.Trim(codes[i], " ")
		if len(strings.Trim(codes[i], " ")) != 2 {
			return PurchaseOptions{}, fmt.Errorf("invalid code %q", codes[i])
		}
	}

	po := PurchaseOptions{
		PopulationSize: int(population),
		TradeCodes:     codes,
		IllegalGoods:   !legalGoods,
		BrokerSkill:    brokerSkill,
	}

	return po, nil
}
