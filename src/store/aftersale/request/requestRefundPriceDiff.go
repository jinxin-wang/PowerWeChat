package request

type RequestRefundPriceDiff struct {
	OrderID         string `json:"order_id"`
	PriceDiffAmount int    `json:"price_diff_amount"`
}
