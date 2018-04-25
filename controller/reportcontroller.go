package controller

import (
	"net/http"

	"github.com/ilmiawan/ssinventory/api"
)

// GenerateInventoryReport function is to generate main inventory report
func GenerateInventoryReport(w http.ResponseWriter, r *http.Request) {
	err := api.CreateInventoryReport()
	checkInternalServerError(err, w)

	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Inventory Report has been successfully generated"))
	}
}

// GenerateSalesReport function is to generate main sales report
func GenerateSalesReport(w http.ResponseWriter, r *http.Request) {
	err := api.CreateSalesReport()
	checkInternalServerError(err, w)

	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Sales Report has been successfully generated"))
	}
}
