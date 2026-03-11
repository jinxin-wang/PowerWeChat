package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// ResponseEwaybillGetDeliveryList 查询开通的快递公司列表响应
type ResponseEwaybillGetDeliveryList struct {
	response.ResponseStore
	// 快递公司列表
	DeliveryList []EwaybillDelivery `json:"delivery_list"`
}

// EwaybillDelivery 快递公司信息
type EwaybillDelivery struct {
	// 快递公司ID
	DeliveryID string `json:"delivery_id"`
	// 快递公司名称
	DeliveryName string `json:"delivery_name"`
}
