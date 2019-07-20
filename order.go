package godeliverit

import "time"

//Store defines an Store from Deliverit POS
type Order struct {
	StoreID     string       `json:"StoreID"`
	OrderID     string       `json:"OrderID"`
	OrderDate   string       `json:"OrderDate"`
	InTime      string       `json:"InTime"`
	OrderDetail OrderDetails `json:"OrderDetail"`
	AmountPaid  string       `json:"AmountPaid"`
}

func (v *Order) CreatedAt() time.Time {
	return time.Now()
}

type Orders []Order
type OrderDetails []OrderDetail

type OrderDetail struct {
	OrderDetail string `json:"OrderDetail"`
	Qty         string `json:"Qty"`
	UnitSell    string `json:"UnitSell"`
}
