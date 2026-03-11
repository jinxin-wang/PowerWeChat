package request

// RequestEwaybillGetPrintContent 获取打印报文请求
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybillgetprintcontent.html
type RequestEwaybillGetPrintContent struct {
	// 电子面单ID列表
	WaybillIDs []string `json:"waybill_ids"`
	// 模板ID
	TemplateID string `json:"template_id,omitempty"`
}
