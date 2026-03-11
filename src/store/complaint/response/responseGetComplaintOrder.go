package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

type ResponseGetComplaintOrder struct {
	response.ResponseStore
	ComplaintInfo *ComplaintInfo `json:"complaint_info"`
}

type ComplaintInfo struct {
	ComplaintID      string `json:"complaint_id"`
	OrderID          string `json:"order_id"`
	OpenID           string `json:"openid"`
	ComplaintType    int    `json:"complaint_type"`
	ComplaintDetail  string `json:"complaint_detail"`
	ComplaintTime    string `json:"complaint_time"`
	Status           int    `json:"status"`
}
