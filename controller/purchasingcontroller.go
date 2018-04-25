package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/ilmiawan/ssinventory/api"
)

//GetPurchase is the function to get one purchase object
func GetPurchase(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "GET")

	purchaseID, _ := strconv.Atoi(r.FormValue("id"))
	pur, err := api.GetPurchasingData(purchaseID)
	checkInternalServerError(err, w)

	json.NewEncoder(w).Encode(pur)
}

//ListPurchase is the function to get all purchase list
func ListPurchase(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "GET")

	purs, err := api.ListAllPurchasingData()
	checkInternalServerError(err, w)

	json.NewEncoder(w).Encode(purs)
}

//AddPurchase is the function to get add new purchase
func AddPurchase(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "POST")

	pur := parseRequestToPurchasing(w, r)
	result, err := api.SavePurchase(pur)
	checkInternalServerError(err, w)

	checkQueryResult(result, "Successfully add Purchase data.", "Failed to add Purchase data.", w)
}

//AddBulkPurchase is to save Multiple Purchasing data
func AddBulkPurchase(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "POST")

	purs := parseArrayRequestToPurchasing(w, r)

	for index := 0; index < len(purs); index++ {
		pur := purs[index]
		_, err := api.SavePurchase(pur)
		checkInternalServerError(err, w)
	}

	http.Redirect(w, r, "/purchasing/list", http.StatusPermanentRedirect)
}

//UpdatePurchase is the function to get update existing purchasing data
func UpdatePurchase(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "PUT")

	pur := parseRequestToPurchasing(w, r)
	result, err := api.EditPuchasingData(pur)
	checkInternalServerError(err, w)

	checkQueryResult(result, "Successfully update Purchase data.", "Failed to update Purchase data.", w)
}

//DeletePurchase is the function to delete one purchasing object
func DeletePurchase(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "DELETE")

	var purchaseID, _ = strconv.Atoi(r.FormValue("id"))
	result, err := api.DeletePurchaseData(purchaseID)
	checkInternalServerError(err, w)

	checkQueryResult(result, "Successfully delete Purchase data.", "Failed to delete Purchase data.", w)
}

//MigratePurchasingFromFile function is to migrate data from purchasing csv
func MigratePurchasingFromFile(w http.ResponseWriter, r *http.Request) {
	records, err := api.ReadCSVFile(r.FormValue("filename"))
	checkInternalServerError(err, w)
	purs := api.ConvertRecordsToPurchasingFile(records)

	for index := range purs {
		// pass the header
		if index == 0 {
			continue
		}
		pur := purs[index]
		_, err := api.SavePurchase(pur)
		checkInternalServerError(err, w)
	}

	http.Redirect(w, r, "/purchasing/list", http.StatusPermanentRedirect)
}
