package entities

import (
	"fmt"

	"github.com/hculpan/travtools/pkg/util"
)

const (
	MAJOR_CARGO      = "Major"
	MINOR_CARGO      = "Minor"
	INCIDENTAL_CARGO = "Incidental"
	MAIL             = "Mail"
)

type FreightLot struct {
	CargoType string
	LotCount  int
	Lots      []FreightTrade
}

func NewFreightLot(cargoType string) *FreightLot {
	result := FreightLot{
		CargoType: cargoType,
		LotCount:  util.RollDice(1, 0),
		Lots:      make([]FreightTrade, 0),
	}

	return &result
}

func (f *FreightLot) Description() string {
	if f.CargoType == MAIL {
		return "Mail: 5 tons (25,000 Cr)"
	}
	return fmt.Sprintf("%s: %d", f.CargoType, f.LotCount)
}
