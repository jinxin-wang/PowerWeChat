package request

// RequestDeleteItem 删除联盟商品请求
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/item/api_deleteitem.html
type RequestDeleteItem struct {
	ProductID string `json:"product_id,omitempty"` // 商品ID
}
