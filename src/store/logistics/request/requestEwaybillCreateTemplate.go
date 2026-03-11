package request

// RequestEwaybillCreateTemplate 新增面单模板请求
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybillcreatetemplate.html
type RequestEwaybillCreateTemplate struct {
	// 模板配置信息
	ConfigInfo EwaybillConfigInfo `json:"config_info"`
}

// EwaybillConfigInfo 电子面单模板配置
type EwaybillConfigInfo struct {
	// 模板名称
	TemplateName string `json:"template_name"`
	// 模板ID（从标准模板获取）
	TemplateCode string `json:"template_code"`
	// 打印方向：1-横版，2-竖版
	Orientation int `json:"orientation"`
	// 缩放比例：0-不缩放，1-自动缩放
	ScaleType int `json:"scale_type"`
	// 是否打印商品信息
	PrintProductInfo bool `json:"print_product_info"`
	// 是否打印备注
	PrintRemark bool `json:"print_remark"`
	// 是否打印二维码
	PrintQRCode bool `json:"print_qr_code"`
}
