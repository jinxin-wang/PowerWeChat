package order

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestRequestGetOrderList(t *testing.T) {
	req := &RequestGetOrderList{
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

func TestRequestGetOrder(t *testing.T) {
	req := &RequestGetOrder{
		OrderID: "order_123456",
	}

	assert.Equal(t, "order_123456", req.OrderID)
}

func TestRequestSearchOrder(t *testing.T) {
	req := &RequestSearchOrder{
		Keyword:  "测试",
		PageSize: 20,
		NextKey:  "next_key",
	}

	assert.Equal(t, "测试", req.Keyword)
	assert.Equal(t, 20, req.PageSize)
	assert.Equal(t, "next_key", req.NextKey)
}

func TestRequestUpdateOrderPrice(t *testing.T) {
	req := &RequestUpdateOrderPrice{
		OrderID:    "order_123",
		ChangeType: 1,
		Price:      999,
		Remark:     "修改价格",
	}

	assert.Equal(t, "order_123", req.OrderID)
	assert.Equal(t, 1, req.ChangeType)
	assert.Equal(t, 999, req.Price)
	assert.Equal(t, "修改价格", req.Remark)
}

func TestRequestUpdateOrderMerchantNote(t *testing.T) {
	req := &RequestUpdateOrderMerchantNote{
		OrderID: "order_456",
		Note:    "重要客户",
	}

	assert.Equal(t, "order_456", req.OrderID)
	assert.Equal(t, "重要客户", req.Note)
}

func TestRequestUpdateOrderAddress(t *testing.T) {
	req := &RequestUpdateOrderAddress{
		OrderID:      "order_789",
		ReceiverName: "张三",
		DetailInfo:   "某某小区1号楼",
		TelNumber:    "13800138000",
		PostalCode:   "100000",
		ProvinceName: "北京市",
		CityName:     "北京市",
		CountyName:   "朝阳区",
	}

	assert.Equal(t, "order_789", req.OrderID)
	assert.Equal(t, "张三", req.ReceiverName)
	assert.Equal(t, "13800138000", req.TelNumber)
}

func TestRequestUpdateOrderDelivery(t *testing.T) {
	req := &RequestUpdateOrderDelivery{
		OrderID:        "order_001",
		DeliveryID:     "SF",
		WaybillID:      "SF123456789",
		DeliveryName:   "顺丰速运",
		DeliveryRemark: "加急",
	}

	assert.Equal(t, "order_001", req.OrderID)
	assert.Equal(t, "SF", req.DeliveryID)
	assert.Equal(t, "SF123456789", req.WaybillID)
}

func TestRequestAcceptOrderAddressModify(t *testing.T) {
	req := &RequestAcceptOrderAddressModify{
		OrderID: "order_modify_1",
	}

	assert.Equal(t, "order_modify_1", req.OrderID)
}

func TestRequestRejectOrderAddressModify(t *testing.T) {
	req := &RequestRejectOrderAddressModify{
		OrderID: "order_modify_2",
		Reason:  "地址不详细",
	}

	assert.Equal(t, "order_modify_2", req.OrderID)
	assert.Equal(t, "地址不详细", req.Reason)
}

func TestRequestDecodeSensitiveInfo(t *testing.T) {
	req := &RequestDecodeSensitiveInfo{
		OrderID:       "order_secret",
		EncryptedData: "encrypted_data_here",
	}

	assert.Equal(t, "order_secret", req.OrderID)
	assert.Equal(t, "encrypted_data_here", req.EncryptedData)
}

func TestResponseGetOrderList(t *testing.T) {
	resp := &ResponseGetOrderList{
		OrderList: []OrderInfo{
			{OrderID: "order_1"},
			{OrderID: "order_2"},
		},
		NextKey: "next_page_key",
		HasMore: true,
	}

	assert.Equal(t, 2, len(resp.OrderList))
	assert.Equal(t, "order_1", resp.OrderList[0].OrderID)
	assert.Equal(t, "next_page_key", resp.NextKey)
	assert.Equal(t, true, resp.HasMore)
}

func TestResponseGetOrder(t *testing.T) {
	resp := &ResponseGetOrder{
		Order: OrderInfo{
			OrderID: "order_detail",
			Status:  1,
		},
	}

	assert.Equal(t, "order_detail", resp.Order.OrderID)
	assert.Equal(t, 1, resp.Order.Status)
}

func TestResponseSearchOrder(t *testing.T) {
	resp := &ResponseSearchOrder{
		OrderList: []OrderInfo{},
		HasMore:   false,
	}

	assert.Equal(t, 0, len(resp.OrderList))
	assert.Equal(t, false, resp.HasMore)
}

func TestOrderInfo(t *testing.T) {
	info := OrderInfo{
		OrderID:    "test_order",
		Status:     2,
		CreateTime: "2024-01-01 10:00:00",
		OpenID:     "openid_123",
	}

	assert.Equal(t, "test_order", info.OrderID)
	assert.Equal(t, 2, info.Status)
	assert.Equal(t, "openid_123", info.OpenID)
}

func TestProductInfo(t *testing.T) {
	info := ProductInfo{
		ProductID:  "prod_1",
		SkuID:      "sku_1",
		Title:      "测试商品",
		SalePrice:  100,
		ProductCnt: 2,
	}

	assert.Equal(t, "prod_1", info.ProductID)
	assert.Equal(t, "测试商品", info.Title)
	assert.Equal(t, 2, info.ProductCnt)
}

func TestAddressInfo(t *testing.T) {
	info := AddressInfo{
		UserName:     "李四",
		TelNumber:    "13900139000",
		ProvinceName: "广东省",
		CityName:     "深圳市",
		DetailInfo:   "科技园",
	}

	assert.Equal(t, "李四", info.UserName)
	assert.Equal(t, "深圳市", info.CityName)
}

func TestResponseDecodeSensitiveInfo(t *testing.T) {
	resp := &ResponseDecodeSensitiveInfo{
		ReceiverName: "王五",
		TelNumber:    "13700137000",
		DetailInfo:   "详细地址",
	}

	assert.Equal(t, "王五", resp.ReceiverName)
	assert.Equal(t, "13700137000", resp.TelNumber)
	assert.Equal(t, "详细地址", resp.DetailInfo)
}

// Batch 2 API Tests

func TestRequestUploadFreshInsurance(t *testing.T) {
	req := &RequestUploadFreshInsurance{
		OrderID:       "order_fresh_123",
		InsuranceInfo: "质检通过",
	}

	assert.Equal(t, "order_fresh_123", req.OrderID)
	assert.Equal(t, "质检通过", req.InsuranceInfo)
}

func TestRequestAddGiftOrderNote(t *testing.T) {
	req := &RequestAddGiftOrderNote{
		OrderID: "order_gift_123",
		Note:    "礼品订单备注",
	}

	assert.Equal(t, "order_gift_123", req.OrderID)
	assert.Equal(t, "礼品订单备注", req.Note)
}

func TestRequestGetGiftOrderSubList(t *testing.T) {
	req := &RequestGetGiftOrderSubList{
		OrderID:  "order_gift_456",
		PageSize: 20,
		NextKey:  "next_key",
	}

	assert.Equal(t, "order_gift_456", req.OrderID)
	assert.Equal(t, 20, req.PageSize)
	assert.Equal(t, "next_key", req.NextKey)
}

func TestResponseGetGiftOrderSubList(t *testing.T) {
	resp := &ResponseGetGiftOrderSubList{
		SubOrderList: []SubOrderInfo{
			{SubOrderID: "sub_1"},
			{SubOrderID: "sub_2"},
		},
		NextKey: "next_page",
		HasMore: true,
	}

	assert.Equal(t, 2, len(resp.SubOrderList))
	assert.Equal(t, "sub_1", resp.SubOrderList[0].SubOrderID)
	assert.Equal(t, true, resp.HasMore)
}

func TestRequestGetSKUChangeList(t *testing.T) {
	req := &RequestGetSKUChangeList{
		OrderID:  "order_sku_123",
		PageSize: 10,
		NextKey:  "key",
	}

	assert.Equal(t, "order_sku_123", req.OrderID)
	assert.Equal(t, 10, req.PageSize)
}

func TestResponseGetSKUChangeList(t *testing.T) {
	resp := &ResponseGetSKUChangeList{
		SKUChangeList: []SKUChangeInfo{
			{
				SKUChangeID: "change_1",
				OrderID:     "order_1",
				Status:      1,
			},
		},
		HasMore: false,
	}

	assert.Equal(t, 1, len(resp.SKUChangeList))
	assert.Equal(t, "change_1", resp.SKUChangeList[0].SKUChangeID)
}

func TestRequestAcceptSKUChange(t *testing.T) {
	req := &RequestAcceptSKUChange{
		OrderID:     "order_123",
		SKUChangeID: "change_123",
	}

	assert.Equal(t, "order_123", req.OrderID)
	assert.Equal(t, "change_123", req.SKUChangeID)
}

func TestRequestRejectSKUChange(t *testing.T) {
	req := &RequestRejectSKUChange{
		OrderID:     "order_456",
		SKUChangeID: "change_456",
		Reason:      "库存不足",
	}

	assert.Equal(t, "order_456", req.OrderID)
	assert.Equal(t, "库存不足", req.Reason)
}

func TestRequestApplyRealNumber(t *testing.T) {
	req := &RequestApplyRealNumber{
		OrderID: "order_real_123",
		Reason:  "售后需要",
	}

	assert.Equal(t, "order_real_123", req.OrderID)
	assert.Equal(t, "售后需要", req.Reason)
}

func TestResponseApplyRealNumber(t *testing.T) {
	resp := &ResponseApplyRealNumber{
		ApplyID: "apply_123",
	}

	assert.Equal(t, "apply_123", resp.ApplyID)
}

func TestRequestGetRealNumberStatus(t *testing.T) {
	req := &RequestGetRealNumberStatus{
		ApplyID: "apply_456",
	}

	assert.Equal(t, "apply_456", req.ApplyID)
}

func TestResponseGetRealNumberStatus(t *testing.T) {
	resp := &ResponseGetRealNumberStatus{
		Status:      1,
		PhoneNumber: "13800138000",
		AuditResult: "通过",
	}

	assert.Equal(t, 1, resp.Status)
	assert.Equal(t, "13800138000", resp.PhoneNumber)
}

func TestRequestReapplyVirtualNumber(t *testing.T) {
	req := &RequestReapplyVirtualNumber{
		OrderID: "order_virtual_123",
	}

	assert.Equal(t, "order_virtual_123", req.OrderID)
}

func TestRequestDelayVirtualNumber(t *testing.T) {
	req := &RequestDelayVirtualNumber{
		OrderID:   "order_virtual_456",
		DelayDays: 7,
	}

	assert.Equal(t, "order_virtual_456", req.OrderID)
	assert.Equal(t, 7, req.DelayDays)
}

func TestRequestAddPhoneVerifyCode(t *testing.T) {
	req := &RequestAddPhoneVerifyCode{
		OrderID: "order_phone_123",
		Phone:   "13900139000",
	}

	assert.Equal(t, "order_phone_123", req.OrderID)
	assert.Equal(t, "13900139000", req.Phone)
}

func TestRequestSendPhoneVerifyCode(t *testing.T) {
	req := &RequestSendPhoneVerifyCode{
		OrderID: "order_phone_456",
	}

	assert.Equal(t, "order_phone_456", req.OrderID)
}

func TestRequestGetPhoneStatus(t *testing.T) {
	req := &RequestGetPhoneStatus{
		OrderID: "order_phone_status_123",
	}

	assert.Equal(t, "order_phone_status_123", req.OrderID)
}

func TestResponseGetPhoneStatus(t *testing.T) {
	resp := &ResponseGetPhoneStatus{
		Status: 1,
		Phone:  "13700137000",
	}

	assert.Equal(t, 1, resp.Status)
	assert.Equal(t, "13700137000", resp.Phone)
}
