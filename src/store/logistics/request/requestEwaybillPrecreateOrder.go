package request

// RequestEwaybillPrecreateOrder 电子面单预取号请求
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybillprecreateorder.html
type RequestEwaybillPrecreateOrder struct {
	// 快递公司ID
	DeliveryID string `json:"delivery_id"`
	// 运单信息
	WaybillInfo EwaybillWaybillInfo `json:"waybill_info"`
}

// EwaybillWaybillInfo 运单信息
type EwaybillWaybillInfo struct {
	// 发货地址ID
	SendAddressID string `json:"send_address_id"`
	// 收货地址信息
	ReceiveAddressInfo EwaybillAddressInfo `json:"receive_address_info"`
	// 包裹信息
	PackageInfo EwaybillPackageInfo `json:"package_info"`
}

// EwaybillAddressInfo 地址信息
type EwaybillAddressInfo struct {
	// 收件人姓名
	ReceiverName string `json:"receiver_name"`
	// 联系电话
	Tel string `json:"tel"`
	// 省份
	Province string `json:"province"`
	// 城市
	City string `json:"city"`
	// 区县
	District string `json:"district"`
	// 详细地址
	Detail string `json:"detail"`
}

// EwaybillPackageInfo 包裹信息
type EwaybillPackageInfo struct {
	// 包裹重量（克）
	Weight int `json:"weight"`
	// 包裹体积（立方厘米）
	Volume int `json:"volume,omitempty"`
	// 包裹商品列表
	ProductList []EwaybillProduct `json:"product_list"`
}

// EwaybillProduct 包裹商品
type EwaybillProduct struct {
	// 商品名称
	ProductName string `json:"product_name"`
	// 商品数量
	ProductNum int `json:"product_num"`
}
