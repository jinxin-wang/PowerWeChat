package request

// RequestUpdatePromoter 编辑达人请求
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/promoter/api_updpromoter.html
type RequestUpdatePromoter struct {
	FinderID   string `json:"finder_id,omitempty"`   // 视频号finder_id，待废除后续以promoter_id为准
	PromoterID string `json:"promoter_id,omitempty"` // 达人带货id，和finder_id二选一
	Type       int    `json:"type,omitempty"`        // 操作枚举值
}
