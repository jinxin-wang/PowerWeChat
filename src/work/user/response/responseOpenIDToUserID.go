package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseOpenIDToUserID struct {
	response.ResponseWork

	UserID string `json:"userid"`
}
