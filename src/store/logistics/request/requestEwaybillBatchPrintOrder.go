package request

// RequestEwaybillBatchPrintOrder 批量打印通知请求
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybillbatchprintorder.html
type RequestEwaybillBatchPrintOrder struct {
	// 电子面单ID列表
	WaybillIDs []string `json:"waybill_ids"`
}
