package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseOCRBankcard struct {
	response.ResponseMiniProgram
	Number string `json:"number"`
}
