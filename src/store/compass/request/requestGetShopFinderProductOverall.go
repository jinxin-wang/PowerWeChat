package request

type RequestGetShopFinderProductOverall struct {
	FinderID  string `json:"finder_id,omitempty"`  // 视频号ID
	ProductID string `json:"product_id,omitempty"` // 商品ID
	StartTime int64  `json:"start_time,omitempty"` // 开始时间，秒级时间戳
	EndTime   int64  `json:"end_time,omitempty"`   // 结束时间，秒级时间戳
}
