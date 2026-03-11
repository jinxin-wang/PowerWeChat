package logistics

import (
	"context"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel"
	"github.com/jinxin-wang/PowerWeChat/v3/src/store/logistics/request"
	"github.com/jinxin-wang/PowerWeChat/v3/src/store/logistics/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// ===== 地址管理 =====

// AddAddress 添加地址
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/addlogisticsaddress.html
func (comp *Client) AddAddress(ctx context.Context, data *request.RequestAddAddress) (*response.ResponseAddAddress, error) {
	result := &response.ResponseAddAddress{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIAddAddress, params, nil, nil, result)

	return result, err
}

// GetAddressList 获取地址列表
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/getlogisticsaddresslist.html
func (comp *Client) GetAddressList(ctx context.Context, data *request.RequestGetAddressList) (*response.ResponseGetAddressList, error) {
	result := &response.ResponseGetAddressList{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpGet(ctx, APIGetAddressList, params, nil, result)

	return result, err
}

// GetAddress 获取地址详情
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/getlogisticsaddress.html
func (comp *Client) GetAddress(ctx context.Context, data *request.RequestGetAddress) (*response.ResponseGetAddress, error) {
	result := &response.ResponseGetAddress{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetAddress, params, nil, nil, result)

	return result, err
}

// UpdateAddress 更新地址
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/updatelogisticsaddress.html
func (comp *Client) UpdateAddress(ctx context.Context, data *request.RequestUpdateAddress) (*response.ResponseUpdateAddress, error) {
	result := &response.ResponseUpdateAddress{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIUpdateAddress, params, nil, nil, result)

	return result, err
}

// DeleteAddress 删除地址
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/deletelogisticsaddress.html
func (comp *Client) DeleteAddress(ctx context.Context, data *request.RequestDeleteAddress) (*response.ResponseDeleteAddress, error) {
	result := &response.ResponseDeleteAddress{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIDeleteAddress, params, nil, nil, result)

	return result, err
}

// ===== 运费模板 =====

// AddFreightTemplate 增加运费模版
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/addfreighttemplate.html
func (comp *Client) AddFreightTemplate(ctx context.Context, data *request.RequestAddFreightTemplate) (*response.ResponseAddFreightTemplate, error) {
	result := &response.ResponseAddFreightTemplate{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIAddFreightTemplate, params, nil, nil, result)

	return result, err
}

// GetFreightTemplateDetail 查询运费模版
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/getfreighttemplatedetail.html
func (comp *Client) GetFreightTemplateDetail(ctx context.Context, data *request.RequestGetFreightTemplateDetail) (*response.ResponseGetFreightTemplateDetail, error) {
	result := &response.ResponseGetFreightTemplateDetail{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetFreightTemplateDetail, params, nil, nil, result)

	return result, err
}

// GetFreightTemplateList 获取运费模板列表
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/getfreighttemplatelist.html
func (comp *Client) GetFreightTemplateList(ctx context.Context, data *request.RequestGetFreightTemplateList) (*response.ResponseGetFreightTemplateList, error) {
	result := &response.ResponseGetFreightTemplateList{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetFreightTemplateList, params, nil, nil, result)

	return result, err
}

// UpdateFreightTemplate 更新运费模版
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/updatefreighttemplate.html
func (comp *Client) UpdateFreightTemplate(ctx context.Context, data *request.RequestUpdateFreightTemplate) (*response.ResponseUpdateFreightTemplate, error) {
	result := &response.ResponseUpdateFreightTemplate{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIUpdateFreightTemplate, params, nil, nil, result)

	return result, err
}

// ===== 电子面单 =====

// EwaybillGetTemplateConfig 获取面单标准模板
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybillgettemplateconfig.html
func (comp *Client) EwaybillGetTemplateConfig(ctx context.Context) (*response.ResponseEwaybillGetTemplateConfig, error) {
	result := &response.ResponseEwaybillGetTemplateConfig{}

	_, err := comp.BaseClient.HttpGet(ctx, APIEwaybillGetTemplateConfig, nil, nil, result)

	return result, err
}

// EwaybillCreateTemplate 新增面单模板
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybillcreatetemplate.html
func (comp *Client) EwaybillCreateTemplate(ctx context.Context, data *request.RequestEwaybillCreateTemplate) (*response.ResponseEwaybillCreateTemplate, error) {
	result := &response.ResponseEwaybillCreateTemplate{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIEwaybillCreateTemplate, params, nil, nil, result)

	return result, err
}

// EwaybillDelTemplate 删除面单模版
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybilldeltemplate.html
func (comp *Client) EwaybillDelTemplate(ctx context.Context, data *request.RequestEwaybillDelTemplate) (*response.ResponseEwaybillDelTemplate, error) {
	result := &response.ResponseEwaybillDelTemplate{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIEwaybillDelTemplate, params, nil, nil, result)

	return result, err
}

// EwaybillUpdateTemplate 更新面单模版
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybillupdatetemplate.html
func (comp *Client) EwaybillUpdateTemplate(ctx context.Context, data *request.RequestEwaybillUpdateTemplate) (*response.ResponseEwaybillUpdateTemplate, error) {
	result := &response.ResponseEwaybillUpdateTemplate{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIEwaybillUpdateTemplate, params, nil, nil, result)

	return result, err
}

// EwaybillGetTemplate 获取面单模板信息
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybillgettemplate.html
func (comp *Client) EwaybillGetTemplate(ctx context.Context, data *request.RequestEwaybillGetTemplate) (*response.ResponseEwaybillGetTemplate, error) {
	result := &response.ResponseEwaybillGetTemplate{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIEwaybillGetTemplate, params, nil, nil, result)

	return result, err
}

// EwaybillGetTemplateByID 根据模板ID获取面单模板信息
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybillgettemplatebyid.html
func (comp *Client) EwaybillGetTemplateByID(ctx context.Context, data *request.RequestEwaybillGetTemplateByID) (*response.ResponseEwaybillGetTemplateByID, error) {
	result := &response.ResponseEwaybillGetTemplateByID{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIEwaybillGetTemplateByID, params, nil, nil, result)

	return result, err
}

// EwaybillGetAcct 查询开通的电子面单网点/账号信息
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybillgetacct.html
func (comp *Client) EwaybillGetAcct(ctx context.Context) (*response.ResponseEwaybillGetAcct, error) {
	result := &response.ResponseEwaybillGetAcct{}

	_, err := comp.BaseClient.HttpGet(ctx, APIEwaybillGetAcct, nil, nil, result)

	return result, err
}

// EwaybillGetDeliveryList 查询开通的快递公司列表
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybillgetdeliverylist.html
func (comp *Client) EwaybillGetDeliveryList(ctx context.Context) (*response.ResponseEwaybillGetDeliveryList, error) {
	result := &response.ResponseEwaybillGetDeliveryList{}

	_, err := comp.BaseClient.HttpGet(ctx, APIEwaybillGetDeliveryList, nil, nil, result)

	return result, err
}

// EwaybillPrecreateOrder 电子面单预取号
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybillprecreateorder.html
func (comp *Client) EwaybillPrecreateOrder(ctx context.Context, data *request.RequestEwaybillPrecreateOrder) (*response.ResponseEwaybillPrecreateOrder, error) {
	result := &response.ResponseEwaybillPrecreateOrder{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIEwaybillPrecreateOrder, params, nil, nil, result)

	return result, err
}

// EwaybillCreateOrder 电子面单取号
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybillcreateorder.html
func (comp *Client) EwaybillCreateOrder(ctx context.Context, data *request.RequestEwaybillCreateOrder) (*response.ResponseEwaybillCreateOrder, error) {
	result := &response.ResponseEwaybillCreateOrder{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIEwaybillCreateOrder, params, nil, nil, result)

	return result, err
}

// EwaybillAddSubOrder 电子面单子件追加
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybilladdsuborder.html
func (comp *Client) EwaybillAddSubOrder(ctx context.Context, data *request.RequestEwaybillAddSubOrder) (*response.ResponseEwaybillAddSubOrder, error) {
	result := &response.ResponseEwaybillAddSubOrder{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIEwaybillAddSubOrder, params, nil, nil, result)

	return result, err
}

// EwaybillCancelOrder 电子面单取消下单
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybillcancelorder.html
func (comp *Client) EwaybillCancelOrder(ctx context.Context, data *request.RequestEwaybillCancelOrder) (*response.ResponseEwaybillCancelOrder, error) {
	result := &response.ResponseEwaybillCancelOrder{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIEwaybillCancelOrder, params, nil, nil, result)

	return result, err
}

// EwaybillGetOrder 查询面单详情
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybillgetorder.html
func (comp *Client) EwaybillGetOrder(ctx context.Context, data *request.RequestEwaybillGetOrder) (*response.ResponseEwaybillGetOrder, error) {
	result := &response.ResponseEwaybillGetOrder{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIEwaybillGetOrder, params, nil, nil, result)

	return result, err
}

// EwaybillGetPrintContent 获取打印报文
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybillgetprintcontent.html
func (comp *Client) EwaybillGetPrintContent(ctx context.Context, data *request.RequestEwaybillGetPrintContent) (*response.ResponseEwaybillGetPrintContent, error) {
	result := &response.ResponseEwaybillGetPrintContent{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIEwaybillGetPrintContent, params, nil, nil, result)

	return result, err
}

// EwaybillPrintOrder 打印成功通知
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybillprintorder.html
func (comp *Client) EwaybillPrintOrder(ctx context.Context, data *request.RequestEwaybillPrintOrder) (*response.ResponseEwaybillPrintOrder, error) {
	result := &response.ResponseEwaybillPrintOrder{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIEwaybillPrintOrder, params, nil, nil, result)

	return result, err
}

// EwaybillBatchPrintOrder 批量打印通知
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/ewaybillbatchprintorder.html
func (comp *Client) EwaybillBatchPrintOrder(ctx context.Context, data *request.RequestEwaybillBatchPrintOrder) (*response.ResponseEwaybillBatchPrintOrder, error) {
	result := &response.ResponseEwaybillBatchPrintOrder{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIEwaybillBatchPrintOrder, params, nil, nil, result)

	return result, err
}

// ===== 发货 =====

// SendDelivery 订单发货
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/senddelivery.html
func (comp *Client) SendDelivery(ctx context.Context, data *request.RequestSendDelivery) (*response.ResponseSendDelivery, error) {
	result := &response.ResponseSendDelivery{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APISendDelivery, params, nil, nil, result)

	return result, err
}

// DeliveryCompensation 订单补发货
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/deliverycompensation.html
func (comp *Client) DeliveryCompensation(ctx context.Context, data *request.RequestDeliveryCompensation) (*response.ResponseDeliveryCompensation, error) {
	result := &response.ResponseDeliveryCompensation{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIDeliveryCompensation, params, nil, nil, result)

	return result, err
}

// GetDeliveryCompanyListNew 获取快递公司列表
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/getdeliverycompanylist.html
func (comp *Client) GetDeliveryCompanyListNew(ctx context.Context) (*response.ResponseGetDeliveryCompanyListNew, error) {
	result := &response.ResponseGetDeliveryCompanyListNew{}

	_, err := comp.BaseClient.HttpGet(ctx, APIGetDeliveryCompanyListNew, nil, nil, result)

	return result, err
}

// GetDeliveryCompanyList 获取快递公司列表-旧
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/getdeliverycompanylistold.html
func (comp *Client) GetDeliveryCompanyList(ctx context.Context) (*response.ResponseGetDeliveryCompanyList, error) {
	result := &response.ResponseGetDeliveryCompanyList{}

	_, err := comp.BaseClient.HttpGet(ctx, APIGetDeliveryCompanyList, nil, nil, result)

	return result, err
}

// ===== 虚拟号码 =====

// GetPrivateNumberPool 获取虚拟号码池
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/getprivatenumberpool.html
func (comp *Client) GetPrivateNumberPool(ctx context.Context) (*response.ResponseGetPrivateNumberPool, error) {
	result := &response.ResponseGetPrivateNumberPool{}

	_, err := comp.BaseClient.HttpGet(ctx, APIGetPrivateNumberPool, nil, nil, result)

	return result, err
}

// GetRealNumber 根据运单号获取真实手机号
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/getrealnumber.html
func (comp *Client) GetRealNumber(ctx context.Context, data *request.RequestGetRealNumber) (*response.ResponseGetRealNumber, error) {
	result := &response.ResponseGetRealNumber{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetRealNumber, params, nil, nil, result)

	return result, err
}

// GetVirtualNumber 根据运单号获取虚拟手机号
// https://developers.weixin.qq.com/doc/store/shop/API/logistics/getvirtualnumber.html
func (comp *Client) GetVirtualNumber(ctx context.Context, data *request.RequestGetVirtualNumber) (*response.ResponseGetVirtualNumber, error) {
	result := &response.ResponseGetVirtualNumber{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetVirtualNumber, params, nil, nil, result)

	return result, err
}
