package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

type ResponseUrlScheme struct {
	response.ResponseMiniProgram

	OpenLink string `json:"openlink"`
}
