package api

import (
	"fmt"
	"net/http"

	"github.com/ilmiawan/ssinventory/model"
)

//GetPurchasingData is the function to get one purchase object
func GetPurchasingData(purchaseID int, w http.ResponseWriter) model.Purchasing {
	stmt := `select id, sku, purchasing_date, req_amount, rec_amount, price, total, receipt_no, notes from purchasing where id=$1;`

	openDBConn()
	defer db.Close()
	row := db.QueryRow(stmt, purchaseID)

	var pur model.Purchasing

	err := row.Scan(&pur.ID, &pur.SKU, &pur.PurchasingDate, &pur.ReqAmount, &pur.RecAmount, &pur.Price, &pur.Total, &pur.ReceiptNo, &pur.Notes)

	if err != nil {
		panic(err)
	}

	return pur
}

//ListAllPurchasingData function is to retrieve all purchasing data
func ListAllPurchasingData(w http.ResponseWriter) []model.Purchasing {
	openDBConn()
	defer db.Close()
	rows, err := db.Query("select id,sku,purchasing_date,req_amount,rec_amount,price,total,receipt_no,notes from purchasing")
	checkInternalServerError(err, w)

	purs := []model.Purchasing{}
	var pur model.Purchasing

	if rows != nil {
		for rows.Next() {
			err = rows.Scan(&pur.ID, &pur.SKU, &pur.PurchasingDate, &pur.ReqAmount, &pur.RecAmount, &pur.Price, &pur.Total, &pur.ReceiptNo, &pur.Notes)
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
		INSERT INTO purchasing(sku, purchasing_date, req_amount, rec_amount, price, total, receipt_no, notes) 
		VALUES(?, ?, ?, ?, ?, ?, ?, ?)
		`)
	checkInternalServerError(err, w)

	_, err = stmt.Exec(pur.SKU, pur.PurchasingDate, pur.ReqAmount, pur.RecAmount, pur.Price, pur.Total, pur.ReceiptNo, pur.Notes)
	checkInternalServerError(err, w)

	//SumInventoryAmount(pur.RecAmount, pur.Price, pur.SKU, w)
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
		total = ?,
		receipt_no = ?,
		notes = ? 
		WHERE id=?
		`)
	checkInternalServerError(err, w)

	_, err = stmt.Exec(pur.SKU, pur.PurchasingDate, pur.ReqAmount, pur.RecAmount, pur.Price, pur.Total, pur.ReceiptNo, pur.Notes, pur.ID)
	checkInternalServerError(err, w)
}

//DeletePurchaseData is to delete purchasing data
func DeletePurchaseData(purchaseID int, w http.ResponseWriter) {
	openDBConn()
	defer db.Close()

	stmt, err := db.Prepare("delete	from purchasing where id = ?")
	checkInternalServerError(err, w)

	res, err := stmt.Exec(purchaseID)
	checkInternalServerError(err, w)

	affect, err := res.RowsAffected()
	checkInternalServerError(err, w)

	fmt.Println(affect)
}
