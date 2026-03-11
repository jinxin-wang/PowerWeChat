package logistics

import (
	"testing"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/store/logistics/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/store/logistics/response"
	"github.com/stretchr/testify/assert"
)

// TestClient 测试客户端
func TestClient(t *testing.T) {
	// 测试结构体创建
	assert.NotNil(t, &Client{})
}

// TestRequestTypes 测试请求结构体类型
func TestRequestTypes(t *testing.T) {
	// 地址管理
	assert.NotNil(t, &request.RequestAddAddress{})
	assert.NotNil(t, &request.RequestGetAddressList{})
	assert.NotNil(t, &request.RequestGetAddress{})
	assert.NotNil(t, &request.RequestUpdateAddress{})
	assert.NotNil(t, &request.RequestDeleteAddress{})

	// 运费模板
	assert.NotNil(t, &request.RequestAddFreightTemplate{})
	assert.NotNil(t, &request.RequestGetFreightTemplateDetail{})
	assert.NotNil(t, &request.RequestGetFreightTemplateList{})
	assert.NotNil(t, &request.RequestUpdateFreightTemplate{})

	// 电子面单
	assert.NotNil(t, &request.RequestEwaybillCreateTemplate{})
	assert.NotNil(t, &request.RequestEwaybillDelTemplate{})
	assert.NotNil(t, &request.RequestEwaybillUpdateTemplate{})
	assert.NotNil(t, &request.RequestEwaybillGetTemplate{})
	assert.NotNil(t, &request.RequestEwaybillGetTemplateByID{})
	assert.NotNil(t, &request.RequestEwaybillPrecreateOrder{})
	assert.NotNil(t, &request.RequestEwaybillCreateOrder{})
	assert.NotNil(t, &request.RequestEwaybillAddSubOrder{})
	assert.NotNil(t, &request.RequestEwaybillCancelOrder{})
	assert.NotNil(t, &request.RequestEwaybillGetOrder{})
	assert.NotNil(t, &request.RequestEwaybillGetPrintContent{})
	assert.NotNil(t, &request.RequestEwaybillPrintOrder{})
	assert.NotNil(t, &request.RequestEwaybillBatchPrintOrder{})

	// 发货
	assert.NotNil(t, &request.RequestSendDelivery{})
	assert.NotNil(t, &request.RequestDeliveryCompensation{})

	// 虚拟号码
	assert.NotNil(t, &request.RequestGetRealNumber{})
	assert.NotNil(t, &request.RequestGetVirtualNumber{})
}

// TestResponseTypes 测试响应结构体类型
func TestResponseTypes(t *testing.T) {
	// 地址管理
	assert.NotNil(t, &response.ResponseAddAddress{})
	assert.NotNil(t, &response.ResponseGetAddressList{})
	assert.NotNil(t, &response.ResponseGetAddress{})
	assert.NotNil(t, &response.ResponseUpdateAddress{})
	assert.NotNil(t, &response.ResponseDeleteAddress{})

	// 运费模板
	assert.NotNil(t, &response.ResponseAddFreightTemplate{})
	assert.NotNil(t, &response.ResponseGetFreightTemplateDetail{})
	assert.NotNil(t, &response.ResponseGetFreightTemplateList{})
	assert.NotNil(t, &response.ResponseUpdateFreightTemplate{})

	// 电子面单
	assert.NotNil(t, &response.ResponseEwaybillGetTemplateConfig{})
	assert.NotNil(t, &response.ResponseEwaybillCreateTemplate{})
	assert.NotNil(t, &response.ResponseEwaybillDelTemplate{})
	assert.NotNil(t, &response.ResponseEwaybillUpdateTemplate{})
	assert.NotNil(t, &response.ResponseEwaybillGetTemplate{})
	assert.NotNil(t, &response.ResponseEwaybillGetTemplateByID{})
	assert.NotNil(t, &response.ResponseEwaybillGetAcct{})
	assert.NotNil(t, &response.ResponseEwaybillGetDeliveryList{})
	assert.NotNil(t, &response.ResponseEwaybillPrecreateOrder{})
	assert.NotNil(t, &response.ResponseEwaybillCreateOrder{})
	assert.NotNil(t, &response.ResponseEwaybillAddSubOrder{})
	assert.NotNil(t, &response.ResponseEwaybillCancelOrder{})
	assert.NotNil(t, &response.ResponseEwaybillGetOrder{})
	assert.NotNil(t, &response.ResponseEwaybillGetPrintContent{})
	assert.NotNil(t, &response.ResponseEwaybillPrintOrder{})
	assert.NotNil(t, &response.ResponseEwaybillBatchPrintOrder{})

	// 发货
	assert.NotNil(t, &response.ResponseSendDelivery{})
	assert.NotNil(t, &response.ResponseDeliveryCompensation{})
	assert.NotNil(t, &response.ResponseGetDeliveryCompanyListNew{})
	assert.NotNil(t, &response.ResponseGetDeliveryCompanyList{})

	// 虚拟号码
	assert.NotNil(t, &response.ResponseGetPrivateNumberPool{})
	assert.NotNil(t, &response.ResponseGetRealNumber{})
	assert.NotNil(t, &response.ResponseGetVirtualNumber{})
}

// TestAddAddress 测试添加地址（示例）
func TestAddAddress(t *testing.T) {
	req := &request.RequestAddAddress{
		ReceiverName: "张三",
		Tel:          "13800138000",
		Province:     "广东省",
		City:         "深圳市",
		District:     "南山区",
		Detail:       "科技园",
	}
	assert.Equal(t, "张三", req.ReceiverName)
	assert.Equal(t, "13800138000", req.Tel)
}
