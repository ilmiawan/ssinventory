package api

import (
	"database/sql"

	"github.com/ilmiawan/ssinventory/model"
)

//GetInventoryBySKU function is to get one selected inventory only
func GetInventoryBySKU(sku string) (model.Inventory, error) {

	sqlStatement := `SELECT sku, name, amount, avg_price from inventory WHERE sku = ?`
	row := runQueryRow(sqlStatement, sku)

	var inv model.Inventory
	err := row.Scan(&inv.SKU, &inv.Name, &inv.Amount, &inv.AvgPrice)

	return inv, err
}

//ListInventories is to get all inventories data
func ListInventories() ([]model.Inventory, error) {
	stmt := "select sku, name, amount, avg_price from inventory"
	rows, err := runQuery(stmt)

	invs := []model.Inventory{}
	var inv model.Inventory

	if rows != nil {
		for rows.Next() {
			err = rows.Scan(&inv.SKU, &inv.Name, &inv.Amount, &inv.AvgPrice)
			invs = append(invs, inv)
		}
	}

	return invs, err
}

//SaveInventory is the function to create new inventory
func SaveInventory(inv model.Inventory) (sql.Result, error) {
	queryString := "INSERT INTO inventory(sku, name, amount, avg_price) VALUES(?, ?, ?, ?)"
	return runExecPreparedStatement(queryString, inv.SKU, inv.Name, inv.Amount, inv.AvgPrice)
}

//EditInventoryData function to run inventory modification
func EditInventoryData(inv model.Inventory) (sql.Result, error) {
	queryString := "UPDATE inventory SET name=?, amount=?, avg_price=? WHERE sku=?"
	return runExecPreparedStatement(queryString, inv.Name, inv.Amount, inv.AvgPrice, inv.SKU)
}

//DeleteInventory is the process of deleting single object
func DeleteInventory(sku string) (sql.Result, error) {
	queryString := "delete	from inventory where sku = ?"
	return runExecPreparedStatement(queryString, sku)
}

//SumInventoryAmount function is to update inventory when purchasing or sales
func SumInventoryAmount(amount int, transactionAveragePrice int, sku string) (sql.Result, error) {
	inv, err := GetInventoryBySKU(sku)

	inv.Amount = inv.Amount + amount
	if transactionAveragePrice == 0 {
		inv.AvgPrice = (inv.AvgPrice + transactionAveragePrice) / 2
	}

	result, err := EditInventoryData(inv)

	return result, err
}
