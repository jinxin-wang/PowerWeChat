package request

type RequestUpdateOrderDelivery struct {
	OrderID        string `json:"order_id"`
	DeliveryID     string `json:"delivery_id,omitempty"`
	WaybillID      string `json:"waybill_id,omitempty"`
	DeliveryName   string `json:"delivery_name,omitempty"`
	DeliveryRemark string `json:"delivery_remark,omitempty"`
}
