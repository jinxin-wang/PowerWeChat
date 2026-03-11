package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// FinderProductOverall 带货达人商品详情数据
type FinderProductOverall struct {
	PayAmount     int64 `json:"pay_amount,omitempty"`     // 支付金额(分)
	PayCount      int64 `json:"pay_count,omitempty"`      // 支付件数
	Uv            int64 `json:"uv,omitempty"`             // 访客数
	AddCartCount  int64 `json:"add_cart_count,omitempty"` // 加购件数
	ExposureCount int64 `json:"exposure_count,omitempty"` // 曝光次数
}

type ResponseGetShopFinderProductOverall struct {
	response.ResponseStore
	ProductOverall FinderProductOverall `json:"product_overall,omitempty"`
}
