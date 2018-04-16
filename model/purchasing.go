package model

//Purchasing is the model for stock purchasing database
type Purchasing struct {
	ID             int    `json:"id,string,omitempty"`
	SKU            string `json:"sku"`
	PurchasingDate string `json:"purchasing_date,omitempty"`
	ReqAmount      int    `json:"req_amount,string"`
	RecAmount      int    `json:"rec_amount,string"`
	Price          int    `json:"price,string"`
	Total          int    `json:"total,string"`
	ReceiptNo      string `json:"receipt_no"`
	Notes          string `json:"notes"`
}
