package generators

import (
	"fmt"
	"log"
	"math/rand"
)

const (
	MALE int = iota
	FEMALE
	EITHER
)

var firstFemaleChain *Chain
var firstMaleChain *Chain
var lastNameChain *Chain

func initNameGenerators() error {
	var err error
	firstFemaleChain, err = NewChainFromFile("female_first.txt", 2, 12)
	if err != nil {
		return err
	}
	firstFemaleChain.Build()

	firstMaleChain, err = NewChainFromFile("male_first.txt", 2, 12)
	if err != nil {
		return err
	}
	firstMaleChain.Build()

	lastNameChain, err = NewChainFromFile("last.txt", 3, 15)
	if err != nil {
		return err
	}
	lastNameChain.Build()

	return nil
}

func GenerateName(nameType int) string {
	if firstFemaleChain == nil {
		err := initNameGenerators()
		if err != nil {
			log.Println(err)
			return ""
		}
	}

	if nameType == FEMALE || (nameType == EITHER && rand.Intn(100) < 51) {
		return fmt.Sprintf("%s %s", firstFemaleChain.GenerateName(), lastNameChain.GenerateName())
	} else {
		return fmt.Sprintf("%s %s", firstMaleChain.GenerateName(), lastNameChain.GenerateName())
	}
}
