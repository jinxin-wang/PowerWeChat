package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// ResponseEwaybillGetTemplateByID 根据模板ID获取面单模板信息响应
type ResponseEwaybillGetTemplateByID struct {
	response.ResponseStore
	// 模板信息
	TemplateInfo EwaybillTemplateInfo `json:"template_info"`
}
