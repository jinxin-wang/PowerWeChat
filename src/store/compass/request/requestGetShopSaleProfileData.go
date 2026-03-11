package request

type RequestGetShopSaleProfileData struct {
	StartTime int64  `json:"start_time,omitempty"` // 开始时间，秒级时间戳
	EndTime   int64  `json:"end_time,omitempty"`   // 结束时间，秒级时间戳
	Type      string `json:"type,omitempty"`       // 数据类型：age-年龄分布，gender-性别分布，province-地域分布
}
