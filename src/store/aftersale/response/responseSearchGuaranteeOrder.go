package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

type GuaranteeInfo struct {
	GuaranteeID string `json:"guarantee_id"`
	OrderID     string `json:"order_id"`
	Status      int    `json:"status"`
	Type        int    `json:"type"`
	Amount      int    `json:"amount"`
	Reason      string `json:"reason,omitempty"`
	ProofInfo   string `json:"proof_info,omitempty"`
	ExtraInfo   string `json:"extra_info,omitempty"`
	CreateTime  string `json:"create_time"`
	UpdateTime  string `json:"update_time"`
}

type ResponseSearchGuaranteeOrder struct {
	response.ResponseStore
	GuaranteeList []GuaranteeInfo `json:"guarantee_list"`
	NextKey       string          `json:"next_key"`
	HasMore       bool            `json:"has_more"`
}
