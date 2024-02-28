package generators

import (
	"bufio"
	"fmt"
	"math/rand"
	"strings"

	"github.com/hculpan/travtools/pkg/embed"
	"github.com/hculpan/travtools/pkg/util"
)

var cargoNames []string = make([]string, 0)

func initCargoNames() error {
	if len(cargoNames) > 0 {
		return nil
	}

	cargos, err := embed.ReadDataFile("common_goods.txt")
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(cargos)))
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		cargoNames = append(cargoNames, line)
	}

	return nil
}

func GenerateCargoName(cargoType string) string {
	initCargoNames()

	if cargoType == "Mail" {
		return cargoType
	}

	return cargoNames[rand.Intn(len(cargoNames))]
}

func GenerateCargoTons(cargoType string) int {
	switch cargoType {
	case "Major":
		return util.RollDice(1, 0) * 10
	case "Minor":
		return util.RollDice(1, 0) * 5
	case "Mail":
		return 5
	default:
		return util.RollDice(1, 0)
	}
}
