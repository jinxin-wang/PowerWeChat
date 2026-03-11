package response

type ItemResultInfo struct {
	ProductID    int    `json:"product_id"`     // 商品ID
	Ratio        int    `json:"ratio"`          // 佣金比例
	ServiceRatio int    `json:"service_ratio"` // 服务费比例
}

type ResponseBatchAddItem struct {
	response.ResponseStore
	ResultList []ItemResultInfo `json:"result_list"`
}

type ResponseDeleteItem struct {
	response.ResponseStore
}

type ItemInfo struct {
	ProductID    int           `json:"product_id"`      // 商品ID
	Title        string        `json:"title"`           // 商品标题
	ThumbImg     string        `json:"thumb_img"`       // 商品图片
	PriceInfo    PriceInfo     `json:"price_info"`      // 价格信息
	CommissionInfo CommissionInfo `json:"commission_info"` // 佣金信息
	CategoryID   int           `json:"category_id"`     // 类目ID
	BrandID      int           `json:"brand_id"`        // 品牌ID
	Status       int           `json:"status"`          // 商品状态
	CreateTime   int           `json:"create_time"`     // 创建时间
	UpdateTime   int           `json:"update_time"`     // 更新时间
}

type PriceInfo struct {
	Price     int    `json:"price"`      // 价格
	MarketPrice int   `json:"market_price"` // 市场价
	CostPrice  int    `json:"cost_price"`  // 成本价
}

type CommissionInfo struct {
	Ratio       int    `json:"ratio"`        // 佣金比例
	ServiceRatio int   `json:"service_ratio"` // 服务费比例
}

type ResponseGetItem struct {
	response.ResponseStore
	Item ItemInfo `json:"item"`
}

type ResponseBatchAddHeadSupplierItem struct {
	response.ResponseStore
	ResultInfoList []ItemResultInfo `json:"result_info_list"`
}

type ResponseGetItemList struct {
	response.ResponseStore
	ItemList []ItemInfo `json:"item_list"`
	NextKey  string    `json:"next_key"`
	HasMore  bool      `json:"has_more"`
}

type ResponseUpdateItem struct {
	response.ResponseStore
}
