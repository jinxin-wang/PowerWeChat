package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/power"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseSubscribeMessageGetCategory struct {
	response.ResponseMiniProgram
	Data []*power.HashMap `json:"data"`
}
