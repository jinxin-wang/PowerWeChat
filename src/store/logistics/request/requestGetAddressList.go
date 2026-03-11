package request

// RequestGetAddressList 获取地址列表请求
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/getlogisticsaddresslist.html
type RequestGetAddressList struct {
	// 分页大小，默认10，最大100
	PageSize int64 `json:"page_size,omitempty"`
	// 分页游标
	NextKey string `json:"next_key,omitempty"`
}
