package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseLivingGetLivingCode struct {
	response.ResponseWork

	LivingCode int `json:"living_code"`
}
