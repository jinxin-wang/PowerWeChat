package request

type RequestGetShopFinderList struct {
	PageSize  int    `json:"page_size,omitempty"`  // 每页获取记录数，默认10，最大100
	NextKey   string `json:"next_key,omitempty"`   // 翻页标记
	StartTime int64  `json:"start_time,omitempty"` // 开始时间，秒级时间戳
	EndTime   int64  `json:"end_time,omitempty"`   // 结束时间，秒级时间戳
}
