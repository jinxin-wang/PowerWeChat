package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseAccountServiceSendMsgOnEvent struct {
	response.ResponseWork

	MsgID string `json:"msgid"`
}
