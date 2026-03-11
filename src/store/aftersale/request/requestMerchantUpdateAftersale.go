package request

type RequestMerchantUpdateAftersale struct {
	AftersaleID  string `json:"aftersale_id"`
	RefundAmount int    `json:"refund_amount,omitempty"`
}
