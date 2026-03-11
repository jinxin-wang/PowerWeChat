package request

// RequestGetVirtualNumber 根据运单号获取虚拟手机号请求
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/getvirtualnumber.html
type RequestGetVirtualNumber struct {
	// 电子面单ID
	WaybillID string `json:"waybill_id"`
}
