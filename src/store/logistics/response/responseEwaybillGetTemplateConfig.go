package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

// ResponseEwaybillGetTemplateConfig 获取面单标准模板配置响应
type ResponseEwaybillGetTemplateConfig struct {
	response.ResponseStore
	// 模板配置列表
	TemplateConfigs []TemplateConfig `json:"template_configs"`
}

// TemplateConfig 模板配置
type TemplateConfig struct {
	// 模板ID
	TemplateCode string `json:"template_code"`
	// 模板名称
	TemplateName string `json:"template_name"`
	// 模板尺寸
	Size string `json:"size"`
	// 模板描述
	Description string `json:"description"`
}
