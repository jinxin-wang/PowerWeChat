package request

type RequestGetPromoterList struct {
	PageSize int    `json:"page_size,omitempty"` // 每页数量
	NextKey  string `json:"next_key,omitempty"`   // 分页拉取标识
}
