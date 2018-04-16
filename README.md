# ssinventory

Hi, 

This is a basic inventory management system, that records 3 things. 

Inventory, Purchasing, Sales

Inside is the basic crud webservice api with basic code structuring in golang.

There's also a sample on how to generate a report using go csv library.

The only framework I used for this is the go-sqlite3, so you'll need to add it first to your library:

$go get github.com/mattn/go-sqlite3

another thing is you need to have sqlite installed in your machine.

=======================================================================

The Project is based on 3 main tables :

CREATE TABLE IF NOT EXISTS inventory (
			sku TEXT NOT NULL PRIMARY KEY,
			name TEXT,
			amount INT,
			avg_price INT
		);

CREATE TABLE IF NOT EXISTS purchasing (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT ,
			sku TEXT NOT NULL,
			purchasing_date NUMERIC,
			req_amount INT,
			rec_amount INT,
			price INT,
			total INT,
			receipt_no TEXT,
			notes TEXT
		);
		
CREATE TABLE IF NOT EXISTS sales (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			sales_id TEXT NOT NULL,
			sku TEXT NOT NULL,
			sales_date NUMERIC,
			amount INT,
			price INT,
			total INT,
			notes TEXT
		);
		
		
To Generate the tables, all you have to do is run the application and then hit the url : http://localhost:8080/createTables

and the next thing to try is add data to your tables by posting to 

Inventory :

http://localhost:8080/inventory/addBulk

with body example :
`[{
 "sku" : "SSI-D01401050-MM-RED",
 "name" : "Zeomila Zipper Casual Blouse (M,Red)",
 "amount" : "73",
 "avgprice" : "70102"
},
{
 "sku" : "SSI-D01037822-XX-BLA",
 "name" : "Dellaya Plain Loose Big Blouse (XXL,Black)",
 "amount" : "8",
 "avgprice" : "70587"
},
{
 "sku" : "SSI-D01322275-XL-WHI",
 "name" : "Thafqya Plain Raglan Blouse (XL,White)",
 "amount" : "116",
 "avgprice" : "69382"
},
{
 "sku" : "SSI-D01401064-XL-RED",
 "name" : "Zeomila Zipper Casual Blouse (XL,Red)",
 "amount" : "44",
 "avgprice" : "70854"
},
{
 "sku" : "SSI-D01220338-XX-SAL",
 "name" : "Zeomila Zipper Casual Blouse (XL,Red)",
 "amount" : "65",
 "avgprice" : "70198"
}]`


Purchasing :

http://localhost:8080/purchasing/addBulk

with body example :
`[{
	"sku" : "SSI-D01401050-MM-RED",
	"purchasing_date" : "25/11/2017  05:37:00",
	"req_amount" : "28",
	"rec_amount" : "28",
	"price" : "69000",
	"total" : "1932000",
	"receipt_no" : "20171125-41065",
	"notes" : "2017/11/28 terima 28"
},
{
	"sku" : "SSI-D01037822-XX-BLA",
	"purchasing_date" : "03/08/2017  15:35:00",
	"req_amount" : "39",
	"rec_amount" : "39",
	"price" : "78000",
	"total" : "3042000",
	"receipt_no" : "20170803-40409",
	"notes" : "2017/08/05 terima 39"
},
{
	"sku" : "SSI-D01322275-XL-WHI",
	"purchasing_date" : "11/05/2017  14:31:00",
	"req_amount" : "136",
	"rec_amount" : "136",
	"price" : "61000",
	"total" : "8296000",
	"receipt_no" : "20170511-78054",
	"notes" : "2017/05/15 terima 136"
},
{
	"sku" : "SSI-D01401064-XL-RED",
	"purchasing_date" : "02/01/2018  07:13:00",
	"req_amount" : "63",
	"rec_amount" : "35",
	"price" : "61000",
	"total" : "3843000",
	"receipt_no" : "20180102-69458",
	"notes" : "2018/01/06 terima 35; Masih menunggu"
},
{
	"sku" : "SSI-D01220338-XX-SAL",
	"purchasing_date" : "30/10/2017  13:34:00",
	"req_amount" : "33",
	"rec_amount" : "33",
	"price" : "70000",
	"total" : "2310000",
	"receipt_no" : "20171030-72459",
	"notes" : "2017/11/01 terima 33"
}]`

the id will be autoincremental integer

Sales :

http://localhost:8080/sales/addBulk

with body example :
`[{
	"sales_id" : "ID-20180109-853724",
	"sku" : "SSI-D01401050-MM-RED",
	"sales_date" : "09/01/2018  02:38:36",
	"amount" : "1",
	"price" : "115000",
	"Total" : "115000",
	"notes" : "Pesanan ID-20180109-853724"
},
{
	"sales_id" : "ID-20180109-853724",
	"sku" : "SSI-D01037822-XX-BLA",
	"sales_date" : "09/01/2018  02:38:36",
	"amount" : "1",
	"price" : "125000",
	"Total" : "125000",
	"notes" : "Pesanan ID-20180109-853724"
},
{
	"sales_id" : "ID-20180108-149680",
	"sku" : "SSI-D01322275-XL-WHI",
	"sales_date" : "09/01/2018  02:07:08",
	"amount" : "2",
	"price" : "115000",
	"Total" : "230000",
	"notes" : "Pesanan ID-20180108-149680"
},
{
	"sales_id" : "ID-20180108-548167",
	"sku" : "SSI-D01401064-XL-RED",
	"sales_date" : "08/01/2018  22:05:15",
	"amount" : "1",
	"price" : "115000",
	"Total" : "115000",
	"notes" : "Pesanan ID-20180108-548167"
},
{
	"sales_id" : "ID-20180108-170723",
	"sku" : "SSI-D01220338-XX-SAL",
	"sales_date" : "08/01/2018  15:26:44",
	"amount" : "1",
	"price" : "130000",
	"Total" : "130000",
	"notes" : "Pesanan ID-20180108-170723"
}]`

so does the sales id will also be automatically inputted

there's also a single insert endpoints :

http://localhost:8080/inventory/add

http://localhost:8080/purchasing/add

http://localhost:8080/sales/add


to generate the inventory report just access these url:

http://localhost:8080/report/inventory

http://localhost:8080/report/inventory

and the report will be generated in the resources folder of the projects, since it is currently set it like so.

that is for now maybe, I'll soon be updating the project again if it is possible.

the other endpoint that's provided is as in this list below

o ya, one more thing. since i don't have any access to linux/*nix environment. 
I have building you an .exe binary file to run on windows machine.

Enjoy,

cheers,



