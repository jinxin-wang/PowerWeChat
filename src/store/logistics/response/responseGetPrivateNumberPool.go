package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// ResponseGetPrivateNumberPool 获取虚拟号码池响应
type ResponseGetPrivateNumberPool struct {
	response.ResponseStore
	// 号码池信息
	NumberPool PrivateNumberPool `json:"number_pool"`
}

// PrivateNumberPool 虚拟号码池
type PrivateNumberPool struct {
	// 号码池ID
	PoolID string `json:"pool_id"`
	// 可用号码数量
	AvailableCount int `json:"available_count"`
	// 已使用号码数量
	UsedCount int `json:"used_count"`
}
