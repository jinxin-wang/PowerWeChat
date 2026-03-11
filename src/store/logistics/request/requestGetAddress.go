package request

// RequestGetAddress 获取地址详情请求
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/getlogisticsaddress.html
type RequestGetAddress struct {
	// 地址ID
	AddressID string `json:"address_id"`
}
