package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseWeDriveFileUpload struct {
	response.ResponseWork

	FileID string `json:"fileid"`
}
