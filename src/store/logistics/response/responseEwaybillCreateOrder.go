package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// ResponseEwaybillCreateOrder 电子面单取号响应
type ResponseEwaybillCreateOrder struct {
	response.ResponseStore
	// 电子面单ID
	WaybillID string `json:"waybill_id"`
	// 面单数据（用于打印）
	WaybillData string `json:"waybill_data"`
}
