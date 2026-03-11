package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseGetAftersaleRejectReason struct {
	response.ResponseStore
	ReasonList []ReasonInfo `json:"reason_list"`
}
