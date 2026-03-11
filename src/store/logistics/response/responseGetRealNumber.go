package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// ResponseGetRealNumber 根据运单号获取真实手机号响应
type ResponseGetRealNumber struct {
	response.ResponseStore
	// 手机号码
	PhoneNumber string `json:"phone_number"`
}
