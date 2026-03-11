package request

// RequestGetItemList 获取联盟商品推广列表请求
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/item/api_getitemlist.html
type RequestGetItemList struct {
	PageSize int    `json:"page_size,omitempty"` // 每页数量
	NextKey  string `json:"next_key,omitempty"`  // 分页拉取标识
	Status   int    `json:"status,omitempty"`    // 商品状态
}
