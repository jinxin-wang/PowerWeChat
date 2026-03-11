package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseAccountServiceAddContactWay struct {
	response.ResponseWork

	URL string `json:"url"`
}
