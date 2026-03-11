package request

// RequestGetPromoter 获取达人详情信息请求
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/promoter/api_getpromoter.html
type RequestGetPromoter struct {
	FinderID   string `json:"finder_id,omitempty"`   // 视频号finder_id，待废除后续以promoter_id为准
	PromoterID string `json:"promoter_id,omitempty"` // 达人带货id，和finder_id二选一
}
