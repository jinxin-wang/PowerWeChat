package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/power"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseCorpCheckInRules struct {
	response.ResponseWork
	Group []*power.HashMap `json:"group"`
}
