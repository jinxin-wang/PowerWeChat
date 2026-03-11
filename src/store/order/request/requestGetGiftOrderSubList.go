package request

type RequestGetGiftOrderSubList struct {
	OrderID  string `json:"order_id"`
	PageSize int    `json:"page_size,omitempty"`
	NextKey  string `json:"next_key,omitempty"`
}
