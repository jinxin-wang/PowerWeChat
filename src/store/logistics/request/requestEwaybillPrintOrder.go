package request

// RequestEwaybillPrintOrder 打印成功通知请求
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybillprintorder.html
type RequestEwaybillPrintOrder struct {
	// 电子面单ID
	WaybillID string `json:"waybill_id"`
}
