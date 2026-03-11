package request

// RequestEwaybillGetOrder 查询面单详情请求
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybillgetorder.html
type RequestEwaybillGetOrder struct {
	// 电子面单ID
	WaybillID string `json:"waybill_id"`
}
