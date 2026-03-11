package request

type RequestGetShopFinderAuthorizationList struct {
	PageSize int    `json:"page_size,omitempty"` // 每页获取记录数，默认10，最大100
	NextKey  string `json:"next_key,omitempty"`  // 翻页标记，从第一次返回结果中获取
}
