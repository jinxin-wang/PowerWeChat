package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

type ResponseGenAftersaleOrder struct {
	response.ResponseStore
	AftersaleID string `json:"aftersale_id"`
}
