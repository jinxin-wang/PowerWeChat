package request

type RequestGetSKUChangeList struct {
	OrderID  string `json:"order_id"`
	PageSize int    `json:"page_size,omitempty"`
	NextKey  string `json:"next_key,omitempty"`
}

type RequestAcceptSKUChange struct {
	OrderID     string `json:"order_id"`
	SKUChangeID string `json:"sku_change_id"`
}

type RequestRejectSKUChange struct {
	OrderID     string `json:"order_id"`
	SKUChangeID string `json:"sku_change_id"`
	Reason      string `json:"reason"`
}
