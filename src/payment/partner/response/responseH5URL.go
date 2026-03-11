package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseH5URL struct {
	response.ResponsePayment

	H5URL string `json:"h5_url"`
}
