package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// ResponseGetVIPUserScore 获取用户积分响应
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/vip/api_getvipuserscore.html
type ResponseGetVIPUserScore struct {
	response.ResponseStore
	Score      int `json:"score"`       // 当前可用积分
	TotalScore int `json:"total_score"` // 累计积分
}
