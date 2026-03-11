package request

type RequestAddGiftOrderNote struct {
	OrderID string `json:"order_id"`
	Note    string `json:"note"`
}
