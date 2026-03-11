package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/power"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseMsgAuditGetAgreeInfo struct {
	response.ResponseWork
	AgreeInfo []*power.HashMap `json:"agreeinfo"`
}
