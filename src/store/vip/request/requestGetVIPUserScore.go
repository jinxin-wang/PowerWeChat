package request

// RequestGetVIPUserScore 获取用户积分请求
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/vip/api_getvipuserscore.html
type RequestGetVIPUserScore struct {
	Openid string `json:"openid"` // 用户openid
}
