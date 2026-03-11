package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

type ResponseGetAPIQuota struct {
	response.ResponseStore
	Quota      string `json:"quota"`
	QuotaLimit string `json:"quota_limit"`
}
