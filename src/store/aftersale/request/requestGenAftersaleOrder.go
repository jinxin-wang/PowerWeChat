package request

type RequestGenAftersaleOrder struct {
	OrderID   string `json:"order_id"`
	ProductID string `json:"product_id,omitempty"`
	Type      int    `json:"type"`
	Reason    string `json:"reason,omitempty"`
}
