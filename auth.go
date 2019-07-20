package godeliverit

//Auth defines an Auth from Deliverit POS
type Auth struct {
	Key         int    `json:"key"`
	Token       string `json:"token"`
	AccountName string `json:"account_name"`
}
