package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// UserListItem 用户列表项结构
type UserListItem struct {
	Openid  string `json:"openid"`  // 用户openid
	Unionid string `json:"unionid"` // 用户unionid
}

// ResponseGetUserList 获取用户列表响应
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/vip/api_getuserlist.html
type ResponseGetUserList struct {
	response.ResponseStore
	UserList []UserListItem `json:"user_list"` // 用户列表
	NextKey  string         `json:"next_key"`  // 分页游标
	HasMore  bool           `json:"has_more"`  // 是否还有更多数据
}
