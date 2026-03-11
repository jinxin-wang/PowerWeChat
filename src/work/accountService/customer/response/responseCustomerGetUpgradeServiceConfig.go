package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/power"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseCustomerGetUpgradeServiceConfig struct {
	response.ResponseWork

	MemberRange    *power.HashMap `json:"member_range"`
	GroupChatRange *power.HashMap `json:"groupchat_range"`
}
