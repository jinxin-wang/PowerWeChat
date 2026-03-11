package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseDepartmentCreate struct {
	response.ResponseWork
	ID int `json:"id"`
}
