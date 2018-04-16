package api

import (
	"net/http"

	"github.com/muslimilmiawan/ssinventory/model"
)

//GetSpecificInventory function is to get one selected inventory only
func GetSpecificInventory(sku string, w http.ResponseWriter) model.Inventory {
	sqlStatement := `SELECT sku, name, amount, avg_price from inventory WHERE sku=$1;`

	openDBConn()
	defer db.Close()
	row := db.QueryRow(sqlStatement, sku)

	var inv model.Inventory

	err := row.Scan(&inv.SKU, &inv.Name, &inv.Amount, &inv.AvgPrice)
	if err != nil {
		panic(err)
	}

	return inv
}

//ListAllInventories is to get all inventories data
func ListAllInventories(w http.ResponseWriter) []model.Inventory {
	openDBConn()
	defer db.Close()
	rows, err := db.Query("select sku, name, amount, avg_price from inventory")
	checkInternalServerError(err, w)

	invs := []model.Inventory{}
	var inv model.Inventory

	if rows != nil {
		for rows.Next() {
			err = rows.Scan(&inv.SKU, &inv.Name, &inv.Amount, &inv.AvgPrice)
			checkInternalServerError(err, w)
			invs = append(invs, inv)
		}
	}

	return invs
}

//SaveInventory is the function to create new inventory
func SaveInventory(inv model.Inventory, w http.ResponseWriter) {
	openDBConn()
	defer db.Close()

	stmt, err := db.Prepare(`
		INSERT INTO inventory(sku, name, amount, avg_price) 
		VALUES(?, ?, ?, ?)
		`)
	checkInternalServerError(err, w)

	_, err = stmt.Exec(inv.SKU, inv.Name, inv.Amount, inv.AvgPrice)
	checkInternalServerError(err, w)
}

//EditInventoryData function to run inventory modification
func EditInventoryData(inv model.Inventory, w http.ResponseWriter) {
	openDBConn()
	defer db.Close()

	stmt, err := db.Prepare(`
		UPDATE inventory SET name=?, amount=?, avg_price=?
		WHERE sku=?
		`)
	checkInternalServerError(err, w)

	_, err = stmt.Exec(inv.Name, inv.Amount, inv.AvgPrice, inv.SKU)
	checkInternalServerError(err, w)

}

//DeleteInventory is the process of deleting single object
func DeleteInventory(sku string, w http.ResponseWriter) {
	openDBConn()
	defer db.Close()

	stmt, err := db.Prepare(`
		"delete	from inventory 
		where sku = ?"`)
	checkInternalServerError(err, w)

	stmt.Exec(sku)
}

//SumInventoryAmount function is to update inventory when purchasing or sales
func SumInventoryAmount(amount int, transactionAveragePrice int, sku string, w http.ResponseWriter) {
	inv := GetSpecificInventory(sku, w)

	inv.Amount = inv.Amount + amount
	if transactionAveragePrice == 0 {
		inv.AvgPrice = (inv.AvgPrice + transactionAveragePrice) / 2
	}

	EditInventoryData(inv, w)
}
