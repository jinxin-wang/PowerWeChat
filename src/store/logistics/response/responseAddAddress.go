package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

// ResponseAddAddress 添加地址响应
type ResponseAddAddress struct {
	response.ResponseStore
	// 地址ID
	AddressID string `json:"address_id"`
}
