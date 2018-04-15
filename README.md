# ssinventory

Hi, 

This is a basic inventory management system, that records 3 things.
Inventory
Purchasing
Sales

Inside is the basic crud webservice api with basic code structuring in golang.
there's also a sample on how to generate a report using go csv library.

The only framework I used for this is the go-sqlite3, so you'll need to add it first to your library:

$go get github.com/mattn/go-sqlite3

another thing is you need to have sqlite installed in your machine.

======================================================================================================

The Project is based on 3 main tables :

CREATE TABLE IF NOT EXISTS inventory (
			sku TEXT NOT NULL PRIMARY KEY,
			name TEXT,
			amount INT,
			avgprice INT,
			created_date NUMERIC,
			updated_daet NUMERIC
		);
		
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

CREATE TABLE IF NOT EXISTS sales (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			sku TEXT NOT NULL,
			sales_date NUMERIC,
			amount INT,
			price INT,
			notes TEXT
		);
		
		
To Generate the tables all you have to do is run the application first and then access :

http://localhost:8080/createTables

this will create the tables in your database.

and the next thing to try is add data to your tables by posting to 

Inventory :

http://localhost:8080/inventory/add

with body example :

{
 "sku" : "SSI-D01401064-XL-RED",
 "name" : "Salyara Plain Casual Big Blouse (XXL,Black)",
 "amount" : "200",
 "avgprice" : "80200",
 "created_date" : "",
 "updated_daet" : ""
}

don't worry, the created_date and updated_daet will automatically input with current datetime.

Purchasing :

http://localhost:8080/purchasing/add

with body example :
{
	"id" : "",
	"sku" : "SSI-D01401064-XL-RED",
	"purchasing_date" : "",
	"req_amount" : "200",
	"rec_amount" : "75",
	"price" : "50300",
	"receipt_no" : "SSI-D01401064-XL-RED",
	"notes" : "Pesanan ID-20170726-003674"
}

the id will be autoincremental integer, purchasingdate as always is current datetime

Sales :

http://localhost:8080/sales/add

with body example :
{
	"id" : ""
	"sku" : "SSI-D01401064-XL-RED"
	"sales_date" : ""
	"amount" : "1"
	"price" : "125000"
	"notes" : "Pesanan ID-20180108-548167"
}

so does the sales id and sales_date also automatically be inputted

there's also a bulk insert endpoint if you don't want to put it one by one
http://localhost:8080/inventory/addBulk
http://localhost:8080/purchasing/addBulk
http://localhost:8080/sales/addBulk

just make an array from the body examples earlier for input body

to generate the inventory report just access this url
http://localhost:8080/report/inventory

and the report will be generated in the resources folder of the projects, since it is currently set it like so.

that is for now maybe, I'll soon be updating the project again if it is possible.

the other endpoint that's provided is as in this list below

"/inventory"
"/inventory/list"
"/inventory/add"
"/inventory/addBulk"
"/inventory/update"
"/inventory/delete"
"/inventory/readfile"

"/purchasing"
"/purchasing/list"
"/purchasing/add"
"/purchasing/addBulk"
"/purchasing/update"
"/purchasing/delete"

"/sales"
"/sales/list"
"/sales/add"
"/sales/addBulk"
"/sales/update"
"/sales/delete"

o ya, one more thing. since i don't have any access to linux/*nix environment. 
I have building you an .exe binary file to run on windows machine.

Enjoy,

cheers,



