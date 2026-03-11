package request

// RequestUpdateAddress 更新地址请求
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/updatelogisticsaddress.html
type RequestUpdateAddress struct {
	// 地址ID
	AddressID string `json:"address_id"`
	// 收件人姓名
	ReceiverName string `json:"receiver_name"`
	// 联系电话
	Tel string `json:"tel"`
	// 省份
	Province string `json:"province"`
	// 城市
	City string `json:"city"`
	// 区县
	District string `json:"district"`
	// 详细地址
	Detail string `json:"detail"`
}
