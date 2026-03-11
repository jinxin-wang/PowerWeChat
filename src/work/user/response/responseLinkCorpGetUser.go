package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/power"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseLinkCorpGetUser struct {
	response.ResponseWork

	UserInfo *power.HashMap `json:"user_info"`
}
