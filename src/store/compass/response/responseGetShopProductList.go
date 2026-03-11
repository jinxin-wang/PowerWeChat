package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// ShopProductListItem 商品列表项
type ShopProductListItem struct {
	ProductID   string `json:"product_id,omitempty"`   // 商品ID
	ProductName string `json:"product_name,omitempty"` // 商品名称
	PayAmount   int64  `json:"pay_amount,omitempty"`   // 支付金额(分)
	PayCount    int64  `json:"pay_count,omitempty"`    // 支付件数
	Uv          int64  `json:"uv,omitempty"`           // 访客数
}

type ResponseGetShopProductList struct {
	response.ResponseStore
	ProductList []ShopProductListItem `json:"product_list,omitempty"`
	NextKey     string                `json:"next_key,omitempty"`
	HasMore     bool                  `json:"has_more,omitempty"`
}
