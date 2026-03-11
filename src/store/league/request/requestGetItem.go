package request

type RequestGetItem struct {
	Type      int    `json:"type"`       // 获取商品推广类别
	ProductID int    `json:"product_id,omitempty"` // 商品id，type为普通推广商品时必填
	InfoID    int    `json:"info_id,omitempty"`   // 特殊推广商品计划id，type为特殊推广商品时必填
}
