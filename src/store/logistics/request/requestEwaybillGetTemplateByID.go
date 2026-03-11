package request

// RequestEwaybillGetTemplateByID 根据模板ID获取面单模板信息请求
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybillgettemplatebyid.html
type RequestEwaybillGetTemplateByID struct {
	// 模板ID
	TemplateID string `json:"template_id"`
}
