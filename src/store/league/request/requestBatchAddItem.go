package request

type RequestBatchAddItem struct {
	ProductIDs []string `json:"product_ids"` // 商品ID列表
}
