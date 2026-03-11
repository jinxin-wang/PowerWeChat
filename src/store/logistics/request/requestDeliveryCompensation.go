package request

// RequestDeliveryCompensation 订单补发货请求
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/deliverycompensation.html
type RequestDeliveryCompensation struct {
	// 订单ID
	OrderID string `json:"order_id"`
	// 发货信息列表
	DeliveryList []DeliveryInfo `json:"delivery_list"`
}
