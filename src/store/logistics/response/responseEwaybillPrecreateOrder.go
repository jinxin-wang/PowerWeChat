package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

// ResponseEwaybillPrecreateOrder 电子面单预取号响应
type ResponseEwaybillPrecreateOrder struct {
	response.ResponseStore
	// 电子面单ID
	WaybillID string `json:"waybill_id"`
}
