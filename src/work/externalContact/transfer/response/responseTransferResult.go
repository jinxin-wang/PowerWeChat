package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/power"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseTransferResult struct {
	response.ResponseWork

	Customer   []*power.HashMap `json:"customer"`
	NextCursor string           `json:"next_cursor"`
}
