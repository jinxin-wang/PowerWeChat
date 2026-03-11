package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

// ResponseEwaybillGetTemplate 获取面单模板信息响应
type ResponseEwaybillGetTemplate struct {
	response.ResponseStore
	// 模板信息
	TemplateInfo EwaybillTemplateInfo `json:"template_info"`
}

// EwaybillTemplateInfo 电子面单模板信息
type EwaybillTemplateInfo struct {
	// 模板ID
	TemplateID string `json:"template_id"`
	// 模板名称
	TemplateName string `json:"template_name"`
	// 模板编码
	TemplateCode string `json:"template_code"`
	// 打印方向：1-横版，2-竖版
	Orientation int `json:"orientation"`
	// 缩放比例
	ScaleType int `json:"scale_type"`
	// 是否打印商品信息
	PrintProductInfo bool `json:"print_product_info"`
}
