package api

import (
	"database/sql"

	"github.com/ilmiawan/ssinventory/model"
)

//GetSalesData is the function to get one sales object
func GetSalesData(ID int) (model.Sales, error) {

	stmt := `select id, sales_id, sku, sales_date, amount, price, total, notes from sales where id = ?`
	row := runQueryRow(stmt, ID)

	var sale model.Sales
	err := row.Scan(&sale.ID, &sale.SalesID, &sale.SKU, &sale.SalesDate, &sale.Amount, &sale.Price, &sale.Total, &sale.Notes)

	return sale, err
}

//ListAllSalesData to list all sales data
func ListAllSalesData() ([]model.Sales, error) {
	queryString := "select id, sales_id, sku, sales_date, amount, price, total, notes from sales"
	rows, err := runQuery(queryString)

	sales := []model.Sales{}
	var sale model.Sales

	if rows != nil {
		for rows.Next() {
			err = rows.Scan(&sale.ID, &sale.SalesID, &sale.SKU, &sale.SalesDate, &sale.Amount, &sale.Price, &sale.Total, &sale.Notes)
			sales = append(sales, sale)
		}
	}

	return sales, err
}

//SaveSales is to save sales data
func SaveSales(sale model.Sales) (sql.Result, error) {
	queryString := `
		INSERT INTO sales(sales_id, sku, sales_date, amount, price, total, notes) 
		VALUES(?, ?, ?, ?, ?, ?, ?)
		`
	return runExecPreparedStatement(queryString, sale.SalesID, sale.SKU, sale.SalesDate, sale.Amount, sale.Price, sale.Total, sale.Notes)

	//SumInventoryAmount(-sale.Amount, 0, sale.SKU, w)
}

//EditSales function to edit sales data
func EditSales(sale model.Sales) (sql.Result, error) {
	queryString := `
		UPDATE sales SET sales_id=?, sku=?, sales_date=?, amount=?, price=?, total=?, notes=? 
		WHERE id=?
		`
	return runExecPreparedStatement(queryString, sale.SalesID, sale.SKU, sale.SalesDate, sale.Amount, sale.Price, sale.Total, sale.Notes, sale.ID)
}

//DeleteSalesData function to delete requested sales data
func DeleteSalesData(ID int) (sql.Result, error) {
	queryString := "delete	from sales where id = ?"
	return runExecPreparedStatement(queryString, ID)
}
