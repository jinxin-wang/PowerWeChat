package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// ScoreFlowItem 积分流水项结构
type ScoreFlowItem struct {
	OrderId    string `json:"order_id"`    // 订单ID
	Score      int    `json:"score"`       // 积分变动值
	Type       int    `json:"type"`        // 变动类型: 1-获取 2-消耗
	CreateTime int64  `json:"create_time"` // 变动时间戳
	Desc       string `json:"desc"`        // 变动描述
}

// ResponseGetUserScoreFlowRecord 获取用户积分流水响应
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/vip/api_getuserscoreflowrecord.html
type ResponseGetUserScoreFlowRecord struct {
	response.ResponseStore
	FlowList []ScoreFlowItem `json:"flow_list"` // 积分流水列表
	NextKey  string          `json:"next_key"`  // 分页游标
	HasMore  bool            `json:"has_more"`  // 是否还有更多数据
}
