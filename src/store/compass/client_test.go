package compass

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/jinxin-wang/PowerWeChat/v3/src/store/compass/request"
	"github.com/jinxin-wang/PowerWeChat/v3/src/store/compass/response"
)

// Test Request Structures

func TestRequestGetShopFinderAuthorizationList(t *testing.T) {
	req := &request.RequestGetShopFinderAuthorizationList{
		PageSize: 20,
		NextKey:  "test_next_key",
	}

	assert.Equal(t, 20, req.PageSize)
	assert.Equal(t, "test_next_key", req.NextKey)
}

func TestRequestGetShopFinderList(t *testing.T) {
	req := &request.RequestGetShopFinderList{
		PageSize:  30,
		NextKey:   "finder_next_key",
		StartTime: 1704067200,
		EndTime:   1706745600,
	}

	assert.Equal(t, 30, req.PageSize)
	assert.Equal(t, "finder_next_key", req.NextKey)
	assert.Equal(t, int64(1704067200), req.StartTime)
	assert.Equal(t, int64(1706745600), req.EndTime)
}

func TestRequestGetShopFinderOverall(t *testing.T) {
	req := &request.RequestGetShopFinderOverall{
		FinderID:  "finder_123",
		StartTime: 1704067200,
		EndTime:   1706745600,
	}

	assert.Equal(t, "finder_123", req.FinderID)
	assert.Equal(t, int64(1704067200), req.StartTime)
	assert.Equal(t, int64(1706745600), req.EndTime)
}

func TestRequestGetShopFinderProductList(t *testing.T) {
	req := &request.RequestGetShopFinderProductList{
		FinderID: "finder_456",
		PageSize: 50,
		NextKey:  "product_next_key",
	}

	assert.Equal(t, "finder_456", req.FinderID)
	assert.Equal(t, 50, req.PageSize)
	assert.Equal(t, "product_next_key", req.NextKey)
}

func TestRequestGetShopFinderProductOverall(t *testing.T) {
	req := &request.RequestGetShopFinderProductOverall{
		FinderID:  "finder_789",
		ProductID: "product_123",
		StartTime: 1704067200,
		EndTime:   1706745600,
	}

	assert.Equal(t, "finder_789", req.FinderID)
	assert.Equal(t, "product_123", req.ProductID)
	assert.Equal(t, int64(1704067200), req.StartTime)
	assert.Equal(t, int64(1706745600), req.EndTime)
}

func TestRequestGetShopLiveList(t *testing.T) {
	req := &request.RequestGetShopLiveList{
		PageSize:  25,
		NextKey:   "live_next_key",
		StartTime: 1704067200,
		EndTime:   1706745600,
	}

	assert.Equal(t, 25, req.PageSize)
	assert.Equal(t, "live_next_key", req.NextKey)
	assert.Equal(t, int64(1704067200), req.StartTime)
	assert.Equal(t, int64(1706745600), req.EndTime)
}

func TestRequestGetShopOverall(t *testing.T) {
	req := &request.RequestGetShopOverall{
		StartTime: 1704067200,
		EndTime:   1706745600,
	}

	assert.Equal(t, int64(1704067200), req.StartTime)
	assert.Equal(t, int64(1706745600), req.EndTime)
}

func TestRequestGetShopProductData(t *testing.T) {
	req := &request.RequestGetShopProductData{
		ProductID: "product_456",
		StartTime: 1704067200,
		EndTime:   1706745600,
	}

	assert.Equal(t, "product_456", req.ProductID)
	assert.Equal(t, int64(1704067200), req.StartTime)
	assert.Equal(t, int64(1706745600), req.EndTime)
}

func TestRequestGetShopProductList(t *testing.T) {
	req := &request.RequestGetShopProductList{
		PageSize:  40,
		NextKey:   "shop_product_next_key",
		StartTime: 1704067200,
		EndTime:   1706745600,
	}

	assert.Equal(t, 40, req.PageSize)
	assert.Equal(t, "shop_product_next_key", req.NextKey)
	assert.Equal(t, int64(1704067200), req.StartTime)
	assert.Equal(t, int64(1706745600), req.EndTime)
}

func TestRequestGetShopSaleProfileData(t *testing.T) {
	req := &request.RequestGetShopSaleProfileData{
		StartTime: 1704067200,
		EndTime:   1706745600,
		Type:      "age",
	}

	assert.Equal(t, int64(1704067200), req.StartTime)
	assert.Equal(t, int64(1706745600), req.EndTime)
	assert.Equal(t, "age", req.Type)
}

// Test Response Structures

func TestResponseGetShopFinderAuthorizationList(t *testing.T) {
	resp := &response.ResponseGetShopFinderAuthorizationList{
		FinderList: []response.ShopFinderAuthorizationListItem{
			{
				FinderID:      "finder_001",
				FinderName:    "视频号1",
				AuthorizeTime: 1704067200,
			},
			{
				FinderID:      "finder_002",
				FinderName:    "视频号2",
				AuthorizeTime: 1704153600,
			},
		},
		NextKey: "next_page_key",
		HasMore: true,
	}

	assert.Equal(t, 2, len(resp.FinderList))
	assert.Equal(t, "finder_001", resp.FinderList[0].FinderID)
	assert.Equal(t, "视频号1", resp.FinderList[0].FinderName)
	assert.Equal(t, "next_page_key", resp.NextKey)
	assert.Equal(t, true, resp.HasMore)
}

func TestResponseGetShopFinderList(t *testing.T) {
	resp := &response.ResponseGetShopFinderList{
		FinderList: []response.ShopFinderItem{
			{
				FinderID:      "finder_001",
				FinderName:    "带货达人1",
				FinderHeadImg: "https://example.com/head1.jpg",
			},
		},
		NextKey: "finder_next_key",
		HasMore: false,
	}

	assert.Equal(t, 1, len(resp.FinderList))
	assert.Equal(t, "finder_001", resp.FinderList[0].FinderID)
	assert.Equal(t, "带货达人1", resp.FinderList[0].FinderName)
	assert.Equal(t, false, resp.HasMore)
}

func TestResponseGetShopFinderOverall(t *testing.T) {
	resp := &response.ResponseGetShopFinderOverall{
		OverallData: response.FinderOverallData{
			PayGmv:           100000,
			PayOrderCount:    50,
			Uv:               1000,
			RefundAmount:     5000,
			RefundOrderCount: 2,
		},
	}

	assert.Equal(t, int64(100000), resp.OverallData.PayGmv)
	assert.Equal(t, int64(50), resp.OverallData.PayOrderCount)
	assert.Equal(t, int64(1000), resp.OverallData.Uv)
}

func TestResponseGetShopFinderProductList(t *testing.T) {
	resp := &response.ResponseGetShopFinderProductList{
		ProductList: []response.ShopFinderProductItem{
			{
				ProductID:   "product_001",
				ProductName: "商品1",
				SkuID:       "sku_001",
				PayAmount:   50000,
				PayCount:    10,
			},
		},
		NextKey: "product_list_next_key",
		HasMore: true,
	}

	assert.Equal(t, 1, len(resp.ProductList))
	assert.Equal(t, "product_001", resp.ProductList[0].ProductID)
	assert.Equal(t, int64(50000), resp.ProductList[0].PayAmount)
}

func TestResponseGetShopFinderProductOverall(t *testing.T) {
	resp := &response.ResponseGetShopFinderProductOverall{
		ProductOverall: response.FinderProductOverall{
			PayAmount:     100000,
			PayCount:      20,
			Uv:            500,
			AddCartCount:  100,
			ExposureCount: 2000,
		},
	}

	assert.Equal(t, int64(100000), resp.ProductOverall.PayAmount)
	assert.Equal(t, int64(20), resp.ProductOverall.PayCount)
	assert.Equal(t, int64(500), resp.ProductOverall.Uv)
}

func TestResponseGetShopLiveList(t *testing.T) {
	resp := &response.ResponseGetShopLiveList{
		LiveList: []response.ShopLiveItem{
			{
				LiveID:        "live_001",
				LiveTitle:     "直播标题1",
				StartTime:     1704067200,
				EndTime:       1704070800,
				PayGmv:        50000,
				PayOrderCount: 10,
			},
		},
		NextKey: "live_list_next_key",
		HasMore: false,
	}

	assert.Equal(t, 1, len(resp.LiveList))
	assert.Equal(t, "live_001", resp.LiveList[0].LiveID)
	assert.Equal(t, "直播标题1", resp.LiveList[0].LiveTitle)
}

func TestResponseGetShopOverall(t *testing.T) {
	resp := &response.ResponseGetShopOverall{
		OverallData: response.ShopOverallData{
			PayGmv:           500000,
			PayOrderCount:    200,
			PayUserCount:     150,
			Uv:               5000,
			RefundAmount:     10000,
			RefundOrderCount: 5,
		},
	}

	assert.Equal(t, int64(500000), resp.OverallData.PayGmv)
	assert.Equal(t, int64(200), resp.OverallData.PayOrderCount)
	assert.Equal(t, int64(150), resp.OverallData.PayUserCount)
}

func TestResponseGetShopProductData(t *testing.T) {
	resp := &response.ResponseGetShopProductData{
		ProductData: response.ShopProductData{
			ProductID:     "product_789",
			ProductName:   "测试商品",
			PayAmount:     100000,
			PayCount:      25,
			Uv:            800,
			ExposureCount: 3000,
			AddCartCount:  150,
		},
	}

	assert.Equal(t, "product_789", resp.ProductData.ProductID)
	assert.Equal(t, "测试商品", resp.ProductData.ProductName)
	assert.Equal(t, int64(100000), resp.ProductData.PayAmount)
}

func TestResponseGetShopProductList(t *testing.T) {
	resp := &response.ResponseGetShopProductList{
		ProductList: []response.ShopProductListItem{
			{
				ProductID:   "product_111",
				ProductName: "商品A",
				PayAmount:   30000,
				PayCount:    6,
				Uv:          200,
			},
			{
				ProductID:   "product_222",
				ProductName: "商品B",
				PayAmount:   45000,
				PayCount:    9,
				Uv:          300,
			},
		},
		NextKey: "shop_product_list_next_key",
		HasMore: true,
	}

	assert.Equal(t, 2, len(resp.ProductList))
	assert.Equal(t, "product_111", resp.ProductList[0].ProductID)
	assert.Equal(t, "商品B", resp.ProductList[1].ProductName)
}

func TestResponseGetShopSaleProfileData(t *testing.T) {
	resp := &response.ResponseGetShopSaleProfileData{
		ProfileData: []response.SaleProfileItem{
			{
				Key:   "18-24",
				Value: 150,
			},
			{
				Key:   "25-34",
				Value: 300,
			},
			{
				Key:   "35-44",
				Value: 200,
			},
		},
	}

	assert.Equal(t, 3, len(resp.ProfileData))
	assert.Equal(t, "18-24", resp.ProfileData[0].Key)
	assert.Equal(t, int64(150), resp.ProfileData[0].Value)
	assert.Equal(t, "25-34", resp.ProfileData[1].Key)
	assert.Equal(t, int64(300), resp.ProfileData[1].Value)
}
