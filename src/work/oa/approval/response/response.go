package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

type ResponseCreateTemplate struct {
	response.ResponseWork

	TemplateId string `json:"template_id"`
}
