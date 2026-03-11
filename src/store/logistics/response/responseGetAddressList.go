package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// ResponseGetAddressList 获取地址列表响应
type ResponseGetAddressList struct {
	response.ResponseStore
	// 地址列表
	AddressList []AddressInfo `json:"address_list"`
	// 分页游标
	NextKey string `json:"next_key"`
	// 是否还有更多
	HasMore bool `json:"has_more"`
}
