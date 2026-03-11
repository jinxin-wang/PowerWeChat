package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ProductInfo struct {
	ProductID string `json:"product_id"`
	SkuID     string `json:"sku_id"`
	Title     string `json:"title"`
	ThumbImg  string `json:"thumb_img"`
	SalePrice int    `json:"sale_price"`
	Quantity  int    `json:"quantity"`
}

type AftersaleInfo struct {
	AftersaleID  string       `json:"aftersale_id"`
	OrderID      string       `json:"order_id"`
	Status       int          `json:"status"`
	Type         int          `json:"type"`
	Reason       string       `json:"reason"`
	RefundAmount int          `json:"refund_amount"`
	CreateTime   string       `json:"create_time"`
	UpdateTime   string       `json:"update_time"`
	ProductInfo  *ProductInfo `json:"product_info,omitempty"`
	OpenID       string       `json:"openid,omitempty"`
}

type ResponseGetAftersaleList struct {
	ResponseStore response.ResponseStore `json:"-"`
	AftersaleList []AftersaleInfo        `json:"aftersale_list"`
	NextKey       string                 `json:"next_key"`
	HasMore       bool                   `json:"has_more"`
}
