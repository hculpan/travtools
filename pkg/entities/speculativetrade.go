package entities

import "fmt"

type SpeculativeTrade struct {
	Cargo string
	Tons  int
}

func (s *SpeculativeTrade) Description() string {
	return fmt.Sprintf("%s: %d", s.Cargo, s.Tons)
}
