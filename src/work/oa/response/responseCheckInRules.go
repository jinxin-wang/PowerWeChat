package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/power"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseCheckInRules struct {
	response.ResponseWork
	Info []*power.HashMap `json:"info"`
}
