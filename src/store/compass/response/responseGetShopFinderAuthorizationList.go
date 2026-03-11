package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// ShopFinderAuthorizationListItem 授权视频号列表项
type ShopFinderAuthorizationListItem struct {
	FinderID      string `json:"finder_id,omitempty"`      // 视频号ID
	FinderName    string `json:"finder_name,omitempty"`    // 视频号名称
	AuthorizeTime int64  `json:"authorize_time,omitempty"` // 授权时间
}

type ResponseGetShopFinderAuthorizationList struct {
	response.ResponseStore
	FinderList []ShopFinderAuthorizationListItem `json:"finder_list,omitempty"`
	NextKey    string                            `json:"next_key,omitempty"`
	HasMore    bool                              `json:"has_more,omitempty"`
}
