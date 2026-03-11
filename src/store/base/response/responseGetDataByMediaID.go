package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

type ResponseGetDataByMediaID struct {
	response.ResponseStore
	Data string `json:"data"`
}
