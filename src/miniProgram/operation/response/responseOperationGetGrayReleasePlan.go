package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/power"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseOperationGetGrayReleasePlan struct {
	response.ResponseMiniProgram
	GrayReleasePlan *power.HashMap `json:"gray_release_plan"`
}
