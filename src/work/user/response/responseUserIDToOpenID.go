package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseUserIDToOpenID struct {
	response.ResponseWork

	OpenID string `json:"openid"`
}
