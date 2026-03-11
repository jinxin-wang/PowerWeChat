package request

// RequestGetUserScoreFlowRecord 获取用户积分流水请求
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/vip/api_getuserscoreflowrecord.html
type RequestGetUserScoreFlowRecord struct {
	Openid   string `json:"openid"`    // 用户openid
	PageSize int    `json:"page_size"` // 分页大小，最大100
	NextKey  string `json:"next_key"`  // 分页游标，第一页为空
}
