package request

// RequestEwaybillDelTemplate 删除面单模版请求
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybilldeltemplate.html
type RequestEwaybillDelTemplate struct {
	// 模板ID
	TemplateID string `json:"template_id"`
}
