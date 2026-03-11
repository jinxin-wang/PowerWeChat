package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseCustomerServiceMessageGetTempMedia struct {
	response.ResponseMiniProgram
	ContentType string `json:"contentType"`
	Buffer      []byte `json:"buffer"`
}
