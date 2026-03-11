package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ReasonInfo struct {
	ReasonID   int    `json:"reason_id"`
	ReasonText string `json:"reason_text"`
}

type ResponseGetAftersaleReason struct {
	response.ResponseStore
	ReasonList []ReasonInfo `json:"reason_list"`
}
