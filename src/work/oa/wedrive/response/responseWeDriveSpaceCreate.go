package response

import (
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
)

type ResponseWeDriveSpaceCreate struct {
	response.ResponseWork

	SpaceID string `json:"spaceid"`
}
