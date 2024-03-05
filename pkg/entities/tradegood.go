package entities

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"slices"
	"strconv"
	"strings"

	"github.com/hculpan/travtools/pkg/embed"
)

var CodeOffset = map[string]int{}
var allTradeGoods TradeGoods

type NoDmError struct {
	code string
}

type TradeGoods []TradeGood

type TonsType struct {
	NumDice  int `json:"num_dice"`
	Modifier int `json:"modifier"`
}

type TradeGood struct {
	idInt         int
	Id            string         `json:"id"`
	Type          string         `json:"type"`
	Availability  []string       `json:"availability"`
	Tons          TonsType       `json:"tons"`
	BasePrice     int            `json:"base_price"`
	PurchaseDm    map[string]int `json:"purchase_dms"`
	SaleDm        map[string]int `json:"sale_dms"`
	PurchasePrice int            `json:"-"`
	SalePrice     int            `json:"-"`
	AvailableTons int            `json:"-"`
	ReasonAdded   []string       `json:"-"`
}

func NewTons(number, modifier string) TonsType {
	n, err := strconv.Atoi(number)
	if err != nil {
		n = 0
	}
	m, err := strconv.Atoi(modifier)
	if err != nil {
		m = 1
	}
	return TonsType{
		NumDice:  n,
		Modifier: m,
	}
}

func NewTradeGood(id, typeName, numDice, modifier, basePrice string) *TradeGood {
	basePrice = strings.Replace(basePrice, ",", "", -1)
	bp, err := strconv.Atoi(basePrice)
	if err != nil {
		bp = 0
	}

	result := &TradeGood{
		Id:          id,
		Type:        typeName,
		Tons:        NewTons(numDice, modifier),
		BasePrice:   bp,
		ReasonAdded: []string{},
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Default().Printf("invalid tradegood id %q", result.Id)
		idInt = 0
	}
	result.idInt = idInt

	result.Availability = []string{}
	result.PurchaseDm = map[string]int{}
	result.SaleDm = map[string]int{}

	return result
}

func init() {
	codes := []string{"Ag", "As", "Ba", "De", "Fl", "Ga", "Hi", "Ht", "Ie", "In", "Lo", "Lt", "Na", "Ni", "Po", "Ri", "Va", "Wa"}
	for i, code := range codes {
		CodeOffset[code] = i
	}

	data, err := embed.ReadDataFile("spec-trade-data.json")
	if err != nil {
		log.Fatal(err)
	}

	tradeGoods, err := unmarshalTradeGoods(data)
	if err != nil {
		log.Fatal(err)
	}
	allTradeGoods = tradeGoods

	for i := range allTradeGoods {
		idInt, err := strconv.Atoi(allTradeGoods[i].Id)
		if err != nil {
			log.Default().Printf("invalid tradegood id %q", allTradeGoods[i].Id)
			idInt = 0
		}
		allTradeGoods[i].idInt = idInt

	}
}

func unmarshalTradeGoods(jsonData []byte) (TradeGoods, error) {
	var goods TradeGoods
	err := json.Unmarshal(jsonData, &goods)
	if err != nil {
		return nil, err
	}

	return goods, nil
}

func (t *TonsType) String() string {
	if t.Modifier != 1 {
		return fmt.Sprintf("%dD * %d", t.NumDice, t.Modifier)
	}
	return fmt.Sprintf("%dD", t.NumDice)
}

func (t *TradeGood) String() string {
	if t.IsCommon() {
		return fmt.Sprintf("Id: %s, Type: %s, Tons: %s, Base Price: %d, Available: All, Purchase DMs: %s, Sale DMs: %s", t.Id, t.Type, t.Tons.String(), t.BasePrice, t.PurchaseDmsString(), t.SaleDmsString())
	}
	return fmt.Sprintf("Id: %s, Type: %s, Tons: %s, Base Price: %d, Available: %s, Purchase DMs: %s, Sale DMs: %s", t.Id, t.Type, t.Tons.String(), t.BasePrice, strings.Join(t.Availability, ", "), t.PurchaseDmsString(), t.SaleDmsString())
}

func (t *TradeGood) PurchaseDmsString() string {
	result := []string{}
	for code, mod := range t.PurchaseDm {
		result = append(result, fmt.Sprintf("%s:%d", code, mod))
	}

	return strings.Join(result, ", ")
}

func (t *TradeGood) SaleDmsString() string {
	result := []string{}
	for code, mod := range t.SaleDm {
		result = append(result, fmt.Sprintf("%s:%d", code, mod))
	}

	return strings.Join(result, ", ")
}

func (t *TradeGood) IsAvailable(code string) bool {
	for _, s := range t.Availability {
		if s == code {
			return true
		}
	}

	return false
}

func (t *TradeGood) IsIllegal() bool {
	return t.Id[0] == '6'
}

func (t *TradeGood) IsCommon() bool {
	return len(t.Availability) == 18
}

func (t *TradeGood) GetPurchaseDm(code string) (int, error) {
	for k, v := range t.PurchaseDm {
		if k == code {
			return v, nil
		}
	}

	return 0, NewNoDmError(code)
}

func (t *TradeGood) CalculateAvailability(mod int) {
	t.AvailableTons = (rollD6(t.Tons.NumDice) + mod) * t.Tons.Modifier
}

func rollD6(num int) int {
	result := 0
	for i := 0; i < num; i++ {
		result += rand.IntN(6) + 1
	}
	return result
}

func (t *TradeGood) SetSalePrice(popSize int, tradeCodes []string, brokerSkill int) {
	result := rollD6(3)
	result -= brokerSkill
	result += 2 // local buyer's broker skill, assumed

	purchaseDm := 0
	for _, code := range tradeCodes {
		dm, _ := t.GetPurchaseDm(code)
		if dm > purchaseDm {
			purchaseDm = dm
		}
	}

	saleDm := 0
	for _, code := range tradeCodes {
		dm, _ := t.GetSaleDm(code)
		if dm > saleDm {
			saleDm = dm
		}
	}

	result -= purchaseDm
	result += saleDm

	pm, err := GetPriceModifier(result)
	if err != nil {
		t.SalePrice = 0
	}

	t.SalePrice = int(float32(t.BasePrice) * pm.SaleModifier)
}

func (t *TradeGood) SetPurchasePrice(popSize int, tradeCodes []string, brokerSkill int) {
	result := rollD6(3)
	result += brokerSkill
	result -= 2 // local seller's broker skill, assumed

	purchaseDm := 0
	for _, code := range tradeCodes {
		dm, _ := t.GetPurchaseDm(code)
		if dm > purchaseDm {
			purchaseDm = dm
		}
	}

	saleDm := 0
	for _, code := range tradeCodes {
		dm, _ := t.GetSaleDm(code)
		if dm > saleDm {
			saleDm = dm
		}
	}

	result += purchaseDm
	result -= saleDm

	pm, err := GetPriceModifier(result)
	if err != nil {
		t.PurchasePrice = 0
	}

	t.PurchasePrice = int(float32(t.BasePrice) * pm.PurchaseModifier)
}

func (t *TradeGood) GetSaleDm(code string) (int, error) {
	for k, v := range t.SaleDm {
		if k == code {
			return v, nil
		}
	}

	return 0, NewNoDmError(code)
}

func GetAllTradeGoods() TradeGoods {
	return allTradeGoods
}

func GetAvailableForCodes(includeIllegal bool, includeAdditional int, availabilityMod int, codes ...string) []TradeGood {
	foundCodes := map[string]TradeGood{}

	for _, code := range codes {
		for _, tradeGood := range allTradeGoods {
			if (tradeGood.IsAvailable(code) || tradeGood.IsCommon()) && (includeIllegal || !tradeGood.IsIllegal()) {
				oldTradeGood, ok := foundCodes[tradeGood.Id]
				if ok {
					tradeGood.CalculateAvailability(availabilityMod)
					if !tradeGood.IsCommon() {
						oldTradeGood.ReasonAdded = append(oldTradeGood.ReasonAdded, code)
					}
					oldTradeGood.AvailableTons += tradeGood.AvailableTons
					foundCodes[tradeGood.Id] = oldTradeGood
				} else {
					tradeGood.CalculateAvailability(availabilityMod)
					if tradeGood.IsCommon() {
						tradeGood.ReasonAdded = []string{"common"}
					} else {
						tradeGood.ReasonAdded = append(tradeGood.ReasonAdded, code)
					}
					foundCodes[tradeGood.Id] = tradeGood
				}
			}
		}
	}

	for range includeAdditional {
		var index int
		if includeIllegal {
			index = rand.IntN(len(allTradeGoods))
		} else {
			index = rand.IntN(len(allTradeGoods) - 6)
		}
		tradeGood := allTradeGoods[index]
		tradeGood.CalculateAvailability(availabilityMod)
		oldTradeGood, ok := foundCodes[tradeGood.Id]
		if ok {
			oldTradeGood.AvailableTons += tradeGood.AvailableTons
			foundCodes[tradeGood.Id] = oldTradeGood
		} else {
			foundCodes[tradeGood.Id] = tradeGood
		}
	}

	result := []TradeGood{}
	for _, v := range foundCodes {
		result = append(result, v)
	}

	slices.SortFunc(result, func(i, j TradeGood) int {
		return i.idInt - j.idInt
	})

	return result
}

func NewNoDmError(code string) NoDmError {
	return NoDmError{code: code}
}

func (n NoDmError) Error() string {
	return "no DM found "
}

func IsValidTradeCode(code string) bool {
	_, ok := CodeOffset[code]
	return ok
}

func GetPurchaseResult(tradeGood TradeGood, codes []string) int {
	result := 0
	for _, c := range codes {
		n, err := tradeGood.GetPurchaseDm(c)
		if err == nil && n > result {
			result = n
		}
	}

	return result
}

func GetSaleResult(tradeGood TradeGood, codes []string) int {
	result := 0
	for _, c := range codes {
		n, err := tradeGood.GetSaleDm(c)
		if err == nil && n > result {
			result = n
		}
	}

	return result
}

func GetTradeGoodById(id string) (TradeGood, error) {
	for _, tradeGood := range allTradeGoods {
		if tradeGood.Id == id {
			return tradeGood, nil
		}
	}

	return TradeGood{}, fmt.Errorf("no trade good with id %q found", id)
}

func GetTradeGoodsByIds(ids []string) TradeGoods {
	result := TradeGoods{}

	for _, id := range ids {
		tg, err := GetTradeGoodById(id)
		if err == nil {
			result = append(result, tg)
		}
	}

	return result
}
