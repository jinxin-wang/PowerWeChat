package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// ResponseEwaybillGetAcct 查询开通的电子面单网点/账号信息响应
type ResponseEwaybillGetAcct struct {
	response.ResponseStore
	// 账号列表
	AcctList []EwaybillAcct `json:"acct_list"`
}

// EwaybillAcct 电子面单账号
type EwaybillAcct struct {
	// 快递公司ID
	DeliveryID string `json:"delivery_id"`
	// 快递公司名称
	DeliveryName string `json:"delivery_name"`
	// 账号类型：1-月结账号，2-散户账号
	AcctType int `json:"acct_type"`
	// 电子面单账号
	AcctID string `json:"acct_id"`
}
