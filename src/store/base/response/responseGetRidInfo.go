package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseGetRidInfo struct {
	response.ResponseStore
	RequestInfo string `json:"request_info"`
	RequestMsg  string `json:"request_msg"`
}
