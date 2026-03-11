package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseAuthGetPaidUnionID struct {
	UnionID string `json:"unionid"`

	response.ResponseMiniProgram
}
