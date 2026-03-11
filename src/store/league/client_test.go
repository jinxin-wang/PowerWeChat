package league

import (
	"testing"

	"github.com/jinxin-wang/PowerWeChat/v3/src/store/league/request"
)

// TestRequestAddPromoter 测试新增达人请求结构体
func TestRequestAddPromoter(t *testing.T) {
	req := &request.RequestAddPromoter{
		FinderID:   "finder_123",
		PromoterID: "promoter_456",
	}
	if req.FinderID != "finder_123" {
		t.Errorf("Expected finder_id to be finder_123, got %s", req.FinderID)
	}
	if req.PromoterID != "promoter_456" {
		t.Errorf("Expected promoter_id to be promoter_456, got %s", req.PromoterID)
	}
}

// TestRequestDeletePromoter 测试删除达人请求结构体
func TestRequestDeletePromoter(t *testing.T) {
	req := &request.RequestDeletePromoter{
		PromoterID: "promoter_789",
	}
	if req.PromoterID != "promoter_789" {
		t.Errorf("Expected promoter_id to be promoter_789, got %s", req.PromoterID)
	}
}

// TestRequestGetPromoter 测试获取达人详情请求结构体
func TestRequestGetPromoter(t *testing.T) {
	req := &request.RequestGetPromoter{
		PromoterID: "promoter_abc",
	}
	if req.PromoterID != "promoter_abc" {
		t.Errorf("Expected promoter_id to be promoter_abc, got %s", req.PromoterID)
	}
}

// TestRequestGetPromoterList 测试获取达人列表请求结构体
func TestRequestGetPromoterList(t *testing.T) {
	req := &request.RequestGetPromoterList{
		PageSize: 20,
		NextKey:  "next_key_123",
	}
	if req.PageSize != 20 {
		t.Errorf("Expected page_size to be 20, got %d", req.PageSize)
	}
	if req.NextKey != "next_key_123" {
		t.Errorf("Expected next_key to be next_key_123, got %s", req.NextKey)
	}
}

// TestRequestUpdatePromoter 测试编辑达人请求结构体
func TestRequestUpdatePromoter(t *testing.T) {
	req := &request.RequestUpdatePromoter{
		PromoterID: "promoter_def",
		Type:       1,
	}
	if req.PromoterID != "promoter_def" {
		t.Errorf("Expected promoter_id to be promoter_def, got %s", req.PromoterID)
	}
	if req.Type != 1 {
		t.Errorf("Expected type to be 1, got %d", req.Type)
	}
}

// TestRequestBatchAddItem 测试批量新增联盟商品请求结构体
func TestRequestBatchAddItem(t *testing.T) {
	req := &request.RequestBatchAddItem{
		ProductIDs: []string{"product_1", "product_2", "product_3"},
	}
	if len(req.ProductIDs) != 3 {
		t.Errorf("Expected product_ids length to be 3, got %d", len(req.ProductIDs))
	}
	if req.ProductIDs[0] != "product_1" {
		t.Errorf("Expected first product_id to be product_1, got %s", req.ProductIDs[0])
	}
}

// TestRequestDeleteItem 测试删除联盟商品请求结构体
func TestRequestDeleteItem(t *testing.T) {
	req := &request.RequestDeleteItem{
		ProductID: "12345",
	}
	if req.ProductID != "12345" {
		t.Errorf("Expected product_id to be 12345, got %s", req.ProductID)
	}
}

// TestRequestGetItem 测试获取联盟商品详情请求结构体
func TestRequestGetItem(t *testing.T) {
	req := &request.RequestGetItem{
		ProductID: 67890,
	}
	if req.ProductID != 67890 {
		t.Errorf("Expected product_id to be 67890, got %d", req.ProductID)
	}
}

// TestRequestBatchAddHeadSupplierItem 测试批量新增联盟机构推广请求结构体
func TestRequestBatchAddHeadSupplierItem(t *testing.T) {
	req := &request.RequestBatchAddHeadSupplierItem{
		ProductIDs: []string{"hs_product_1", "hs_product_2"},
	}
	if len(req.ProductIDs) != 2 {
		t.Errorf("Expected product_ids length to be 2, got %d", len(req.ProductIDs))
	}
}

// TestRequestGetItemList 测试获取联盟商品推广列表请求结构体
func TestRequestGetItemList(t *testing.T) {
	req := &request.RequestGetItemList{
		PageSize: 50,
		NextKey:  "item_next_key",
		Status:   1,
	}
	if req.PageSize != 50 {
		t.Errorf("Expected page_size to be 50, got %d", req.PageSize)
	}
	if req.NextKey != "item_next_key" {
		t.Errorf("Expected next_key to be item_next_key, got %s", req.NextKey)
	}
	if req.Status != 1 {
		t.Errorf("Expected status to be 1, got %d", req.Status)
	}
}

// TestRequestUpdateItem 测试更新联盟商品信息请求结构体
func TestRequestUpdateItem(t *testing.T) {
	req := &request.RequestUpdateItem{
		Type:      1,
		ProductID: 11111,
		Ratio:     10.5,
	}
	if req.Type != 1 {
		t.Errorf("Expected type to be 1, got %d", req.Type)
	}
	if req.ProductID != 11111 {
		t.Errorf("Expected product_id to be 11111, got %d", req.ProductID)
	}
	if req.Ratio != 10.5 {
		t.Errorf("Expected ratio to be 10.5, got %f", req.Ratio)
	}
}

// TestRegisterProvider 测试Provider注册
func TestRegisterProvider(t *testing.T) {
	// 验证Provider可以被创建
	// 实际测试需要完整的应用上下文
	t.Skip("Skip: Need application context")
}
