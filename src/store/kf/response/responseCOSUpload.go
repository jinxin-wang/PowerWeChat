package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

type ResponseCOSUpload struct {
	response.ResponseStore
	MediaID string `json:"media_id"`
	URL     string `json:"url"`
}
