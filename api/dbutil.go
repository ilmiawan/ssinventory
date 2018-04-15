package api

import (
	"database/sql"
	"net/http"
)

var (
	db  *sql.DB
	err error
)

func openDBConn() {
	db, err = sql.Open("sqlite3", "inventory.db")
	if err != nil {
		panic(err)
	}

	// test connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

//CreateTables is the function to create tables
func CreateTables(w http.ResponseWriter, r *http.Request) {
	openDBConn()
	defer db.Close()
	createInventoryTable()
	createPurcashingTable()
	createSalesTable()
}

func createInventoryTable() {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS inventory (
			sku TEXT NOT NULL PRIMARY KEY,
			name TEXT,
			amount INT,
			avgprice INT,
			created_date NUMERIC,
			updated_daet NUMERIC
		);
		`)

	stmt.Exec()
}

func createPurcashingTable() {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS purchasing (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT ,
			sku TEXT NOT NULL,
			purchasing_date NUMERIC,
			req_amount INT,
			rec_amount INT,
			price INT,
			receipt_no TEXT,
			notes TEXT
		);
		`)

	stmt.Exec()
}

func createSalesTable() {
	stmt, _ := db.Prepare(`
		CREATE TABLE IF NOT EXISTS sales (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			sku TEXT NOT NULL,
			sales_date NUMERIC,
			amount INT,
			price INT,
			notes TEXT
		);
		`)

	stmt.Exec()
}

func checkInternalServerError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
