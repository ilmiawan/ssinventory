package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/ilmiawan/ssinventory/controller"

	_ "github.com/mattn/go-sqlite3"
	"github.com/ilmiawan/ssinventory/api"
)

func main() {
	http.HandleFunc("/", index)
	// inventory routing
	http.HandleFunc("/inventory", controller.GetInventory)
	http.HandleFunc("/inventory/list", controller.ListInventory)
	http.HandleFunc("/inventory/add", controller.AddInventory)
	http.HandleFunc("/inventory/addBulk", controller.AddBulkInventory)
	http.HandleFunc("/inventory/update", controller.UpdateInventory)
	http.HandleFunc("/inventory/delete", controller.DeleteInventoryController)
	http.HandleFunc("/inventory/readfile", controller.MigrateInventoryFromFile)
	// purchasing routing
	http.HandleFunc("/purchasing", controller.GetPurchase)
	http.HandleFunc("/purchasing/list", controller.ListPurchase)
	http.HandleFunc("/purchasing/add", controller.AddPurchase)
	http.HandleFunc("/purchasing/addBulk", controller.AddBulkPurchase)
	http.HandleFunc("/purchasing/update", controller.UpdatePurchase)
	http.HandleFunc("/purchasing/delete", controller.DeletePurchase)
	// sales routing
	http.HandleFunc("/sales", controller.GetSales)
	http.HandleFunc("/sales/list", controller.ListSales)
	http.HandleFunc("/sales/add", controller.AddSales)
	http.HandleFunc("/sales/addBulk", controller.AddBulkSales)
	http.HandleFunc("/sales/update", controller.UpdateSales)
	http.HandleFunc("/sales/delete", controller.DeleteSales)

	// report routing
	http.HandleFunc("/report/inventory", controller.GenerateInventoryReport)
	http.HandleFunc("/report/sales", controller.GenerateSalesReport)

	http.HandleFunc("/createTables", api.CreateTables)
	http.HandleFunc("/migrate/inventory", controller.MigrateInventoryFromFile)
	http.HandleFunc("/migrate/purchasing", controller.MigratePurchasingFromFile)
	http.HandleFunc("/migrate/sales", controller.MigrateSalesFromFile)
	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hi, There this is Muslim's Sale Stock Code test")
}
