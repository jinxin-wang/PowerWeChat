package aftersale

import (
	"testing"

	"github.com/jinxin-wang/PowerWeChat/v3/src/store/aftersale/request"
	"github.com/jinxin-wang/PowerWeChat/v3/src/store/aftersale/response"
	"github.com/go-playground/assert/v2"
)

func TestRequestGetAftersaleList(t *testing.T) {
	req := &request.RequestGetAftersaleList{
		PageSize:  10,
		NextKey:   "key123",
		Status:    1,
		StartTime: "2024-01-01",
		EndTime:   "2024-12-31",
	}

	assert.Equal(t, 10, req.PageSize)
	assert.Equal(t, "key123", req.NextKey)
	assert.Equal(t, 1, req.Status)
	assert.Equal(t, "2024-01-01", req.StartTime)
	assert.Equal(t, "2024-12-31", req.EndTime)
}

func TestRequestGetAftersaleOrder(t *testing.T) {
	req := &request.RequestGetAftersaleOrder{
		AftersaleID: "aftersale_123456",
	}

	assert.Equal(t, "aftersale_123456", req.AftersaleID)
}

func TestRequestAcceptApply(t *testing.T) {
	req := &request.RequestAcceptApply{
		AftersaleID: "aftersale_123",
		AddressID:   "address_456",
	}

	assert.Equal(t, "aftersale_123", req.AftersaleID)
	assert.Equal(t, "address_456", req.AddressID)
}

func TestRequestAcceptExchangeReship(t *testing.T) {
	req := &request.RequestAcceptExchangeReship{
		AftersaleID: "aftersale_789",
		DeliveryID:  "SF",
		WaybillID:   "SF123456789",
	}

	assert.Equal(t, "aftersale_789", req.AftersaleID)
	assert.Equal(t, "SF", req.DeliveryID)
	assert.Equal(t, "SF123456789", req.WaybillID)
}

func TestRequestGenAftersaleOrder(t *testing.T) {
	req := &request.RequestGenAftersaleOrder{
		OrderID:   "order_123",
		ProductID: "prod_456",
		Type:      1,
		Reason:    "商品质量问题",
	}

	assert.Equal(t, "order_123", req.OrderID)
	assert.Equal(t, "prod_456", req.ProductID)
	assert.Equal(t, 1, req.Type)
	assert.Equal(t, "商品质量问题", req.Reason)
}

func TestRequestSearchGuaranteeOrder(t *testing.T) {
	req := &request.RequestSearchGuaranteeOrder{
		PageSize: 20,
		NextKey:  "next_key",
		Status:   2,
	}

	assert.Equal(t, 20, req.PageSize)
	assert.Equal(t, "next_key", req.NextKey)
	assert.Equal(t, 2, req.Status)
}

func TestRequestGetGuaranteeOrder(t *testing.T) {
	req := &request.RequestGetGuaranteeOrder{
		GuaranteeID: "guarantee_123",
	}

	assert.Equal(t, "guarantee_123", req.GuaranteeID)
}

func TestRequestMerchantAcceptGuarantee(t *testing.T) {
	req := &request.RequestMerchantAcceptGuarantee{
		GuaranteeID: "guarantee_456",
	}

	assert.Equal(t, "guarantee_456", req.GuaranteeID)
}

func TestRequestMerchantModifyGuarantee(t *testing.T) {
	req := &request.RequestMerchantModifyGuarantee{
		GuaranteeID: "guarantee_789",
		Amount:      500,
	}

	assert.Equal(t, "guarantee_789", req.GuaranteeID)
	assert.Equal(t, 500, req.Amount)
}

func TestRequestMerchantProofGuarantee(t *testing.T) {
	req := &request.RequestMerchantProofGuarantee{
		GuaranteeID: "guarantee_001",
		ProofInfo: &request.ProofInfo{
			Text:     "协商内容",
			ImageIDs: []string{"img1", "img2"},
		},
	}

	assert.Equal(t, "guarantee_001", req.GuaranteeID)
	assert.Equal(t, "协商内容", req.ProofInfo.Text)
	assert.Equal(t, 2, len(req.ProofInfo.ImageIDs))
}

func TestRequestMerchantRefuseGuarantee(t *testing.T) {
	req := &request.RequestMerchantRefuseGuarantee{
		GuaranteeID: "guarantee_002",
		Reason:      "不符合保障条件",
	}

	assert.Equal(t, "guarantee_002", req.GuaranteeID)
	assert.Equal(t, "不符合保障条件", req.Reason)
}

func TestRequestMerchantUpdateAftersale(t *testing.T) {
	req := &request.RequestMerchantUpdateAftersale{
		AftersaleID:  "aftersale_update_123",
		RefundAmount: 1000,
	}

	assert.Equal(t, "aftersale_update_123", req.AftersaleID)
	assert.Equal(t, 1000, req.RefundAmount)
}

func TestRequestGetAftersaleReason(t *testing.T) {
	req := &request.RequestGetAftersaleReason{}

	assert.Equal(t, &request.RequestGetAftersaleReason{}, req)
}

func TestRequestRejectApply(t *testing.T) {
	req := &request.RequestRejectApply{
		AftersaleID: "aftersale_reject_123",
		Reason:      "不符合退货条件",
	}

	assert.Equal(t, "aftersale_reject_123", req.AftersaleID)
	assert.Equal(t, "不符合退货条件", req.Reason)
}

func TestRequestRejectExchangeReship(t *testing.T) {
	req := &request.RequestRejectExchangeReship{
		AftersaleID: "aftersale_exchange_reject_123",
		Reason:      "无法发货",
	}

	assert.Equal(t, "aftersale_exchange_reject_123", req.AftersaleID)
	assert.Equal(t, "无法发货", req.Reason)
}

func TestRequestGetAftersaleRejectReason(t *testing.T) {
	req := &request.RequestGetAftersaleRejectReason{}

	assert.Equal(t, &request.RequestGetAftersaleRejectReason{}, req)
}

func TestRequestUploadRefundCertificate(t *testing.T) {
	req := &request.RequestUploadRefundCertificate{
		AftersaleID: "aftersale_refund_123",
		CertificateInfo: &request.CertificateInfo{
			ImageIDs: []string{"cert_img1", "cert_img2"},
		},
	}

	assert.Equal(t, "aftersale_refund_123", req.AftersaleID)
	assert.Equal(t, 2, len(req.CertificateInfo.ImageIDs))
}

func TestRequestRefundPriceDiff(t *testing.T) {
	req := &request.RequestRefundPriceDiff{
		OrderID:         "order_price_diff_123",
		PriceDiffAmount: 500,
	}

	assert.Equal(t, "order_price_diff_123", req.OrderID)
	assert.Equal(t, 500, req.PriceDiffAmount)
}

func TestRequestApplyVirtualTelnum(t *testing.T) {
	req := &request.RequestApplyVirtualTelnum{
		AftersaleID: "aftersale_virtual_123",
	}

	assert.Equal(t, "aftersale_virtual_123", req.AftersaleID)
}

func TestAftersaleInfo(t *testing.T) {
	info := response.AftersaleInfo{
		AftersaleID:  "aftersale_test",
		OrderID:      "order_test",
		Status:       1,
		Type:         2,
		Reason:       "七天无理由退货",
		RefundAmount: 200,
		CreateTime:   "2024-01-01 10:00:00",
		OpenID:       "openid_123",
	}

	assert.Equal(t, "aftersale_test", info.AftersaleID)
	assert.Equal(t, "order_test", info.OrderID)
	assert.Equal(t, 1, info.Status)
	assert.Equal(t, "七天无理由退货", info.Reason)
}

func TestProductInfo(t *testing.T) {
	info := response.ProductInfo{
		ProductID: "prod_test",
		SkuID:     "sku_test",
		Title:     "测试商品",
		SalePrice: 9900,
		Quantity:  2,
	}

	assert.Equal(t, "prod_test", info.ProductID)
	assert.Equal(t, "测试商品", info.Title)
	assert.Equal(t, 2, info.Quantity)
}

func TestGuaranteeInfo(t *testing.T) {
	info := response.GuaranteeInfo{
		GuaranteeID: "guarantee_test",
		OrderID:     "order_test",
		Status:      1,
		Type:        2,
		Amount:      500,
		Reason:      "商品破损",
		CreateTime:  "2024-01-01 10:00:00",
	}

	assert.Equal(t, "guarantee_test", info.GuaranteeID)
	assert.Equal(t, 1, info.Status)
	assert.Equal(t, 500, info.Amount)
}

func TestReasonInfo(t *testing.T) {
	info := response.ReasonInfo{
		ReasonID:   1,
		ReasonText: "商品质量问题",
	}

	assert.Equal(t, 1, info.ReasonID)
	assert.Equal(t, "商品质量问题", info.ReasonText)
}

func TestResponseGetAftersaleList(t *testing.T) {
	resp := &response.ResponseGetAftersaleList{
		AftersaleList: []response.AftersaleInfo{
			{AftersaleID: "aftersale_1"},
			{AftersaleID: "aftersale_2"},
		},
		NextKey: "next_page_key",
		HasMore: true,
	}

	assert.Equal(t, 2, len(resp.AftersaleList))
	assert.Equal(t, "aftersale_1", resp.AftersaleList[0].AftersaleID)
	assert.Equal(t, "next_page_key", resp.NextKey)
	assert.Equal(t, true, resp.HasMore)
}

func TestResponseGetAftersaleOrder(t *testing.T) {
	resp := &response.ResponseGetAftersaleOrder{
		AftersaleInfo: response.AftersaleInfo{
			AftersaleID: "aftersale_detail",
			Status:      1,
		},
	}

	assert.Equal(t, "aftersale_detail", resp.AftersaleInfo.AftersaleID)
	assert.Equal(t, 1, resp.AftersaleInfo.Status)
}

func TestResponseGenAftersaleOrder(t *testing.T) {
	resp := &response.ResponseGenAftersaleOrder{
		AftersaleID: "new_aftersale_123",
	}

	assert.Equal(t, "new_aftersale_123", resp.AftersaleID)
}

func TestResponseSearchGuaranteeOrder(t *testing.T) {
	resp := &response.ResponseSearchGuaranteeOrder{
		GuaranteeList: []response.GuaranteeInfo{},
		HasMore:       false,
	}

	assert.Equal(t, 0, len(resp.GuaranteeList))
	assert.Equal(t, false, resp.HasMore)
}

func TestResponseGetGuaranteeOrder(t *testing.T) {
	resp := &response.ResponseGetGuaranteeOrder{
		GuaranteeInfo: response.GuaranteeInfo{
			GuaranteeID: "guarantee_detail",
			Status:      2,
		},
	}

	assert.Equal(t, "guarantee_detail", resp.GuaranteeInfo.GuaranteeID)
	assert.Equal(t, 2, resp.GuaranteeInfo.Status)
}

func TestResponseGetAftersaleReason(t *testing.T) {
	resp := &response.ResponseGetAftersaleReason{
		ReasonList: []response.ReasonInfo{
			{ReasonID: 1, ReasonText: "质量问题"},
			{ReasonID: 2, ReasonText: "物流问题"},
		},
	}

	assert.Equal(t, 2, len(resp.ReasonList))
	assert.Equal(t, "质量问题", resp.ReasonList[0].ReasonText)
}

func TestResponseGetAftersaleRejectReason(t *testing.T) {
	resp := &response.ResponseGetAftersaleRejectReason{
		ReasonList: []response.ReasonInfo{
			{ReasonID: 1, ReasonText: "已超时"},
		},
	}

	assert.Equal(t, 1, len(resp.ReasonList))
	assert.Equal(t, "已超时", resp.ReasonList[0].ReasonText)
}

func TestResponseRefundPriceDiff(t *testing.T) {
	resp := &response.ResponseRefundPriceDiff{
		AftersaleID: "aftersale_refund_diff_123",
	}

	assert.Equal(t, "aftersale_refund_diff_123", resp.AftersaleID)
}
