package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/hculpan/travtools/pkg/entities"
)

/*func main() {
	data, err := embed.ReadDataFile("spec-trade-data.json")
	if err != nil {
		log.Fatal(err)
	}

	tradeGoods, err := unmarshalTradeGoods(data)
	if err != nil {
		log.Fatal(err)
	}

	for _, tradeGood := range tradeGoods {
		fmt.Println(tradeGood.String())
	}
}

func unmarshalTradeGoods(jsonData []byte) (entities.TradeGoods, error) {
	var goods entities.TradeGoods
	err := json.Unmarshal(jsonData, &goods)
	if err != nil {
		return nil, err
	}

	return goods, nil
}*/

func main() {
	// Open the CSV file
	file, err := os.Open("spec trade - base data.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a new CSV reader
	reader := csv.NewReader(file)

	// Read all lines
	lines, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	// Process each line
	tradeGoods := entities.TradeGoods{}
	for _, line := range lines {
		if line[0] != "#" && strings.Trim(line[1], " ") != "" {
			tradeGood := NewTradeGoodFromCsv(line)
			tradeGoods = append(tradeGoods, *tradeGood)
		}
	}
	writeJSONToFile("spec-trade-data.json", tradeGoods)
}

func writeJSONToFile(filename string, data entities.TradeGoods) error {
	// Marshal the TradeGoods slice to JSON
	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}

	// Create or truncate the file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the JSON data to the file
	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}

func NewTradeGoodFromCsv(line []string) *entities.TradeGood {
	result := entities.NewTradeGood(line[1], line[2], line[22], line[23], line[25])

	result.Availability = GetAvailable(line)
	result.PurchaseDm = GetDms(line, 27)
	result.SaleDm = GetDms(line, 46)
	n, err := strconv.Atoi(strings.Trim(line[64], " "))
	if err == nil {
		result.SaleDm["ZA"] = n
	}
	n, err = strconv.Atoi(strings.Trim(line[65], " "))
	if err == nil {
		result.SaleDm["ZR"] = n
	}

	return result
}

func GetDms(line []string, baseIndex int) map[string]int {
	result := map[string]int{}
	for code, i := range entities.CodeOffset {
		if strings.Trim(line[baseIndex+i], " ") != "" {
			mod, err := strconv.Atoi(line[baseIndex+i])
			if err != nil {
				log.Fatalf("error at %s: %s", line[1], err.Error())
			}
			result[code] = mod
		}
	}

	return result
}

func GetAvailable(line []string) []string {
	result := []string{}
	for code, i := range entities.CodeOffset {
		if strings.Trim(line[i+3], " ") != "" {
			result = append(result, code)
		}
	}
	return result
}
