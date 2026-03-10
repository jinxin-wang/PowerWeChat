package order

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

// OrderInfo 订单信息
type OrderInfo struct {
	OrderID      string       `json:"order_id"`
	Status       int          `json:"status"`
	CreateTime   string       `json:"create_time"`
	UpdateTime   string       `json:"update_time"`
	OpenID       string       `json:"openid"`
	UnionID      string       `json:"unionid,omitempty"`
	OrderDetail  OrderDetail  `json:"order_detail"`
	AddressInfo  AddressInfo  `json:"address_info"`
	DeliveryInfo DeliveryInfo `json:"delivery_info,omitempty"`
}

// OrderDetail 订单详情
type OrderDetail struct {
	ProductInfos []ProductInfo `json:"product_infos"`
	PriceInfo    PriceInfo     `json:"price_info"`
}

// ProductInfo 商品信息
type ProductInfo struct {
	ProductID   string `json:"product_id"`
	SkuID       string `json:"sku_id"`
	ThumbImg    string `json:"thumb_img"`
	SalePrice   int    `json:"sale_price"`
	MarketPrice int    `json:"market_price"`
	RealPrice   int    `json:"real_price"`
	ProductCnt  int    `json:"product_cnt"`
	Title       string `json:"title"`
}

// PriceInfo 价格信息
type PriceInfo struct {
	ProductPrice    int `json:"product_price"`
	OrderPrice      int `json:"order_price"`
	Freight         int `json:"freight"`
	DiscountedPrice int `json:"discounted_price"`
}

// AddressInfo 地址信息
type AddressInfo struct {
	UserName     string `json:"user_name"`
	PostalCode   string `json:"postal_code"`
	ProvinceName string `json:"province_name"`
	CityName     string `json:"city_name"`
	CountyName   string `json:"county_name"`
	DetailInfo   string `json:"detail_info"`
	TelNumber    string `json:"tel_number"`
}

// DeliveryInfo 物流信息
type DeliveryInfo struct {
	DeliveryID   string `json:"delivery_id"`
	WaybillID    string `json:"waybill_id"`
	DeliveryName string `json:"delivery_name"`
}

// ResponseGetOrderList 获取订单列表响应
type ResponseGetOrderList struct {
	response.ResponseStore
	OrderList []OrderInfo `json:"order_list"`
	NextKey   string      `json:"next_key"`
	HasMore   bool        `json:"has_more"`
}

// ResponseGetOrder 获取订单详情响应
type ResponseGetOrder struct {
	response.ResponseStore
	Order OrderInfo `json:"order"`
}

// ResponseSearchOrder 订单搜索响应
type ResponseSearchOrder struct {
	response.ResponseStore
	OrderList []OrderInfo `json:"order_list"`
	NextKey   string      `json:"next_key"`
	HasMore   bool        `json:"has_more"`
}

// ResponseUpdateOrder 修改订单响应（通用）
type ResponseUpdateOrder struct {
	response.ResponseStore
}

// ResponseDecodeSensitiveInfo 解密订单中的详细收货信息响应
type ResponseDecodeSensitiveInfo struct {
	response.ResponseStore
	ReceiverName string `json:"receiver_name"`
	TelNumber    string `json:"tel_number"`
	DetailInfo   string `json:"detail_info"`
}
