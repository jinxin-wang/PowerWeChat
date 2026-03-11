package request

// RequestGetUserList 获取用户列表请求
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/vip/api_getuserlist.html
type RequestGetUserList struct {
	PageSize int    `json:"page_size"` // 分页大小，最大100
	NextKey  string `json:"next_key"`  // 分页游标，第一页为空
}
