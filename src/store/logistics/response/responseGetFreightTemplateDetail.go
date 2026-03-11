package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// ResponseGetFreightTemplateDetail 查询运费模版详情响应
type ResponseGetFreightTemplateDetail struct {
	response.ResponseStore
	// 运费模板信息
	TemplateInfo FreightTemplateInfo `json:"template_info"`
}

// FreightTemplateInfo 运费模板信息
type FreightTemplateInfo struct {
	// 模板ID
	TemplateID string `json:"template_id"`
	// 模板名称
	Name string `json:"name"`
	// 运费信息
	FreightInfo FreightInfoResponse `json:"freight_info"`
}

// FreightInfoResponse 运费信息
type FreightInfoResponse struct {
	// 计费类型：1-按件计费，2-按重量计费
	ValuationType int `json:"valuation_type"`
	// 运送方式列表
	DeliveryMethods []DeliveryMethodResponse `json:"delivery_methods"`
}

// DeliveryMethodResponse 运送方式
type DeliveryMethodResponse struct {
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
	AddressList []DeliveryAddressResponse `json:"address_list"`
}

// DeliveryAddressResponse 运送范围
type DeliveryAddressResponse struct {
	// 是否默认配送范围
	IsDefault bool `json:"is_default"`
	// 可配送区域列表
	DeliveryAddressIDList []string `json:"delivery_address_id_list"`
}
