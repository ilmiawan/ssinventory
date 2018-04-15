package controller

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"

	"github.com/muslimilmiawan/ssinventory/model"
)

//FloatToString convert float to String
func FloatToString(inputNum float64) string {
	return strconv.FormatFloat(inputNum, 'f', 2, 64)
}

func readInventoryFiles(filename string) []model.Inventory {
	csvFile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ';'

	lines, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	skus := make([]string, len(lines)-1)
	names := make([]string, len(lines)-1)
	amounts := make([]string, len(lines)-1)
	inventories := make([]model.Inventory, len(lines)-1)

	for index, line := range lines {
		if index == 0 {
			//to skip the header
			continue
		}
		skus[index-1] = line[0]
		names[index-1] = line[1]
		amounts[index-1] = line[2]
	}

	for index := range inventories {
		amountInt, err := strconv.Atoi(amounts[2])
		if err != nil {
			panic(err)
		}
		inventories = append(inventories, model.Inventory{
			SKU:         skus[index],
			Name:        names[index],
			Amount:      amountInt,
			AvgPrice:    0,
			CreatedDate: "",
			UpdatedDate: "",
		})
	}

	return inventories
}
