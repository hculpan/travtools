package generators

import (
	"strings"

	"github.com/hculpan/travtools/pkg/entities"
	"github.com/hculpan/travtools/pkg/util"
)

func modifiersForPassenger(population string, starport string, amberZone, redZone bool) int {
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

	if amberZone {
		result += 1
	} else if redZone {
		result -= 4
	}

	return result
}

func generatePassengers(passageClass string, startWorldPop, endWorldPop string, startStarport, endStarport string, startAmberZone, startRedZone, endAmberZone, endRedZone bool, distanceToDest int) int {
	result := util.RollDice(2, 0)

	if passageClass == entities.HIGH {
		result -= 4
	}

	result += modifiersForPassenger(startWorldPop, startStarport, startAmberZone, startRedZone)
	result += modifiersForPassenger(endWorldPop, endStarport, endAmberZone, endRedZone)

	result -= util.Max(distanceToDest-1, 0)

	result = calcPassengers(result)

	return util.Max(result, 0)
}

func calcPassengers(dieRoll int) int {
	if dieRoll <= 1 {
		return 0
	}

	switch dieRoll {
	case 2, 3:
		return util.RollDice(1, 0)
	case 4, 5, 6:
		return util.RollDice(2, 0)
	case 7, 8, 9, 10:
		return util.RollDice(3, 0)
	case 11, 12, 13:
		return util.RollDice(4, 0)
	case 14, 15:
		return util.RollDice(5, 0)
	case 16:
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

func GeneratePassengerTrade(tradeParams *TradeParams) ([]entities.PassengerTrade, error) {
	result := []entities.PassengerTrade{}

	result = append(result, entities.PassengerTrade{
		Passage: entities.HIGH,
		Count: generatePassengers(
			entities.HIGH,
			tradeParams.StartPopulation,
			tradeParams.EndPopulation,
			tradeParams.StartStarport,
			tradeParams.EndStarport,
			tradeParams.StartAmberZone,
			tradeParams.StartRedZone,
			tradeParams.EndAmberZone,
			tradeParams.EndRedZone,
			tradeParams.Jumps,
		)})
	result = append(result, entities.PassengerTrade{
		Passage: entities.MIDDLE,
		Count: generatePassengers(
			entities.MIDDLE,
			tradeParams.StartPopulation,
			tradeParams.EndPopulation,
			tradeParams.StartStarport,
			tradeParams.EndStarport,
			tradeParams.StartAmberZone,
			tradeParams.StartRedZone,
			tradeParams.EndAmberZone,
			tradeParams.EndRedZone,
			tradeParams.Jumps,
		)})
	result = append(result, entities.PassengerTrade{
		Passage: entities.BASIC,
		Count: generatePassengers(
			entities.BASIC,
			tradeParams.StartPopulation,
			tradeParams.EndPopulation,
			tradeParams.StartStarport,
			tradeParams.EndStarport,
			tradeParams.StartAmberZone,
			tradeParams.StartRedZone,
			tradeParams.EndAmberZone,
			tradeParams.EndRedZone,
			tradeParams.Jumps,
		)})
	result = append(result, entities.PassengerTrade{
		Passage: entities.LOW,
		Count: generatePassengers(
			entities.LOW,
			tradeParams.StartPopulation,
			tradeParams.EndPopulation,
			tradeParams.StartStarport,
			tradeParams.EndStarport,
			tradeParams.StartAmberZone,
			tradeParams.StartRedZone,
			tradeParams.EndAmberZone,
			tradeParams.EndRedZone,
			tradeParams.Jumps,
		)})

	return result, nil
}
