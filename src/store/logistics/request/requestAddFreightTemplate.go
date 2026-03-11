package request

// RequestAddFreightTemplate 增加运费模版请求
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/addfreighttemplate.html
type RequestAddFreightTemplate struct {
	// 模板名称
	Name string `json:"name"`
	// 运费信息
	FreightInfo FreightInfo `json:"freight_info"`
}

// FreightInfo 运费信息
type FreightInfo struct {
	// 计费类型：1-按件计费，2-按重量计费
	ValuationType int `json:"valuation_type"`
	// 运送方式列表
	DeliveryMethods []DeliveryMethod `json:"delivery_methods"`
}

// DeliveryMethod 运送方式
type DeliveryMethod struct {
	// 运送方式：1-快递，2-同城配送
	DeliveryMethodID int `json:"delivery_method_id"`
	// 是否默认运送方式
	IsDefault bool `json:"is_default"`
	// 首件数量
	FirstVal int `json:"first_val"`
	// 首费金额（分）
	FirstFee int `json:"first_fee"`
	// 续件数量
	AddVal int `json:"add_val"`
	// 续费金额（分）
	AddFee int `json:"add_fee"`
	// 运送范围列表
	AddressList []DeliveryAddress `json:"address_list"`
}

// DeliveryAddress 运送范围
type DeliveryAddress struct {
	// 是否默认配送范围
	IsDefault bool `json:"is_default"`
	// 可配送区域列表
	DeliveryAddressIDList []string `json:"delivery_address_id_list"`
}
