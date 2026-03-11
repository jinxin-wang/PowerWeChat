package request

// RequestEwaybillAddSubOrder 电子面单子件追加请求
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybilladdsuborder.html
type RequestEwaybillAddSubOrder struct {
	// 电子面单ID
	WaybillID string `json:"waybill_id"`
	// 子件信息
	SubOrderInfo EwaybillSubOrderInfo `json:"sub_order_info"`
}

// EwaybillSubOrderInfo 子件信息
type EwaybillSubOrderInfo struct {
	// 子件单号
	SubWaybillID string `json:"sub_waybill_id"`
}
