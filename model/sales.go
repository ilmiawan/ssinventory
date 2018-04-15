package model

//Sales is the model for stock sales database
type Sales struct {
	ID        int    `json:"id,string,omitempty"`
	SKU       string `json:"sku,omitempty"`
	SalesDate string `json:"sales_date,omitempty"`
	Amount    int    `json:"amount,string,omitempty"`
	Price     int    `json:"price,string,omitempty"`
	Notes     string `json:"notes,omitempty"`
}
