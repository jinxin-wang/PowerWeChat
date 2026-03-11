package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseLivingCreate struct {
	response.ResponseWork

	LivingID int `json:"livingid"`
}
