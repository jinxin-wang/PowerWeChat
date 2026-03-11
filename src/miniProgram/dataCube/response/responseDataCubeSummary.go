package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/power"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseDataCubeSummary struct {
	response.ResponseMiniProgram

	List []*power.HashMap `json:"list"`
}
