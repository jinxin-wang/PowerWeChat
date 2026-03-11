package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// ShopProductData 商品数据
type ShopProductData struct {
	ProductID     string `json:"product_id,omitempty"`     // 商品ID
	ProductName   string `json:"product_name,omitempty"`   // 商品名称
	PayAmount     int64  `json:"pay_amount,omitempty"`     // 支付金额(分)
	PayCount      int64  `json:"pay_count,omitempty"`      // 支付件数
	Uv            int64  `json:"uv,omitempty"`             // 访客数
	ExposureCount int64  `json:"exposure_count,omitempty"` // 曝光次数
	AddCartCount  int64  `json:"add_cart_count,omitempty"` // 加购件数
}

type ResponseGetShopProductData struct {
	response.ResponseStore
	ProductData ShopProductData `json:"product_data,omitempty"`
}
