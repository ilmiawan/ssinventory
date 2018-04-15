package model

//Inventory is the model for stock inventory database
type Inventory struct {
	SKU         string `json:"sku"`
	Name        string `json:"name"`
	Amount      int    `json:"amount,string"`
	AvgPrice    int    `json:"avgprice,string"`
	CreatedDate string `json:"created_date,omitempty"`
	UpdatedDate string `json:"updated_daet,omitempty"`
}
