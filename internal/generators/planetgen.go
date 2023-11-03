package generators

import (
	"strconv"

	"github.com/hculpan/travtools/internal/entities"
	"github.com/hculpan/travtools/internal/util"
)

type GeneratePlanetConfig struct {
	GenerateName        bool
	VerifyUniqueName    bool
	EnforceEnvTechLevel bool
}

func GenerateNewPlanet(name, sector, hex string, config *GeneratePlanetConfig) (*entities.Planet, error) {
	result := &entities.Planet{
		Sector: sector,
		Hex:    hex,
		Name:   name,
		Uwp:    entities.Uwp{},
	}

	if config.GenerateName {
		result.Name = "Planet X"
	}

	result.Uwp.Size = generateSize(result)
	result.Uwp.Atmosphere = generateAtmosphere(result)
	result.Uwp.Hydrographic = generateHydrographic(result)
	result.Uwp.Population = generatePopulation(result)
	result.Uwp.Government = generateGovernment(result)
	result.Uwp.LawLevel = generateLawLevel(result)
	result.Uwp.Starport = generateStarport(result)
	result.Uwp.TechLevel = generateTechLevel(result)

	result.Highport = hasHighport(result)
	if result.Highport {
		result.HighportString = "Yes"
	} else if result.Uwp.Starport == "X" {
		result.HighportString = "--"
	} else {
		result.HighportString = "No"
	}

	if config.EnforceEnvTechLevel {
		result.Uwp.TechLevel = limitTechLevel(result)
	}

	result.UwpString = result.Uwp.String()

	return result, nil
}

func hexDigitToInt(digit string) int {
	switch digit {
	case "A":
		return 10
	case "B":
		return 11
	case "C":
		return 12
	case "D":
		return 13
	case "E":
		return 14
	case "F":
		return 15
	default:
		v, _ := strconv.ParseInt(digit, 10, 64)
		return int(v)
	}
}

func hasHighport(planet *entities.Planet) bool {
	result := util.RollDice(2, 0)

	if planet.Uwp.Population >= 9 && planet.Uwp.Population <= 11 {
		result += 1
	} else if planet.Uwp.Population >= 12 {
		result += 2
	}

	switch planet.Uwp.Starport {
	case "A":
		return result >= 6
	case "B":
		return result >= 8
	case "C":
		return result >= 10
	case "D":
		return result >= 12
	default:
		return false
	}
}

func generateSize(planet *entities.Planet) int {
	return util.RollDice(2, -2)
}

func generateAtmosphere(planet *entities.Planet) int {
	i := util.RollDice(2, -7) + planet.Uwp.Size
	if i < 0 {
		i = 0
	}
	return i
}

func generateHydrographic(planet *entities.Planet) int {
	i := util.RollDice(2, -7) + planet.Uwp.Atmosphere
	if planet.Uwp.Size == 0 || planet.Uwp.Size == 1 {
		i = 0
	}
	if planet.Uwp.Atmosphere == 0 || planet.Uwp.Atmosphere == 1 || planet.Uwp.Atmosphere > 9 {
		i -= 4
	}

	if i < 0 {
		i = 0
	}

	return i
}

func generatePopulation(planet *entities.Planet) int {
	i := util.RollDice(2, -2)

	if i < 0 {
		i = 0
	}
	return i
}

func generateGovernment(planet *entities.Planet) int {
	i := util.RollDice(2, -7) + planet.Uwp.Population

	if i < 0 || planet.Uwp.Population == 0 {
		i = 0
	}
	return i
}

func generateLawLevel(planet *entities.Planet) int {
	i := util.RollDice(2, -7) + planet.Uwp.Government

	if i < 0 || planet.Uwp.Population == 0 {
		i = 0
	}
	return i
}

func generateStarport(planet *entities.Planet) string {
	pop := planet.Uwp.Population
	i := util.RollDice(2, 0)

	switch pop {
	case 8, 9:
		i += 1
	case 10, 11, 12, 13, 14, 15:
		i += 2
	case 3, 4:
		i -= 1
	case 0, 1, 2:
		i -= 2
	}

	if i < 0 {
		i = 0
	}
	switch i {
	case 0, 1, 2:
		return "X"
	case 3, 4:
		return "E"
	case 5, 6:
		return "D"
	case 7, 8:
		return "C"
	case 9, 10:
		return "B"
	default:
		return "A"
	}
}

func generateTechLevel(planet *entities.Planet) int {
	if planet.Uwp.Population == 0 {
		return 0
	}

	i := util.RollDice(1, 0)
	i += bonusForStarport(planet)
	i += bonusForSize(planet)
	i += bonusForAtmosphere(planet)
	i += bonusForHydrographics(planet)
	i += bonusForPopulation(planet)
	i += bonusForGovernment(planet)

	if i < 0 {
		i = 0
	} else if i > 15 {
		i = 15
	}
	return i
}

func bonusForStarport(planet *entities.Planet) int {
	switch planet.Uwp.Starport {
	case "A":
		return 6
	case "B":
		return 4
	case "C":
		return 2
	case "X":
		return -4
	default:
		return 0
	}
}

func bonusForSize(planet *entities.Planet) int {
	switch planet.Uwp.Size {
	case 0, 1:
		return 2
	case 2, 3, 4:
		return 1
	default:
		return 0
	}
}

func bonusForAtmosphere(planet *entities.Planet) int {
	v := planet.Uwp.Atmosphere
	if v > 3 && v < 10 {
		return 0
	}
	return 1
}

func bonusForHydrographics(planet *entities.Planet) int {
	v := planet.Uwp.Hydrographic
	if v == 0 || v == 9 {
		return 1
	} else if v == 10 {
		return 2
	}
	return 0
}

func bonusForPopulation(planet *entities.Planet) int {
	v := planet.Uwp.Population
	if (v >= 1 && v <= 5) || v == 8 {
		return 1
	} else if v == 9 {
		return 2
	} else if v == 10 {
		return 4
	}
	return 0
}

func bonusForGovernment(planet *entities.Planet) int {
	v := planet.Uwp.Government
	if v == 0 || v == 5 {
		return 1
	} else if v == 7 {
		return 2
	} else if v == 13 || v == 14 {
		return -2
	}
	return 0
}

func limitTechLevel(planet *entities.Planet) int {
	a := planet.Uwp.Atmosphere
	t := planet.Uwp.TechLevel
	if a == 0 || a == 1 {
		return util.Max(t, 8)
	} else if a == 2 || a == 3 {
		return util.Max(t, 5)
	} else if a == 4 || a == 7 || a == 9 {
		return util.Max(t, 3)
	} else if a == 10 {
		return util.Max(t, 8)
	} else if a == 11 {
		return util.Max(t, 9)
	} else if a == 12 {
		return util.Max(t, 10)
	} else if a == 13 || a == 14 {
		return util.Max(t, 5)
	} else if a == 15 {
		return util.Max(t, 8)
	}

	return t
}
