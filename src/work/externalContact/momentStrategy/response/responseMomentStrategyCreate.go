package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/power"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseMomentStrategyCreate struct {
	response.ResponseWork

	Strategy   []*power.HashMap `json:"strategy"`
	NextCursor string           `json:"next_cursor"`
}
