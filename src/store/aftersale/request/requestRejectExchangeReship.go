package request

type RequestRejectExchangeReship struct {
	AftersaleID string `json:"aftersale_id"`
	Reason      string `json:"reason"`
}
