package api

import (
	"database/sql"

	"github.com/ilmiawan/ssinventory/model"
)

//GetPurchasingData is the function to get one purchase object
func GetPurchasingData(purchaseID int) (model.Purchasing, error) {

	stmt := "select id, sku, purchasing_date, req_amount, rec_amount, price, total, receipt_no, notes from purchasing where id=?"
	row := runQueryRow(stmt, purchaseID)

	var pur model.Purchasing
	err := row.Scan(&pur.ID, &pur.SKU, &pur.PurchasingDate, &pur.ReqAmount, &pur.RecAmount, &pur.Price, &pur.Total, &pur.ReceiptNo, &pur.Notes)

	return pur, err
}

//ListAllPurchasingData function is to retrieve all purchasing data
func ListAllPurchasingData() ([]model.Purchasing, error) {
	queryString := "select id,sku,purchasing_date,req_amount,rec_amount,price,total,receipt_no,notes from purchasing"
	rows, err := runQuery(queryString)

	purs := []model.Purchasing{}
	var pur model.Purchasing

	if rows != nil {
		for rows.Next() {
			err = rows.Scan(&pur.ID, &pur.SKU, &pur.PurchasingDate, &pur.ReqAmount, &pur.RecAmount, &pur.Price, &pur.Total, &pur.ReceiptNo, &pur.Notes)
			purs = append(purs, pur)
		}
	}

	return purs, err
}

//SavePurchase function is to save all purchase data
func SavePurchase(pur model.Purchasing) (sql.Result, error) {
	queryString := `
		INSERT INTO purchasing(sku, purchasing_date, req_amount, rec_amount, price, total, receipt_no, notes) 
		VALUES(?, ?, ?, ?, ?, ?, ?, ?)
		`
	return runExecPreparedStatement(queryString, pur.SKU, pur.PurchasingDate, pur.ReqAmount, pur.RecAmount, pur.Price, pur.Total, pur.ReceiptNo, pur.Notes)

	//SumInventoryAmount(pur.RecAmount, pur.Price, pur.SKU, w)
}

//EditPuchasingData function is to process editing purchase data
func EditPuchasingData(pur model.Purchasing) (sql.Result, error) {
	queryString := `UPDATE purchasing SET sku = ?, purchasing_date = ?, req_amount = ?, rec_amount = ?,
		price = ?, total = ?, receipt_no = ?, notes = ?  WHERE id=? `
	return runExecPreparedStatement(queryString, pur.SKU, pur.PurchasingDate, pur.ReqAmount, pur.RecAmount, pur.Price, pur.Total, pur.ReceiptNo, pur.Notes, pur.ID)
}

//DeletePurchaseData is to delete purchasing data
func DeletePurchaseData(purchaseID int) (sql.Result, error) {
	queryString := "delete from purchasing where id = ?"
	return runExecPreparedStatement(queryString, purchaseID)
}
