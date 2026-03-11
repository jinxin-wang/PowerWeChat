package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// ResponseGetVirtualNumber 根据运单号获取虚拟手机号响应
type ResponseGetVirtualNumber struct {
	response.ResponseStore
	// 虚拟号码
	VirtualNumber string `json:"virtual_number"`
}
