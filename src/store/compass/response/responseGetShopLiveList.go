package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// ShopLiveItem 店铺开播列表项
type ShopLiveItem struct {
	LiveID        string `json:"live_id,omitempty"`       // 直播ID
	LiveTitle     string `json:"live_title,omitempty"`    // 直播标题
	StartTime     int64  `json:"start_time,omitempty"`    // 开播时间
	EndTime       int64  `json:"end_time,omitempty"`      // 结束时间
	PayGmv        int64  `json:"pay_gmv,omitempty"`       // 支付GMV(分)
	PayOrderCount int64  `json:"pay_order_cnt,omitempty"` // 支付订单数
}

type ResponseGetShopLiveList struct {
	response.ResponseStore
	LiveList []ShopLiveItem `json:"live_list,omitempty"`
	NextKey  string         `json:"next_key,omitempty"`
	HasMore  bool           `json:"has_more,omitempty"`
}
