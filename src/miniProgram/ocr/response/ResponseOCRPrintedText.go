package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/power"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseOCRPrintedText struct {
	response.ResponseMiniProgram
	Items   []*power.HashMap `json:"items"`
	ImgSize *power.HashMap   `json:"img_size"`
}
