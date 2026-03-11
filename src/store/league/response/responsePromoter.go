package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

type PromoterInfo struct {
	FinderID          string `json:"finder_id"`                    // 视频号finder_id
	PromoterID        string `json:"promoter_id,omitempty"`        // 达人带货id
	Status            int    `json:"status"`                       // 合作状态
	InviteTime        int    `json:"invite_time"`                 // 达人邀请秒级时间戳
	SaleProductNumber int    `json:"sale_product_number"`          // 带货商品数
	SaleGMV           int    `json:"sale_gmv"`                    // 带货GMV
}

type ResponseAddPromoter struct {
	response.ResponseStore
}

type ResponseDeletePromoter struct {
	response.ResponseStore
}

type ResponseGetPromoter struct {
	response.ResponseStore
	Promoter PromoterInfo `json:"promoter"`
}

type ResponseGetPromoterList struct {
	response.ResponseStore
	PromoterList []PromoterInfo `json:"promoter_list"`
	NextKey      string         `json:"next_key"`
	HasMore      bool           `json:"has_more"`
}

type ResponseUpdatePromoter struct {
	response.ResponseStore
}
