package request

// RequestAddAddress 添加地址请求
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/addlogisticsaddress.html
type RequestAddAddress struct {
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
