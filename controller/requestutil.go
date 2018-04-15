package controller

import (
	"encoding/json"
	"net/http"

	"github.com/muslimilmiawan/ssinventory/model"
)

func checkInternalServerError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func validateRequestMethod(w http.ResponseWriter, r *http.Request, expectedMethod string) {
	if r.Method != expectedMethod {
		http.Error(w, "Method is not allowed", http.StatusBadRequest)
		http.Redirect(w, r, "/", http.StatusBadRequest)
	}
}

func validateEmptyBody(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Redirect(w, r, "/", http.StatusBadRequest)
	}
}

func parseArrayRequestToInventory(w http.ResponseWriter, r *http.Request) []model.Inventory {
	validateEmptyBody(w, r)

	invs := make([]model.Inventory, 0)
	err := json.NewDecoder(r.Body).Decode(&invs)
	checkInternalServerError(err, w)

	return invs
}

func parseRequestToInventory(w http.ResponseWriter, r *http.Request) model.Inventory {
	validateEmptyBody(w, r)

	var inv model.Inventory
	err := json.NewDecoder(r.Body).Decode(&inv)
	checkInternalServerError(err, w)

	return inv
}

func parseArrayRequestToPurchasing(w http.ResponseWriter, r *http.Request) []model.Purchasing {
	validateEmptyBody(w, r)

	purs := make([]model.Purchasing, 0)
	err := json.NewDecoder(r.Body).Decode(&purs)
	checkInternalServerError(err, w)

	return purs
}

func parseRequestToPurchasing(w http.ResponseWriter, r *http.Request) model.Purchasing {
	validateEmptyBody(w, r)

	var pur model.Purchasing
	err := json.NewDecoder(r.Body).Decode(&pur)
	checkInternalServerError(err, w)

	return pur
}

func parseArrayRequestToSales(w http.ResponseWriter, r *http.Request) []model.Sales {
	validateEmptyBody(w, r)

	sales := make([]model.Sales, 0)
	err := json.NewDecoder(r.Body).Decode(&sales)
	checkInternalServerError(err, w)

	return sales
}

func parseRequestToSales(w http.ResponseWriter, r *http.Request) model.Sales {
	validateEmptyBody(w, r)

	var sale model.Sales
	err := json.NewDecoder(r.Body).Decode(&sale)
	checkInternalServerError(err, w)

	return sale
}
