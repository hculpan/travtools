package entities

import "fmt"

const (
	HIGH   = "High"
	MIDDLE = "Middle"
	BASIC  = "Basic"
	LOW    = "Low"
)

type PassengerTrade struct {
	Passage string
	Count   int
}

func (p *PassengerTrade) Description() string {
	return fmt.Sprintf("%s: %d", p.Passage, p.Count)
}
