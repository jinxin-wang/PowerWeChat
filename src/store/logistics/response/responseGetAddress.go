package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// ResponseGetAddress 获取地址详情响应
type ResponseGetAddress struct {
	response.ResponseStore
	// 地址信息
	AddressInfo AddressInfo `json:"address_info"`
}

// AddressInfo 地址信息
type AddressInfo struct {
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
