package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

type ResponseAddGroupWelcomeTemplate struct {
	response.ResponseWork

	TemplateID string `json:"template_id"`
}
