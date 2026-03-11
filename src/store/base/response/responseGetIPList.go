package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

type ResponseGetAPIDomainIP struct {
	response.ResponseStore
	IPList []string `json:"ip_list"`
}

type ResponseGetCallbackIP struct {
	response.ResponseStore
	IPList []string `json:"ip_list"`
}
