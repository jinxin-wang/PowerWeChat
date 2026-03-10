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

// SubOrderInfo 子订单信息
type SubOrderInfo struct {
	SubOrderID string `json:"sub_order_id"`
	Status     int    `json:"status"`
	OpenID     string `json:"openid"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

// ResponseGetGiftOrderSubList 获取礼品订单子单列表响应
type ResponseGetGiftOrderSubList struct {
	response.ResponseStore
	SubOrderList []SubOrderInfo `json:"sub_order_list"`
	NextKey      string         `json:"next_key"`
	HasMore      bool           `json:"has_more"`
}

// SKUChangeInfo SKU变更信息
type SKUChangeInfo struct {
	SKUChangeID string `json:"sku_change_id"`
	OrderID     string `json:"order_id"`
	ProductID   string `json:"product_id"`
	OldSkuID    string `json:"old_sku_id"`
	NewSkuID    string `json:"new_sku_id"`
	CreateTime  string `json:"create_time"`
	Status      int    `json:"status"`
}

// ResponseGetSKUChangeList 获取待发货SKU变更列表响应
type ResponseGetSKUChangeList struct {
	response.ResponseStore
	SKUChangeList []SKUChangeInfo `json:"sku_change_list"`
	NextKey       string          `json:"next_key"`
	HasMore       bool            `json:"has_more"`
}

// ResponseApplyRealNumber 申请查看真实号码响应
type ResponseApplyRealNumber struct {
	response.ResponseStore
	ApplyID string `json:"apply_id"`
}

// ResponseGetRealNumberStatus 查询真实号码审核状态响应
type ResponseGetRealNumberStatus struct {
	response.ResponseStore
	Status      int    `json:"status"`
	PhoneNumber string `json:"phone_number"`
	AuditResult string `json:"audit_result,omitempty"`
}

// ResponseGetPhoneStatus 获取店铺手机号核验状态响应
type ResponseGetPhoneStatus struct {
	response.ResponseStore
	Status int    `json:"status"`
	Phone  string `json:"phone,omitempty"`
}
