package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseGetCallBackIP struct {
	response.ResponseOfficialAccount

	IPList []string `json:"ip_list"`
}
