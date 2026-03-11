package request

// RequestEwaybillCreateOrder 电子面单取号请求
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybillcreateorder.html
type RequestEwaybillCreateOrder struct {
	// 快递公司ID
	DeliveryID string `json:"delivery_id"`
	// 订单ID
	OrderID string `json:"order_id"`
	// 发货人信息
	SenderInfo EwaybillAddressInfo `json:"sender_info"`
	// 收货人信息
	ReceiverInfo EwaybillAddressInfo `json:"receiver_info"`
	// 包裹信息
	PackageInfo EwaybillPackageInfo `json:"package_info"`
}
