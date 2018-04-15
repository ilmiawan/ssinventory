package api

import (
	"encoding/json"
	"net/http"

	"github.com/muslimilmiawan/ssinventory/model"
)

//GetPurchasingData is the function to get one purchase object
func GetPurchasingData(purchaseID int, w http.ResponseWriter) {
	openDBConn()
	defer db.Close()
	stmt := `select id, sku, purchasing_date, req_amount, rec_amount, price, receipt_no, notes from purchasing where sku=$1;`
	checkInternalServerError(err, w)

	row := db.QueryRow(stmt, purchaseID)

	json.NewEncoder(w).Encode(row)
}

//ListAllPurchasingData function is to retrieve all purchasing data
func ListAllPurchasingData(w http.ResponseWriter) []model.Purchasing {
	openDBConn()
	defer db.Close()
	rows, err := db.Query("select id,sku,purchasing_date,req_amount,rec_amount,price,receipt_no,notes from purchasing")
	checkInternalServerError(err, w)

	purs := []model.Purchasing{}
	var pur model.Purchasing

	if rows != nil {
		for rows.Next() {
			err = rows.Scan(&pur.ID, &pur.SKU, &pur.PurchasingDate, &pur.ReqAmount, &pur.RecAmount, &pur.Price, &pur.ReceiptNo, &pur.Notes)
			checkInternalServerError(err, w)
			purs = append(purs, pur)
		}
	}

	return purs
}

//SavePurchase function is to save all purchase data
func SavePurchase(pur model.Purchasing, w http.ResponseWriter) {
	openDBConn()
	defer db.Close()

	stmt, err := db.Prepare(`
		INSERT INTO purchasing(sku, req_amount, rec_amount, price, receipt_no, notes, purchasing_date) 
		VALUES(?, ?, ?, ?, ?, ?, date('now'))
		`)
	checkInternalServerError(err, w)

	_, err = stmt.Exec(pur.SKU, pur.ReqAmount, pur.RecAmount, pur.Price, pur.ReceiptNo, pur.Notes)
	checkInternalServerError(err, w)

	SumInventoryAmount(pur.RecAmount, pur.Price, pur.SKU, w)
}

//EditPuchasingData function is to process editing purchase data
func EditPuchasingData(pur model.Purchasing, w http.ResponseWriter) {
	openDBConn()
	defer db.Close()
	stmt, err := db.Prepare(`
		UPDATE purchasing sku = ?,
		purchasing_date = ?,
		req_amount = ?,
		rec_amount = ?,
		price = ?,
		receipt_no = ?,
		notes = ? 
		WHERE id=?
		`)
	checkInternalServerError(err, w)

	_, err = stmt.Exec(pur.SKU, pur.PurchasingDate, pur.ReqAmount, pur.RecAmount, pur.Price, pur.ReceiptNo, pur.Notes, pur.ID)
	checkInternalServerError(err, w)
}

//DeletePurchaseData is to delete purchasing data
func DeletePurchaseData(purchaseID int, w http.ResponseWriter) {
	openDBConn()
	defer db.Close()

	stmt, err := db.Prepare("delete	from purchasing where id = ?")
	checkInternalServerError(err, w)

	stmt.Exec(purchaseID)
}
