package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// ResponseEwaybillGetOrder 查询面单详情响应
type ResponseEwaybillGetOrder struct {
	response.ResponseStore
	// 电子面单信息
	WaybillInfo WaybillInfoResponse `json:"waybill_info"`
}

// WaybillInfoResponse 电子面单信息
type WaybillInfoResponse struct {
	// 电子面单ID
	WaybillID string `json:"waybill_id"`
	// 快递公司ID
	DeliveryID string `json:"delivery_id"`
	// 快递公司名称
	DeliveryName string `json:"delivery_name"`
	// 订单ID
	OrderID string `json:"order_id"`
	// 电子面单状态：1-待使用，2-已使用，3-已取消
	Status int `json:"status"`
	// 取号时间
	CreateTime string `json:"create_time"`
}
