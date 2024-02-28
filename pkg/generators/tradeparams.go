package generators

import "github.com/hculpan/travtools/pkg/entities"

type TradeParams struct {
	StartPopulation string
	StartStarport   string
	StartTechLevel  string
	StartAmberZone  bool
	StartRedZone    bool

	EndPopulation string
	EndStarport   string
	EndTechLevel  string
	EndAmberZone  bool
	EndRedZone    bool

	BrokerEffect int
	StewardSkill int
	Jumps        int

	PassengerTrades   []entities.PassengerTrade
	CargoLots         []entities.FreightLot
	SpeculativeTrades []entities.SpeculativeTrade
}

func NewTradeParams() *TradeParams {
	return &TradeParams{
		StartStarport: "X",
		EndStarport:   "X",
		BrokerEffect:  0,
		StewardSkill:  0,
		Jumps:         1,
	}
}
