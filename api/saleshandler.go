package api

import (
	"encoding/json"
	"net/http"

	"github.com/muslimilmiawan/ssinventory/model"
)

//GetSalesData is the function to get one sales object
func GetSalesData(salesID int, w http.ResponseWriter) {
	stmt := `select sku, sales_date, amount, price, notes from sales where id = $1;`
	checkInternalServerError(err, w)

	openDBConn()
	defer db.Close()

	row := db.QueryRow(stmt, salesID)

	json.NewEncoder(w).Encode(row)
}

//ListAllSalesData to list all sales data
func ListAllSalesData(w http.ResponseWriter) []model.Sales {
	openDBConn()
	defer db.Close()

	rows, err := db.Query("select id, sku, sales_date, amount, price, notes from sales")
	checkInternalServerError(err, w)

	sales := []model.Sales{}
	var sale model.Sales

	if rows != nil {
		for rows.Next() {
			err = rows.Scan(&sale.ID, &sale.SKU, &sale.SalesDate, &sale.Amount, &sale.Price, &sale.Notes)
			checkInternalServerError(err, w)
			sales = append(sales, sale)
		}
	}

	return sales
}

//SaveSales is to save sales data
func SaveSales(sale model.Sales, w http.ResponseWriter) {
	openDBConn()
	defer db.Close()

	stmt, err := db.Prepare(`
		INSERT INTO sales(sku, amount, price, notes, sales_date) 
		VALUES(?, ?, ?, ?, date('now'))
		`)

	_, err = stmt.Exec(sale.SKU, sale.Amount, sale.Price, sale.Notes)
	checkInternalServerError(err, w)

	SumInventoryAmount(-sale.Amount, 0, sale.SKU, w)
}

//EditSales function to edit sales data
func EditSales(sale model.Sales, w http.ResponseWriter) {
	openDBConn()
	defer db.Close()
	stmt, err := db.Prepare(`
		UPDATE sales SET sku=?, sales_date=?, amount=?, price=?, notes=? 
		WHERE id=?
		`)
	checkInternalServerError(err, w)

	_, err = stmt.Exec(sale.SKU, sale.SalesDate, sale.Amount, sale.Price, sale.Notes, sale.ID)
	checkInternalServerError(err, w)
}

//DeleteSalesData function to delete requested sales data
func DeleteSalesData(salesID int, w http.ResponseWriter) {
	openDBConn()
	defer db.Close()

	stmt, err := db.Prepare(`
		"delete	from sales 
		where id = ?"`)
	checkInternalServerError(err, w)

	stmt.Exec(salesID)
}
