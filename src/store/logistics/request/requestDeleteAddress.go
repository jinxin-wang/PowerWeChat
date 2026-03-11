package request

// RequestDeleteAddress 删除地址请求
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/deletelogisticsaddress.html
type RequestDeleteAddress struct {
	// 地址ID
	AddressID string `json:"address_id"`
}
