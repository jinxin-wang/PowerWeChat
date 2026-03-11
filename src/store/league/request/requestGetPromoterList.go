package request

// RequestGetPromoterList 获取商店达人列表请求
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/promoter/api_getpromoterlist.html
type RequestGetPromoterList struct {
	PageSize int    `json:"page_size,omitempty"` // 每页数量
	NextKey  string `json:"next_key,omitempty"`  // 分页拉取标识
}
