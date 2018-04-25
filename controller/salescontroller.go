package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ilmiawan/ssinventory/api"
)

//GetSales is the function to get one sales object
func GetSales(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "GET")

	salesID, _ := strconv.Atoi(r.FormValue("id"))
	sale, err := api.GetSalesData(salesID)
	checkInternalServerError(err, w)

	json.NewEncoder(w).Encode(sale)
}

//ListSales is the function to get all sales list
func ListSales(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "GET")

	sales, err := api.ListAllSalesData()
	checkInternalServerError(err, w)

	json.NewEncoder(w).Encode(sales)
}

//AddSales is the function to get add new stock sales
func AddSales(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "POST")

	sale := parseRequestToSales(w, r)

	result, err := api.SaveSales(sale)
	checkInternalServerError(err, w)

	checkQueryResult(result, "Successfully add Sales data.", "Failed to add sales data.", w)
}

//AddBulkSales is to save Multiple stock sales data
func AddBulkSales(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "POST")

	sales := parseArrayRequestToSales(w, r)

	for index := 0; index < len(sales); index++ {
		sale := sales[index]
		_, err := api.SaveSales(sale)
		checkInternalServerError(err, w)
	}

	http.Redirect(w, r, "/sales/list", http.StatusPermanentRedirect)
}

//UpdateSales is the function to get update existing stock sales data
func UpdateSales(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "PUT")

	sale := parseRequestToSales(w, r)
	result, err := api.EditSales(sale)
	checkInternalServerError(err, w)

	checkQueryResult(result, "Successfully updated sales data.", "Failed to update sales data.", w)
}

//DeleteSales is the function to delete one stock sales object
func DeleteSales(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "DELETE")

	var salesID, _ = strconv.Atoi(r.FormValue("id"))
	result, err := api.DeleteSalesData(salesID)
	checkInternalServerError(err, w)

	checkQueryResult(result, "Successfully Deleted Sales data.", "Failed to delete sales data", w)
}

//MigrateSalesFromFile function is to migrate data from sales csv
func MigrateSalesFromFile(w http.ResponseWriter, r *http.Request) {
	records, err := api.ReadCSVFile(r.FormValue("filename"))
	checkInternalServerError(err, w)

	sales := api.ConvertRecordsToSalesFile(records)

	for index := range sales {
		// pass the header
		if index == 0 {
			continue
		}
		sale := sales[index]
		_, err := api.SaveSales(sale)
		checkInternalServerError(err, w)
	}

	http.Redirect(w, r, "/sales/list", http.StatusPermanentRedirect)
}
