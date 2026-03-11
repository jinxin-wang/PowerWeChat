package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseCodePayResult struct {
	response.ResponsePayment

	ResponseOrder
}
