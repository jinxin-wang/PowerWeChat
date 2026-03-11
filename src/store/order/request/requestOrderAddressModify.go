package request

type RequestAcceptOrderAddressModify struct {
	OrderID string `json:"order_id"`
}

type RequestRejectOrderAddressModify struct {
	OrderID string `json:"order_id"`
	Reason  string `json:"reason"`
}
