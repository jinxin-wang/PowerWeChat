package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseCallbackCheck struct {
	response.ResponseStore
	Operator string `json:"operator"`
	Result   string `json:"result"`
}
