package model

//Inventory is the model for stock inventory database
type Inventory struct {
	SKU      string `json:"sku"`
	Name     string `json:"name"`
	Amount   int    `json:"amount,string"`
	AvgPrice int    `json:"avg_price,string"`
}
