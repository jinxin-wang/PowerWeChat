package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// ResponseEwaybillCreateTemplate 新增面单模板响应
type ResponseEwaybillCreateTemplate struct {
	response.ResponseStore
	// 模板ID
	TemplateID string `json:"template_id"`
}
