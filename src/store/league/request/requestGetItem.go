package request

// RequestGetItem 获取联盟商品详情请求
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/item/api_getitem.html
type RequestGetItem struct {
	Type      int `json:"type,omitempty"`       // 获取商品推广类别 1-普通推广商品 2-专属推广商品
	ProductID int `json:"product_id,omitempty"` // 商品id，type为普通推广商品时必填
	InfoID    int `json:"info_id,omitempty"`    // 特殊推广商品计划id，type为特殊推广商品时必填
}
