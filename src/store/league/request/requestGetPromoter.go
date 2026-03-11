package request

type RequestGetPromoter struct {
	FinderID   string `json:"finder_id,omitempty"`   // 视频号finder_id，待废除后续以promoter_id为准
	PromoterID string `json:"promoter_id,omitempty"` // 达人带货id，和finder_id二选一
}
