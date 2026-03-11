package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

// ResponseGetDeliveryCompanyListNew 获取快递公司列表响应（新）
type ResponseGetDeliveryCompanyListNew struct {
	response.ResponseStore
	// 快递公司列表
	CompanyList []DeliveryCompany `json:"company_list"`
}

// DeliveryCompany 快递公司
type DeliveryCompany struct {
	// 快递公司ID
	DeliveryID string `json:"delivery_id"`
	// 快递公司名称
	DeliveryName string `json:"delivery_name"`
}
