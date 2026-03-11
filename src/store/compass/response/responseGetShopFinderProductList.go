package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// ShopFinderProductItem 带货达人商品列表项
type ShopFinderProductItem struct {
	ProductID   string `json:"product_id,omitempty"`   // 商品ID
	ProductName string `json:"product_name,omitempty"` // 商品名称
	SkuID       string `json:"sku_id,omitempty"`       // skuID
	PayAmount   int64  `json:"pay_amount,omitempty"`   // 支付金额(分)
	PayCount    int64  `json:"pay_count,omitempty"`    // 支付件数
}

type ResponseGetShopFinderProductList struct {
	response.ResponseStore
	ProductList []ShopFinderProductItem `json:"product_list,omitempty"`
	NextKey     string                  `json:"next_key,omitempty"`
	HasMore     bool                    `json:"has_more,omitempty"`
}
