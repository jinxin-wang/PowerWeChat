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

// RequestUploadFreshInsurance 上传生鲜商品质检信息请求
type RequestUploadFreshInsurance struct {
	OrderID       string `json:"order_id"`
	InsuranceInfo string `json:"insurance_info"`
}

// RequestAddGiftOrderNote 添加礼品订单备注请求
type RequestAddGiftOrderNote struct {
	OrderID string `json:"order_id"`
	Note    string `json:"note"`
}

// RequestGetGiftOrderSubList 获取礼品订单子单列表请求
type RequestGetGiftOrderSubList struct {
	OrderID  string `json:"order_id"`
	PageSize int    `json:"page_size,omitempty"`
	NextKey  string `json:"next_key,omitempty"`
}

// RequestGetSKUChangeList 获取待发货SKU变更列表请求
type RequestGetSKUChangeList struct {
	OrderID  string `json:"order_id"`
	PageSize int    `json:"page_size,omitempty"`
	NextKey  string `json:"next_key,omitempty"`
}

// RequestAcceptSKUChange 接受SKU变更请求
type RequestAcceptSKUChange struct {
	OrderID     string `json:"order_id"`
	SKUChangeID string `json:"sku_change_id"`
}

// RequestRejectSKUChange 拒绝SKU变更请求
type RequestRejectSKUChange struct {
	OrderID     string `json:"order_id"`
	SKUChangeID string `json:"sku_change_id"`
	Reason      string `json:"reason"`
}

// RequestApplyRealNumber 申请查看真实号码请求
type RequestApplyRealNumber struct {
	OrderID string `json:"order_id"`
	Reason  string `json:"reason"`
}

// RequestGetRealNumberStatus 查询真实号码审核状态请求
type RequestGetRealNumberStatus struct {
	ApplyID string `json:"apply_id"`
}

// RequestReapplyVirtualNumber 重新申请虚拟号请求
type RequestReapplyVirtualNumber struct {
	OrderID string `json:"order_id"`
}

// RequestDelayVirtualNumber 延期虚拟号有效期请求
type RequestDelayVirtualNumber struct {
	OrderID   string `json:"order_id"`
	DelayDays int    `json:"delay_days"`
}

// RequestAddPhoneVerifyCode 添加手机号用于核验请求
type RequestAddPhoneVerifyCode struct {
	OrderID string `json:"order_id"`
	Phone   string `json:"phone"`
}

// RequestSendPhoneVerifyCode 获取短信验证码请求
type RequestSendPhoneVerifyCode struct {
	OrderID string `json:"order_id"`
}

// RequestGetPhoneStatus 获取店铺手机号核验状态请求
type RequestGetPhoneStatus struct {
	OrderID string `json:"order_id"`
}
