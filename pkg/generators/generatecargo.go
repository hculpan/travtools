package generators

import (
	"strings"

	"github.com/hculpan/travtools/pkg/entities"
	"github.com/hculpan/travtools/pkg/util"
)

func GenerateCargoLots(tradeInfo *TradeParams) ([]entities.FreightLot, error) {
	result := []entities.FreightLot{}

	result = append(result, *generateCargoLots(entities.MAJOR_CARGO, tradeInfo))
	result = append(result, *generateCargoLots(entities.MINOR_CARGO, tradeInfo))
	result = append(result, *generateCargoLots(entities.INCIDENTAL_CARGO, tradeInfo))
	mailLot := generateCargoLots(entities.MAIL, tradeInfo)
	if mailLot != nil {
		result = append(result, *mailLot)
	}

	return result, nil
}

func generateCargoLots(cargoType string, tradeInfo *TradeParams) *entities.FreightLot {
	roll := util.RollDice(2, 0)

	if cargoType == entities.MAJOR_CARGO {
		roll -= 4
	} else if cargoType == entities.INCIDENTAL_CARGO {
		roll += 2
	}

	dm := modifiersForCargo(tradeInfo.StartPopulation, tradeInfo.StartStarport, tradeInfo.StartTechLevel, tradeInfo.StartAmberZone, tradeInfo.StartRedZone)
	dm += modifiersForCargo(tradeInfo.EndPopulation, tradeInfo.EndStarport, tradeInfo.EndTechLevel, tradeInfo.EndAmberZone, tradeInfo.EndRedZone)

	dm -= util.Max(tradeInfo.Jumps-1, 0)

	roll += dm

	if cargoType == entities.MAIL {
		return calcMail(dm)
	} else {
		return calcLots(cargoType, util.Max(roll, 0))
	}
}

func calcMail(dm int) *entities.FreightLot {
	if util.RollDice(2, dm) >= 12 {
		result := &entities.FreightLot{
			CargoType: entities.MAIL,
			LotCount:  1,
		}
		result.Lots = append(result.Lots, *entities.NewFreightTrade(entities.MAIL, entities.MAIL))
		return result
	} else {
		return nil
	}
}

func modifiersForCargo(population string, starport string, techLevel string, amberZone, redZone bool) int {
	result := 0

	population = strings.ToUpper(population)
	starport = strings.ToUpper(starport)

	switch population {
	case "0", "1":
		result -= 4
	case "6", "7":
		result += 1
	case "8", "9", "A", "B", "C", "D", "E", "F":
		result += 3
	}

	switch starport {
	case "A":
		result += 2
	case "B":
		result += 1
	case "E":
		result -= 1
	case "X":
		result -= 3
	}

	techLevelValue, _ := util.HexDigitToInt(techLevel)
	if techLevelValue <= 6 {
		result -= 1
	} else if techLevelValue >= 9 {
		result += 2
	}

	if amberZone {
		result -= 2
	} else if redZone {
		result -= 6
	}

	return result
}

func numberOfLots(roll int) int {
	switch roll {
	case 0, 1:
		return 0
	case 2, 3:
		return util.RollDice(1, 0)
	case 4, 5:
		return util.RollDice(2, 0)
	case 6, 7, 8:
		return util.RollDice(3, 0)
	case 9, 10, 11:
		return util.RollDice(4, 0)
	case 12, 13, 14:
		return util.RollDice(5, 0)
	case 15, 16:
		return util.RollDice(6, 0)
	case 17:
		return util.RollDice(7, 0)
	case 18:
		return util.RollDice(8, 0)
	case 19:
		return util.RollDice(9, 0)
	default:
		return util.RollDice(10, 0)
	}

}

func calcLots(cargoType string, roll int) *entities.FreightLot {
	return &entities.FreightLot{
		CargoType: cargoType,
		LotCount:  numberOfLots(roll),
	}
}
