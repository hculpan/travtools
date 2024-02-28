package generators

import (
	"bufio"
	"fmt"
	"math/rand"
	"strings"

	"github.com/hculpan/travtools/pkg/embed"
)

var passengerAspects []string = make([]string, 0)

func initPassengerAspects() error {
	if len(passengerAspects) > 0 {
		return nil
	}

	aspects, err := embed.ReadDataFile("passenger_aspects.txt")
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(aspects)))
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		passengerAspects = append(passengerAspects, line)
	}

	return nil
}

func option(choices []string, weights []int, max int) string {
	selected := rand.Intn(max) + 1

	for i := range weights {
		selected -= weights[i]
		if selected <= 0 {
			return choices[i]
		}
	}

	return "error in option()"
}

func GeneratePassengerAspect() string {
	if err := initPassengerAspects(); err != nil {
		fmt.Println(err)
		return err.Error()
	}

	aspect := passengerAspects[rand.Intn(len(passengerAspects))]

	return aspect
}
