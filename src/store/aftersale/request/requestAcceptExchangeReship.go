package request

type RequestAcceptExchangeReship struct {
	AftersaleID string `json:"aftersale_id"`
	DeliveryID  string `json:"delivery_id,omitempty"`
	WaybillID   string `json:"waybill_id,omitempty"`
}
