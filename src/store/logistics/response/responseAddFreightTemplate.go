package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

// ResponseAddFreightTemplate 增加运费模版响应
type ResponseAddFreightTemplate struct {
	response.ResponseStore
	// 模板ID
	TemplateID string `json:"template_id"`
}
