package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// ResponseGetDeliveryCompanyList 获取快递公司列表响应（旧）
type ResponseGetDeliveryCompanyList struct {
	response.ResponseStore
	// 快递公司列表
	CompanyList []DeliveryCompany `json:"company_list"`
}
