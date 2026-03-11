package request

type RequestGetItemList struct {
	PageSize int    `json:"page_size,omitempty"` // 每页数量
	NextKey  string `json:"next_key,omitempty"` // 分页拉取标识
	Status   int    `json:"status,omitempty"`  // 商品状态
}
