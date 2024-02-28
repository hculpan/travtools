package entities

import (
	"encoding/json"
)

type Name struct {
	Text string `json:"Text"`
	Lang string `json:"Lang"`
}

type Sector struct {
	X            int    `json:"X"`
	Y            int    `json:"Y"`
	Milieu       string `json:"Milieu"`
	Abbreviation string `json:"Abbreviation"`
	Tags         string `json:"Tags"`
	Names        []Name `json:"Names"`
}

type sectorset struct {
	Sectors SectorList `json:"Sectors"`
}

type SectorList []Sector

func NewSectorList(data []byte) (SectorList, error) {
	result := sectorset{}

	if err := json.Unmarshal(data, &result); err != nil {
		return SectorList{}, err
	}

	return result.Sectors, nil
}
