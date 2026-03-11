package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/power"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseExpressBatchOrderList struct {
	response.ResponseMiniProgram

	OrderList []*power.HashMap `json:"order_list"`
}
