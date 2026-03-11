package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// UserInfo 用户信息结构
type UserInfo struct {
	Openid      string `json:"openid"`       // 用户openid
	Unionid     string `json:"unionid"`      // 用户unionid
	Nickname    string `json:"nickname"`     // 用户昵称
	Headimgurl  string `json:"headimgurl"`   // 用户头像
	Gender      int    `json:"gender"`       // 用户性别 0:未知 1:男 2:女
	Country     string `json:"country"`      // 国家
	Province    string `json:"province"`     // 省份
	City        string `json:"city"`         // 城市
	PhoneNumber string `json:"phone_number"` // 手机号
}

// ResponseGetUserInfo 获取用户信息响应
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/vip/api_getuserinfo.html
type ResponseGetUserInfo struct {
	response.ResponseStore
	UserInfo UserInfo `json:"user_info"` // 用户信息
}
