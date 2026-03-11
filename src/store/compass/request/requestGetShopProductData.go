package request

type RequestGetShopProductData struct {
	ProductID string `json:"product_id,omitempty"` // 商品ID
	StartTime int64  `json:"start_time,omitempty"` // 开始时间，秒级时间戳
	EndTime   int64  `json:"end_time,omitempty"`   // 结束时间，秒级时间戳
}
