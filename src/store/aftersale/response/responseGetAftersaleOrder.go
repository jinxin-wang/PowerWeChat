package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseGetAftersaleOrder struct {
	response.ResponseStore
	AftersaleInfo AftersaleInfo `json:"aftersale_info"`
}
