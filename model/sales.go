package model

//Sales is the model for stock sales database
type Sales struct {
	ID        int    `json:"id,string,omitempty"`
	SalesID   string `json:"sales_d,omitempty"`
	SKU       string `json:"sku,omitempty"`
	SalesDate string `json:"sales_date,omitempty"`
	Amount    int    `json:"amount,string,omitempty"`
	Price     int    `json:"price,string,omitempty"`
	Total     int    `json:"total,string,omitempty"`
	Notes     string `json:"notes,omitempty"`
}
