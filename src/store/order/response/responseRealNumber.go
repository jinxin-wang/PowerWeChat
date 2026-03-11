package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseApplyRealNumber struct {
	response.ResponseStore
	ApplyID string `json:"apply_id"`
}

type ResponseGetRealNumberStatus struct {
	response.ResponseStore
	Status      int    `json:"status"`
	PhoneNumber string `json:"phone_number"`
	AuditResult string `json:"audit_result,omitempty"`
}

type ResponseGetPhoneStatus struct {
	response.ResponseStore
	Status int    `json:"status"`
	Phone  string `json:"phone,omitempty"`
}
