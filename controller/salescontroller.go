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
	sale := api.GetSalesData(salesID, w)

	json.NewEncoder(w).Encode(sale)
}

//ListSales is the function to get all sales list
func ListSales(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "GET")
	sales := api.ListAllSalesData(w)
	json.NewEncoder(w).Encode(sales)
}

//AddSales is the function to get add new stock sales
func AddSales(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "POST")

	sale := parseRequestToSales(w, r)

	api.SaveSales(sale, w)
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

//AddBulkSales is to save Multiple stock sales data
func AddBulkSales(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "POST")

	sales := parseArrayRequestToSales(w, r)

	for index := 0; index < len(sales); index++ {
		sale := sales[index]
		api.SaveSales(sale, w)
	}

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

//UpdateSales is the function to get update existing stock sales data
func UpdateSales(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "PUT")

	sale := parseRequestToSales(w, r)
	api.EditSales(sale, w)

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

//DeleteSales is the function to delete one stock sales object
func DeleteSales(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "DELETE")

	var salesID, _ = strconv.Atoi(r.FormValue("id"))
	api.DeleteSalesData(salesID, w)

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

//MigrateSalesFromFile function is to migrate data from sales csv
func MigrateSalesFromFile(w http.ResponseWriter, r *http.Request) {
	records := api.ReadCSVFile(r.FormValue("filename"))
	sales := api.ConvertRecordsToSalesFile(records)

	for index := range sales {
		// remove first row, the header
		if index == 0 {
			continue
		}
		sale := sales[index]
		api.SaveSales(sale, w)
	}

	http.Redirect(w, r, "/sales/list", http.StatusPermanentRedirect)
}
