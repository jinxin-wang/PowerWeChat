package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// ResponseEwaybillGetPrintContent 获取打印报文响应
type ResponseEwaybillGetPrintContent struct {
	response.ResponseStore
	// 打印内容列表
	PrintContentList []PrintContent `json:"print_content_list"`
}

// PrintContent 打印内容
type PrintContent struct {
	// 电子面单ID
	WaybillID string `json:"waybill_id"`
	// 打印报文
	PrintData string `json:"print_data"`
}
