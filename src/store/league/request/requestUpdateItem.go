package request

// ItemInfo 商品信息
type ItemInfo struct {
	ProductID int     `json:"product_id,omitempty"` // 商品ID
	Ratio     float64 `json:"ratio,omitempty"`      // 佣金比例
}

// ExclusiveInfo 专属推广信息
type ExclusiveInfo struct {
	FinderIDs     []string `json:"finder_ids,omitempty"`     // 专属达人列表
	PromoterIDs   []string `json:"promoter_ids,omitempty"`   // 专属达人ID列表
	StartTime     string   `json:"start_time,omitempty"`     // 开始时间
	EndTime       string   `json:"end_time,omitempty"`       // 结束时间
	DiscountRatio float64  `json:"discount_ratio,omitempty"` // 专属折扣比例
}

// RequestUpdateItem 更新联盟商品信息请求
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/item/api_upditem.html
type RequestUpdateItem struct {
	Type          int            `json:"type,omitempty"`           // 获取商品推广类别 1-普通推广商品 2-专属推广商品
	ProductID     int            `json:"product_id,omitempty"`     // 商品id，type为普通推广商品时必填
	InfoID        int            `json:"info_id,omitempty"`        // 特殊推广商品计划id，type为特殊推广商品时必填
	OperateType   int            `json:"operate_type,omitempty"`   // 操作类型
	Ratio         float64        `json:"ratio,omitempty"`          // 佣金比例
	ExclusiveInfo *ExclusiveInfo `json:"exclusive_info,omitempty"` // 专属信息
}
