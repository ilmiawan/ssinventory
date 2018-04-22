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
	pur := api.GetPurchasingData(purchaseID, w)

	json.NewEncoder(w).Encode(pur)
}

//ListPurchase is the function to get all purchase list
func ListPurchase(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "GET")
	purs := api.ListAllPurchasingData(w)
	json.NewEncoder(w).Encode(purs)
}

//AddPurchase is the function to get add new purchase
func AddPurchase(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "POST")

	pur := parseRequestToPurchasing(w, r)

	api.SavePurchase(pur, w)
	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

//AddBulkPurchase is to save Multiple Purchasing data
func AddBulkPurchase(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "POST")

	purs := parseArrayRequestToPurchasing(w, r)

	for index := 0; index < len(purs); index++ {
		pur := purs[index]
		api.SavePurchase(pur, w)
	}

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

//UpdatePurchase is the function to get update existing purchasing data
func UpdatePurchase(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "PUT")

	pur := parseRequestToPurchasing(w, r)
	api.EditPuchasingData(pur, w)

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

//DeletePurchase is the function to delete one purchasing object
func DeletePurchase(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "DELETE")

	var purchaseID, _ = strconv.Atoi(r.FormValue("id"))
	api.DeletePurchaseData(purchaseID, w)

	http.Redirect(w, r, "/", http.StatusPermanentRedirect)
}

//MigratePurchasingFromFile function is to migrate data from purchasing csv
func MigratePurchasingFromFile(w http.ResponseWriter, r *http.Request) {
	records := api.ReadCSVFile(r.FormValue("filename"))
	purs := api.ConvertRecordsToPurchasingFile(records)

	for index := range purs {
		// remove first row, the header
		if index == 0 {
			continue
		}
		pur := purs[index]
		api.SavePurchase(pur, w)
	}

	http.Redirect(w, r, "/purchasing/list", http.StatusPermanentRedirect)
}
