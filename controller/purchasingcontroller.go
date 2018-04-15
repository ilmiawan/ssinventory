package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/muslimilmiawan/ssinventory/api"
)

//GetPurchase is the function to get one purchase object
func GetPurchase(w http.ResponseWriter, r *http.Request) {
	validateRequestMethod(w, r, "GET")

	var purchaseID, _ = strconv.Atoi(r.FormValue("id"))
	api.GetPurchasingData(purchaseID, w)
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
