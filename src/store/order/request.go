package order

// RequestGetOrderList 获取订单列表请求
type RequestGetOrderList struct {
	PageSize  int    `json:"page_size,omitempty"`
	NextKey   string `json:"next_key,omitempty"`
	Status    int    `json:"status,omitempty"`
	StartTime string `json:"start_time,omitempty"`
	EndTime   string `json:"end_time,omitempty"`
}

// RequestGetOrder 获取订单详情请求
type RequestGetOrder struct {
	OrderID string `json:"order_id"`
}

// RequestSearchOrder 订单搜索请求
type RequestSearchOrder struct {
	Keyword  string `json:"keyword"`
	PageSize int    `json:"page_size,omitempty"`
	NextKey  string `json:"next_key,omitempty"`
}

// RequestUpdateOrderPrice 修改订单价格请求
type RequestUpdateOrderPrice struct {
	OrderID    string `json:"order_id"`
	ChangeType int    `json:"change_type"`
	Price      int    `json:"price"`
	Remark     string `json:"remark,omitempty"`
}

// RequestUpdateOrderMerchantNote 修改订单备注请求
type RequestUpdateOrderMerchantNote struct {
	OrderID string `json:"order_id"`
	Note    string `json:"note"`
}

// RequestUpdateOrderAddress 修改订单地址请求
type RequestUpdateOrderAddress struct {
	OrderID      string `json:"order_id"`
	ReceiverName string `json:"receiver_name"`
	DetailInfo   string `json:"detail_info"`
	TelNumber    string `json:"tel_number"`
	PostalCode   string `json:"postal_code,omitempty"`
	ProvinceName string `json:"province_name"`
	CityName     string `json:"city_name"`
	CountyName   string `json:"county_name"`
}

// RequestUpdateOrderDelivery 修改物流信息请求
type RequestUpdateOrderDelivery struct {
	OrderID        string `json:"order_id"`
	DeliveryID     string `json:"delivery_id,omitempty"`
	WaybillID      string `json:"waybill_id,omitempty"`
	DeliveryName   string `json:"delivery_name,omitempty"`
	DeliveryRemark string `json:"delivery_remark,omitempty"`
}

// RequestAcceptOrderAddressModify 同意用户修改收货地址申请请求
type RequestAcceptOrderAddressModify struct {
	OrderID string `json:"order_id"`
}

// RequestRejectOrderAddressModify 拒绝用户修改收货地址申请请求
type RequestRejectOrderAddressModify struct {
	OrderID string `json:"order_id"`
	Reason  string `json:"reason"`
}

// RequestDecodeSensitiveInfo 解密订单中的详细收货信息请求
type RequestDecodeSensitiveInfo struct {
	OrderID       string `json:"order_id"`
	EncryptedData string `json:"encrypted_data"`
}
