package request

// RequestGetUserInfo 获取用户信息请求
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/vip/api_getuserinfo.html
type RequestGetUserInfo struct {
	Openid string `json:"openid"` // 用户openid
}
