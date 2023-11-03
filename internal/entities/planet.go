package entities

import "fmt"

type Uwp struct {
	Starport     string `json:"starport"`
	Size         int    `json:"size"`
	Atmosphere   int    `json:"atmosphere"`
	Hydrographic int    `json:"hydrographic"`
	Population   int    `json:"population"`
	Government   int    `json:"government"`
	LawLevel     int    `json:"law_level"`
	TechLevel    int    `json:"tech_level"`
}

type Planet struct {
	Name   string `json:"name"`
	Hex    string `json:"hex"`
	Sector string `json:"sector"`

	Highport       bool   `json:"highport"`
	HighportString string `json:"-"`

	NavalBase string `json:"naval-base"`
	ScoutBase string `json:"scout-base"`

	Uwp       Uwp    `json:"uwp"`
	UwpString string `json:"-"`
}

func (u *Uwp) String() string {
	return fmt.Sprintf("%1s%1X%1X%1X%1X%1X%1X-%1X", u.Starport, u.Size, u.Atmosphere, u.Hydrographic, u.Population, u.Government, u.LawLevel, u.TechLevel)
}
