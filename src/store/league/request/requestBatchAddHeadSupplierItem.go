package request

// SupplierInfo 团长供应商信息
type SupplierInfo struct {
	SupplierID   string `json:"supplier_id,omitempty"`   // 机构ID
	SupplierName string `json:"supplier_name,omitempty"` // 机构名称
}

// RequestBatchAddHeadSupplierItem 批量新增联盟机构推广请求
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/item/api_batchaddheadsupplieritem.html
type RequestBatchAddHeadSupplierItem struct {
	ProductIDs   []string      `json:"product_ids,omitempty"`   // 商品ID列表
	SupplierInfo *SupplierInfo `json:"supplier_info,omitempty"` // 机构信息
}
