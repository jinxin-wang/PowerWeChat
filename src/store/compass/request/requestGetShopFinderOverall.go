package request

type RequestGetShopFinderOverall struct {
	FinderID  string `json:"finder_id,omitempty"`  // 视频号ID
	StartTime int64  `json:"start_time,omitempty"` // 开始时间，秒级时间戳
	EndTime   int64  `json:"end_time,omitempty"`   // 结束时间，秒级时间戳
}
