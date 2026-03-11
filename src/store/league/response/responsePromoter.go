package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// PromoterInfo 达人信息
type PromoterInfo struct {
	FinderID          string `json:"finder_id,omitempty"`           // 视频号finder_id
	PromoterID        string `json:"promoter_id,omitempty"`         // 达人带货id
	Status            int    `json:"status,omitempty"`              // 合作状态
	InviteTime        int    `json:"invite_time,omitempty"`         // 达人邀请秒级时间戳
	SaleProductNumber int    `json:"sale_product_number,omitempty"` // 带货商品数
	SaleGMV           int    `json:"sale_gmv,omitempty"`            // 带货GMV
}

// ResponseAddPromoter 新增达人响应
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/promoter/api_addpromoter.html
type ResponseAddPromoter struct {
	response.ResponseStore
}

// ResponseDeletePromoter 删除达人响应
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/promoter/api_deletepromoter.html
type ResponseDeletePromoter struct {
	response.ResponseStore
}

// ResponseGetPromoter 获取达人详情信息响应
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/promoter/api_getpromoter.html
type ResponseGetPromoter struct {
	response.ResponseStore
	Promoter PromoterInfo `json:"promoter,omitempty"`
}

// ResponseGetPromoterList 获取商店达人列表响应
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/promoter/api_getpromoterlist.html
type ResponseGetPromoterList struct {
	response.ResponseStore
	PromoterList []PromoterInfo `json:"promoter_list,omitempty"`
	NextKey      string         `json:"next_key,omitempty"`
	HasMore      bool           `json:"has_more,omitempty"`
}

// ResponseUpdatePromoter 编辑达人响应
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/promoter/api_updpromoter.html
type ResponseUpdatePromoter struct {
	response.ResponseStore
}
