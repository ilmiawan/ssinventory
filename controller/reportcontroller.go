package controller

import (
	"net/http"

	"github.com/muslimilmiawan/ssinventory/api"
)

// GenerateInventoryReport function is to generate main inventory report
func GenerateInventoryReport(w http.ResponseWriter, r *http.Request) {
	api.CreateInventoryReport(w)
}

// GenerateSalesReport function is to generate main sales report
func GenerateSalesReport(w http.ResponseWriter, r *http.Request) {
	api.CreateSalesReport(w)
}
