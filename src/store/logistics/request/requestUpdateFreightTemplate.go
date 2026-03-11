package request

// RequestUpdateFreightTemplate 更新运费模版请求
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/updatefreighttemplate.html
type RequestUpdateFreightTemplate struct {
	// 模板ID
	TemplateID string `json:"template_id"`
	// 模板名称
	Name string `json:"name"`
	// 运费信息
	FreightInfo FreightInfo `json:"freight_info"`
}
