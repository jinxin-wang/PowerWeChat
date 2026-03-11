package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseUnitfy struct {
	response.ResponsePayment

	PrepayID string `json:"prepay_id"`
}
