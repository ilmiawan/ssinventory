package controller

import (
	"encoding/json"
	"net/http"

	"github.com/ilmiawan/ssinventory/api"
)

//GetInventory is the function to get one inventory object
func GetInventory(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "GET")

	sku := r.FormValue("sku")
	inv := api.GetSpecificInventory(sku, w)

	json.NewEncoder(w).Encode(inv)
}

//ListInventory is the function to get all inventory list
func ListInventory(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "GET")

	invs := api.ListAllInventories(w)
	json.NewEncoder(w).Encode(invs)
}

//AddInventory is the function to get add new inventory
func AddInventory(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "POST")

	inv := parseRequestToInventory(w, r)

	api.SaveInventory(inv, w)
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

//AddBulkInventory is to save Multiple Inventory data
func AddBulkInventory(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "POST")

	invs := parseArrayRequestToInventory(w, r)

	for index := 0; index < len(invs); index++ {
		inv := invs[index]
		api.SaveInventory(inv, w)
	}

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

//UpdateInventory is the function to get update existing inventory data
func UpdateInventory(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "PUT")

	inv := parseRequestToInventory(w, r)
	api.EditInventoryData(inv, w)

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

//DeleteInventoryController is the function to delete one inventory object
func DeleteInventoryController(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "DELETE")

	sku := r.FormValue("sku")
	api.DeleteInventory(sku, w)

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

//MigrateInventoryFromFile function is to migrate data from inventory data file
func MigrateInventoryFromFile(w http.ResponseWriter, r *http.Request) {
	records := api.ReadCSVFile(r.FormValue("filename"))
	invs := api.ConvertRecordsToInventory(records)

	for index := range invs {
		// remove first row, the header
		if index == 0 {
			continue
		}
		inv := invs[index]
		api.SaveInventory(inv, w)
	}

	http.Redirect(w, r, "/inventory/list", http.StatusPermanentRedirect)
}
