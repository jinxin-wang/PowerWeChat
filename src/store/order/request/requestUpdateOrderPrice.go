package request

type RequestUpdateOrderPrice struct {
	OrderID    string `json:"order_id"`
	ChangeType int    `json:"change_type"`
	Price      int    `json:"price"`
	Remark     string `json:"remark,omitempty"`
}
