package model

//ReportHeaders struct for generating Report headers
type ReportHeaders struct {
	Title string
	Value string
}

//SalesReport struct for generation sales report items
type SalesReport struct {
	SalesID   string `json:"sales_id"`
	SalesDate string `json:"sales_date,omitempty"`
	SKU       string `json:"sku"`
	Name      string `json:"name,omitempty"`
	Amount    int    `json:"amount,string"`
	SalePrice int    `json:"sale_price,string,omitempty"`
	Total     int    `json:"total,string,omitempty"`
	BuyPrice  int    `json:"buy_price,string,omitempty"`
	InvAmount int    `json:"inv_amount,string,omitempty"`
}
