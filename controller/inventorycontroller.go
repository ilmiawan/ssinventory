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
	inv, err := api.GetInventoryBySKU(sku)

	checkInternalServerError(err, w)

	json.NewEncoder(w).Encode(inv)
}

//ListInventory is the function to get all inventory list
func ListInventory(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "GET")

	invs, err := api.ListInventories()
	checkInternalServerError(err, w)

	json.NewEncoder(w).Encode(invs)
}

//AddInventory is the function to get add new inventory
func AddInventory(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "POST")

	inv := parseRequestToInventory(w, r)

	result, err := api.SaveInventory(inv)

	checkInternalServerError(err, w)

	checkQueryResult(result, "Inventory has been successfully Added to database.", "Cannot add inventory.", w)
}

//AddBulkInventory is to save Multiple Inventory data
func AddBulkInventory(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "POST")

	invs := parseArrayRequestToInventory(w, r)

	for index := 0; index < len(invs); index++ {
		inv := invs[index]
		_, err := api.SaveInventory(inv)
		checkInternalServerError(err, w)
	}

	http.Redirect(w, r, "/inventory/list", http.StatusPermanentRedirect)
}

//UpdateInventory is the function to get update existing inventory data
func UpdateInventory(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "PUT")

	inv := parseRequestToInventory(w, r)
	result, err := api.EditInventoryData(inv)

	checkInternalServerError(err, w)

	checkQueryResult(result, "Inventory has been successfully Updated.", "Cannot update inventory.", w)
}

//DeleteInventoryController is the function to delete one inventory object
func DeleteInventoryController(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "DELETE")

	sku := r.FormValue("sku")
	result, err := api.DeleteInventory(sku)
	checkInternalServerError(err, w)

	checkQueryResult(result, "Succesfully deleted inventory.", "Failed to delete inventory.", w)
}

//MigrateInventoryFromFile function is to migrate data from inventory data file
func MigrateInventoryFromFile(w http.ResponseWriter, r *http.Request) {
	records, err := api.ReadCSVFile(r.FormValue("filename"))
	checkInternalServerError(err, w)

	invs := api.ConvertRecordsToInventory(records)

	for index := range invs {
		// pass the header
		if index == 0 {
			continue
		}
		inv := invs[index]
		_, err := api.SaveInventory(inv)
		checkInternalServerError(err, w)
	}

	http.Redirect(w, r, "/inventory/list", http.StatusPermanentRedirect)
}
