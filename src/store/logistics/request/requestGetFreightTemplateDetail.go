package request

// RequestGetFreightTemplateDetail 查询运费模版详情请求
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/getfreighttemplatedetail.html
type RequestGetFreightTemplateDetail struct {
	// 模板ID
	TemplateID string `json:"template_id"`
}
