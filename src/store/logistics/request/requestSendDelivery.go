package request

// RequestSendDelivery 订单发货请求
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/senddelivery.html
type RequestSendDelivery struct {
	// 订单ID
	OrderID string `json:"order_id"`
	// 发货信息列表
	DeliveryList []DeliveryInfo `json:"delivery_list"`
}

// DeliveryInfo 发货信息
type DeliveryInfo struct {
	// 快递公司ID
	DeliveryID string `json:"delivery_id"`
	// 快递单号
	WaybillID string `json:"waybill_id"`
	// 地址ID（使用电子面单发货时不需要）
	AddressID string `json:"address_id,omitempty"`
	// 电子面单发货信息
	EwaybillInfo *EwaybillDeliveryInfo `json:"ewaybill_info,omitempty"`
}

// EwaybillDeliveryInfo 电子面单发货信息
type EwaybillDeliveryInfo struct {
	// 电子面单号
	WaybillID string `json:"waybill_id"`
	// 电子面单ID
	EwaybillOrderID string `json:"ewaybill_order_id"`
	// 快递公司ID
	DeliveryID string `json:"delivery_id"`
}
