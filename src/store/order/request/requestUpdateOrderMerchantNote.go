package request

type RequestUpdateOrderMerchantNote struct {
	OrderID string `json:"order_id"`
	Note    string `json:"note"`
}
