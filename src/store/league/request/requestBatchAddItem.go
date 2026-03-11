package request

// RequestBatchAddItem 批量新增联盟商品请求
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/item/api_batchadditem.html
type RequestBatchAddItem struct {
	ProductIDs []string `json:"product_ids,omitempty"` // 商品ID列表
}
