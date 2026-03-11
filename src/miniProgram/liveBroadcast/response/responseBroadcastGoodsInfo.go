package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/power"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseBroadcastGoodsInfo struct {
	response.ResponseMiniProgram

	Total int              `json:"total"`
	Goods []*power.HashMap `json:"goods"`
}
