package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

type ResponseGetGuaranteeOrder struct {
	response.ResponseStore
	GuaranteeInfo GuaranteeInfo `json:"guarantee_info"`
}
