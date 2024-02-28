package entities

import (
	"fmt"

	"github.com/hculpan/travtools/pkg/util"
)

type FreightTrade struct {
	Cargo string
	Tons  int
}

func NewFreightTrade(cargoType, cargo string) *FreightTrade {
	result := FreightTrade{
		Cargo: cargo,
		Tons:  util.RollDice(1, 0),
	}

	switch cargoType {
	case MAJOR_CARGO:
		result.Tons *= 10
	case MINOR_CARGO:
		result.Tons *= 5
	case MAIL:
		result.Tons = 5
	}

	return &result
}

func (f *FreightTrade) Description() string {
	return fmt.Sprintf("%s: %d", f.Cargo, f.Tons)
}
