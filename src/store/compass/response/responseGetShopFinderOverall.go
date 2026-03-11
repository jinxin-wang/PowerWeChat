package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// FinderOverallData 带货数据概览
type FinderOverallData struct {
	PayGmv           int64 `json:"pay_gmv,omitempty"`          // 支付GMV(分)
	PayOrderCount    int64 `json:"pay_order_cnt,omitempty"`    // 支付订单数
	Uv               int64 `json:"uv,omitempty"`               // 访客数
	RefundAmount     int64 `json:"refund_amount,omitempty"`    // 退款金额(分)
	RefundOrderCount int64 `json:"refund_order_cnt,omitempty"` // 退款订单数
}

type ResponseGetShopFinderOverall struct {
	response.ResponseStore
	OverallData FinderOverallData `json:"overall_data,omitempty"`
}
