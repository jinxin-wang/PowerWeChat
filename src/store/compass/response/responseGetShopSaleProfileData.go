package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// SaleProfileItem 销售画像数据项
type SaleProfileItem struct {
	Key   string `json:"key,omitempty"`   // 维度值，如年龄段、性别、省份
	Value int64  `json:"value,omitempty"` // 对应维度的数值
}

type ResponseGetShopSaleProfileData struct {
	response.ResponseStore
	ProfileData []SaleProfileItem `json:"profile_data,omitempty"`
}
