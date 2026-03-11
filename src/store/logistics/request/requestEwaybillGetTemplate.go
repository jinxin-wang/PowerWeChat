package request

// RequestEwaybillGetTemplate 获取面单模板信息请求
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybillgettemplate.html
type RequestEwaybillGetTemplate struct {
	// 模板编码
	TemplateCode string `json:"template_code"`
}
