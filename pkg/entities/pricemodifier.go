package entities

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hculpan/travtools/pkg/embed"
)

type PriceModifier struct {
	Result           int     `json:"result"`
	PurchaseModifier float32 `json:"purchase_mod"`
	SaleModifier     float32 `json:"sale_mod"`
}

type PriceModifiers []PriceModifier

var priceMods PriceModifiers

func init() {
	data, err := embed.ReadDataFile("specpricemod.json")
	if err != nil {
		log.Fatal(err)
	}

	priceModifiers, err := unmarshalPriceModifiers(data)
	if err != nil {
		log.Fatal(err)
	}
	priceMods = priceModifiers
}

func GetPriceModifier(num int) (PriceModifier, error) {
	if num < -3 {
		num = -3
	} else if num > 25 {
		num = 25
	}

	for i := range priceMods {
		if priceMods[i].Result == num {
			return priceMods[i], nil
		}
	}

	return PriceModifier{}, fmt.Errorf("price modifier not found for %d", num)
}

func unmarshalPriceModifiers(jsonData []byte) (PriceModifiers, error) {
	var result PriceModifiers
	err := json.Unmarshal(jsonData, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
