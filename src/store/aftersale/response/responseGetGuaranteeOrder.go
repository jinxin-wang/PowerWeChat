package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseGetGuaranteeOrder struct {
	response.ResponseStore
	GuaranteeInfo GuaranteeInfo `json:"guarantee_info"`
}
