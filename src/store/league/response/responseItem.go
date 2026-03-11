package response

import "github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"

// ItemResultInfo 商品结果信息
type ItemResultInfo struct {
	ProductID    int `json:"product_id,omitempty"`    // 商品ID
	Ratio        int `json:"ratio,omitempty"`         // 佣金比例
	ServiceRatio int `json:"service_ratio,omitempty"` // 服务费比例
}

// ResponseBatchAddItem 批量新增联盟商品响应
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/item/api_batchadditem.html
type ResponseBatchAddItem struct {
	response.ResponseStore
	ResultList []ItemResultInfo `json:"result_list,omitempty"`
}

// ResponseDeleteItem 删除联盟商品响应
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/item/api_deleteitem.html
type ResponseDeleteItem struct {
	response.ResponseStore
}

// ItemInfo 商品信息
type ItemInfo struct {
	ProductID      int            `json:"product_id,omitempty"`      // 商品ID
	Title          string         `json:"title,omitempty"`           // 商品标题
	ThumbImg       string         `json:"thumb_img,omitempty"`       // 商品图片
	PriceInfo      PriceInfo      `json:"price_info,omitempty"`      // 价格信息
	CommissionInfo CommissionInfo `json:"commission_info,omitempty"` // 佣金信息
	CategoryID     int            `json:"category_id,omitempty"`     // 类目ID
	BrandID        int            `json:"brand_id,omitempty"`        // 品牌ID
	Status         int            `json:"status,omitempty"`          // 商品状态
	CreateTime     int            `json:"create_time,omitempty"`     // 创建时间
	UpdateTime     int            `json:"update_time,omitempty"`     // 更新时间
}

// PriceInfo 价格信息
type PriceInfo struct {
	Price       int `json:"price,omitempty"`        // 价格
	MarketPrice int `json:"market_price,omitempty"` // 市场价
	CostPrice   int `json:"cost_price,omitempty"`   // 成本价
}

// CommissionInfo 佣金信息
type CommissionInfo struct {
	Ratio        int `json:"ratio,omitempty"`         // 佣金比例
	ServiceRatio int `json:"service_ratio,omitempty"` // 服务费比例
}

// ResponseGetItem 获取联盟商品详情响应
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/item/api_getitem.html
type ResponseGetItem struct {
	response.ResponseStore
	Item ItemInfo `json:"item,omitempty"`
}

// ResponseBatchAddHeadSupplierItem 批量新增联盟机构推广响应
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/item/api_batchaddheadsupplieritem.html
type ResponseBatchAddHeadSupplierItem struct {
	response.ResponseStore
	ResultInfoList []ItemResultInfo `json:"result_info_list,omitempty"`
}

// ResponseGetItemList 获取联盟商品推广列表响应
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/item/api_getitemlist.html
type ResponseGetItemList struct {
	response.ResponseStore
	ItemList []ItemInfo `json:"item_list,omitempty"`
	NextKey  string     `json:"next_key,omitempty"`
	HasMore  bool       `json:"has_more,omitempty"`
}

// ResponseUpdateItem 更新联盟商品信息响应
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/item/api_upditem.html
type ResponseUpdateItem struct {
	response.ResponseStore
}
