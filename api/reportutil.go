package api

import (
	"encoding/csv"
	"net/http"
	"os"
	"strconv"

	"github.com/muslimilmiawan/ssinventory/model"
)

var header = [][]string{}
var data = [][]string{}

//CreateInventoryReport function for generating inventory reports csv
func CreateInventoryReport(w http.ResponseWriter, head [][]string, invs []model.Inventory) {
	data = append(data, []string{"SKU", "Nama Item", "Jumlah", "Rata-Rata Harga Beli", "Total"})

	for _, inv := range invs {
		total := inv.AvgPrice * inv.Amount
		item := []string{inv.SKU, inv.Name, strconv.Itoa(inv.Amount), strconv.Itoa(inv.AvgPrice), strconv.Itoa(total)}
		data = append(data, item)
	}

	header = head
	generateReport("./resources/Laporan_Nilai_Barang.csv", w)
}

func generateReport(filename string, w http.ResponseWriter) {
	file, err := os.Create(filename)
	checkInternalServerError(err, w)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.WriteAll(header)
	writer.WriteAll(data)
}
