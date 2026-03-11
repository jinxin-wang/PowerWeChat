package request

type RequestGetShopFinderProductList struct {
	FinderID string `json:"finder_id,omitempty"` // 视频号ID
	PageSize int    `json:"page_size,omitempty"` // 每页获取记录数，默认10，最大100
	NextKey  string `json:"next_key,omitempty"`  // 翻页标记
}
