package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/power"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseDataCubeGetPerformanceData struct {
	Data *power.HashMap `json:"data"`

	response.ResponseMiniProgram
}
