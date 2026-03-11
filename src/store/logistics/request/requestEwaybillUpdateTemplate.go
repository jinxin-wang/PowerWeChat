package request

// RequestEwaybillUpdateTemplate 更新面单模版请求
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybillupdatetemplate.html
type RequestEwaybillUpdateTemplate struct {
	// 模板ID
	TemplateID string `json:"template_id"`
	// 模板配置信息
	ConfigInfo EwaybillConfigInfo `json:"config_info"`
}
