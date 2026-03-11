package request

type SupplierInfo struct {
	SupplierID   string `json:"supplier_id"`    // 机构ID
	SupplierName string `json:"supplier_name"`  // 机构名称
}

type RequestBatchAddHeadSupplierItem struct {
	ProductIDs   []string    `json:"product_ids"`    // 商品ID列表
	SupplierInfo *SupplierInfo `json:"supplier_info"` // 机构信息
}
