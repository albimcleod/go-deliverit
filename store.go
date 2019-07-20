package godeliverit

//Store defines an Store from Deliverit POS
type Store struct {
	ID   int    `json:"store_id"`
	Name string `json:"store_name"`
}

type Stores []Store
