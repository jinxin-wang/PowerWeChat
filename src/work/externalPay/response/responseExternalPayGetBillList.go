package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/power"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseExternalPayGetBillList struct {
	response.ResponseWork

	NextCursor string           `json:"next_cursor"`
	BillList   []*power.HashMap `json:"bill_list"`
}
