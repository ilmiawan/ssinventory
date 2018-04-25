package api

import (
	"bufio"
	"encoding/csv"
	"os"
	"strconv"
	"strings"

	"github.com/ilmiawan/ssinventory/model"
)

//ReadCSVFile function is to read csv file and return string arrays
func ReadCSVFile(fullpathName string) ([][]string, error) {
	csvFile, err := os.Open(fullpathName)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ';'

	return reader.ReadAll()
}

//ConvertRecordsToInventory is to read files from inventory csv
func ConvertRecordsToInventory(records [][]string) []model.Inventory {

	var inventories = make([]model.Inventory, 0)

	for i := range records {
		strAmount, _ := strconv.Atoi(records[i][2])
		strAvgPrice, _ := strconv.Atoi(records[i][3])

		inventories = append(inventories, model.Inventory{
			SKU:      records[i][0],
			Name:     records[i][1],
			Amount:   strAmount,
			AvgPrice: strAvgPrice,
		})
	}

	return inventories
}

//ConvertRecordsToPurchasingFile is to read files from purchasing csv
func ConvertRecordsToPurchasingFile(records [][]string) []model.Purchasing {

	var purchasings = make([]model.Purchasing, 0)

	for i := range records {
		strReqAmount, _ := strconv.Atoi(records[i][3])
		strRecAmount, _ := strconv.Atoi(records[i][4])
		strPrice, _ := strconv.Atoi(records[i][5])
		strTotal, _ := strconv.Atoi(records[i][6])

		purchasings = append(purchasings, model.Purchasing{
			PurchasingDate: records[i][0],
			SKU:            records[i][1],
			ReqAmount:      strReqAmount,
			RecAmount:      strRecAmount,
			Price:          strPrice,
			Total:          strTotal,
			ReceiptNo:      records[i][7],
			Notes:          records[i][8],
		})
	}

	return purchasings
}

//ConvertRecordsToSalesFile is to read files from sales csv
func ConvertRecordsToSalesFile(records [][]string) []model.Sales {

	var sales = make([]model.Sales, 0)

	for i := range records {
		strAmount, _ := strconv.Atoi(records[i][3])
		strPrice, _ := strconv.Atoi(records[i][4])
		strTotal, _ := strconv.Atoi(records[i][5])
		notes := records[i][6]
		var salesID string
		if strings.Contains(notes, "Pesanan ID") {
			salesID = notes[(strings.Index(notes, " ") + 1):len(notes)]
		}

		sales = append(sales, model.Sales{
			SalesID:   salesID,
			SalesDate: records[i][0],
			SKU:       records[i][1],
			Amount:    strAmount,
			Price:     strPrice,
			Total:     strTotal,
			Notes:     records[i][6],
		})
	}

	return sales
}
