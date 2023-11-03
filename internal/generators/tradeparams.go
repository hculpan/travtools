package generators

import "github.com/hculpan/travtools/internal/entities"

type TradeParams struct {
	StartPopulation string
	StartStarport   string
	StartAmberZone  bool
	StartRedZone    bool

	EndPopulation string
	EndStarport   string
	EndAmberZone  bool
	EndRedZone    bool

	BrokerEffect int
	StewardSkill int
	Distance     int
	Jumps        int

	PassengerTrades []entities.PassengerTrade
}

func NewTradeParams() *TradeParams {
	return &TradeParams{
		StartStarport: "X",
		EndStarport:   "X",
		BrokerEffect:  0,
		StewardSkill:  0,
		Distance:      1,
		Jumps:         1,
	}
}
