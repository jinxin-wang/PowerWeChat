package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

type ResponseIsTradeManagementConfirmationCompleted struct {
	response.ResponseMiniProgram
	Completed bool `json:"completed"`
}
