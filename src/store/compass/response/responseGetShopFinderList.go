package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// ShopFinderItem 带货达人列表项
type ShopFinderItem struct {
	FinderID      string `json:"finder_id,omitempty"`       // 视频号ID
	FinderName    string `json:"finder_name,omitempty"`     // 视频号名称
	FinderHeadImg string `json:"finder_head_img,omitempty"` // 视频号头像
}

type ResponseGetShopFinderList struct {
	response.ResponseStore
	FinderList []ShopFinderItem `json:"finder_list,omitempty"`
	NextKey    string           `json:"next_key,omitempty"`
	HasMore    bool             `json:"has_more,omitempty"`
}
