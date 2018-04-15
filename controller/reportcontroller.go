package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/muslimilmiawan/ssinventory/api"
	"github.com/muslimilmiawan/ssinventory/model"
)

var header = [][]string{}
var data = [][]string{}

// GenerateInventoryReport function is to generate main inventory report
func GenerateInventoryReport(w http.ResponseWriter, r *http.Request) {

	invs := api.ListAllInventories(w)
	grandTotal := 0
	totalAmounts := 0
	for _, inv := range invs {
		grandTotal += inv.AvgPrice * inv.Amount
		totalAmounts += inv.Amount
	}

	head := generateInventoryReportHeader(len(invs), totalAmounts, grandTotal)

	for _, h := range head {
		header = append(header, []string{h.Title, h.Value})
	}

	api.CreateInventoryReport(w, header, invs)
}

func generateInventoryReportHeader(itemLength int, totalAmount int, grandTotal int) []model.ReportHeaders {
	var head []model.ReportHeaders

	head = append(head, model.ReportHeaders{
		Title: "LAPORAN NILAI BARANG",
		Value: "",
	})

	head = append(head, model.ReportHeaders{
		Title: "Tanggal Cetak",
		Value: time.Now().Format("02 January 2018"),
	})

	head = append(head, model.ReportHeaders{
		Title: "Jumlah SKU",
		Value: strconv.Itoa(itemLength),
	})

	head = append(head, model.ReportHeaders{
		Title: "Jumlah Total Barang",
		Value: strconv.Itoa(totalAmount),
	})

	head = append(head, model.ReportHeaders{
		Title: "Total Nilai",
		Value: strconv.Itoa(grandTotal),
	})

	return head
}
