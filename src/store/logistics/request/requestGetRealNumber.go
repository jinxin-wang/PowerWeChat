package request

// RequestGetRealNumber 根据运单号获取真实手机号请求
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/getrealnumber.html
type RequestGetRealNumber struct {
	// 电子面单ID
	WaybillID string `json:"waybill_id"`
}
