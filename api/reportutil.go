package api

import (
	"encoding/csv"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/muslimilmiawan/ssinventory/model"
)

var header = [][]string{}
var data = [][]string{}

//CreateInventoryReport function for generating inventory reports csv
func CreateInventoryReport(w http.ResponseWriter) {
	data = append(data, []string{"SKU", "Nama Item", "Jumlah", "Rata-Rata Harga Beli", "Total"})

	invs := ListAllInventories(w)
	grandTotal := 0
	totalAmounts := 0

	for _, inv := range invs {
		grandTotal += inv.AvgPrice * inv.Amount
		totalAmounts += inv.Amount
		total := inv.AvgPrice * inv.Amount
		item := []string{inv.SKU, inv.Name, strconv.Itoa(inv.Amount), strconv.Itoa(inv.AvgPrice), strconv.Itoa(total)}
		data = append(data, item)
	}

	head := generateInventoryReportHeader(len(invs), totalAmounts, grandTotal)

	for _, h := range head {
		header = append(header, []string{h.Title, h.Value})
	}

	generateReport("./resources/Laporan_Nilai_Barang.csv", w)
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

//CreateSalesReport is for generating sales report function
func CreateSalesReport(w http.ResponseWriter) {
	data = append(data, []string{"ID Pesanan", "Waktu", "SKU", "Nama Barang", "Jumlah", "Harga Jual", "Total", "Harga Beli", "Laba"})

	salesReports := ListSalesReportData(w)
	totalOmzet := 0
	totalProfit := 0
	totalSales := len(salesReports)
	totalStock := 0

	for _, sr := range salesReports {
		totalOmzet += sr.Total
		totalStock += sr.Amount
		profit := sr.Total - (sr.BuyPrice * sr.Amount)
		totalProfit += profit
		item := []string{sr.SalesID, sr.SalesDate, sr.SKU, sr.Name, strconv.Itoa(sr.Amount), strconv.Itoa(sr.SalePrice), strconv.Itoa(sr.Total), strconv.Itoa(sr.BuyPrice), strconv.Itoa(profit)}
		data = append(data, item)
	}

	head := generateSalesReportHeader(totalOmzet, totalProfit, totalSales, totalStock)

	for _, h := range head {
		header = append(header, []string{h.Title, h.Value})
	}

	generateReport("./resources/Laporan_Penjualan.csv", w)
}

func generateSalesReportHeader(totalOmzet int, totalProfit int, totalSales int, totalStock int) []model.ReportHeaders {
	var head []model.ReportHeaders

	head = append(head, model.ReportHeaders{
		Title: "LAPORAN PENJUALAN",
		Value: "",
	})

	head = append(head, model.ReportHeaders{
		Title: "Tanggal Cetak",
		Value: time.Now().Format("02 January 2018"),
	})

	head = append(head, model.ReportHeaders{
		Title: "Total Omzet",
		Value: strconv.Itoa(totalOmzet),
	})

	head = append(head, model.ReportHeaders{
		Title: "Total Laba Kotor",
		Value: strconv.Itoa(totalProfit),
	})

	head = append(head, model.ReportHeaders{
		Title: "Total Penjualan",
		Value: strconv.Itoa(totalSales),
	})

	head = append(head, model.ReportHeaders{
		Title: "Total Barang",
		Value: strconv.Itoa(totalStock),
	})

	return head
}

//ListSalesReportData function is to retrieve sales report data
func ListSalesReportData(w http.ResponseWriter) []model.SalesReport {
	openDBConn()
	defer db.Close()
	rows, err := db.Query("SELECT s.sales_id, s.sales_date, s.sku, i.name, s.amount, s.price sale_price, s.total, p.price buy_price, i.amount inv_amount FROM sales s LEFT JOIN purchasing p ON s.sku = p.sku LEFT JOIN inventory i ON s.sku = i.sku GROUP BY s.sales_id, s.sales_date")
	checkInternalServerError(err, w)

	srs := []model.SalesReport{}
	var sr model.SalesReport

	if rows != nil {
		for rows.Next() {
			err = rows.Scan(&sr.SalesID, &sr.SalesDate, &sr.SKU, &sr.Name, &sr.Amount, &sr.SalePrice, &sr.Total, &sr.BuyPrice, &sr.InvAmount)
			checkInternalServerError(err, w)
			srs = append(srs, sr)
		}
	}

	return srs
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
