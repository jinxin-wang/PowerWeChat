package request

// RequestEwaybillCancelOrder 电子面单取消下单请求
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybillcancelorder.html
type RequestEwaybillCancelOrder struct {
	// 电子面单ID
	WaybillID string `json:"waybill_id"`
	// 订单ID
	OrderID string `json:"order_id,omitempty"`
}
