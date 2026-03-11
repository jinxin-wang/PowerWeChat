package aftersale

import (
	"context"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel"
	"github.com/jinxin-wang/PowerWeChat/v3/src/store/aftersale/request"
	"github.com/jinxin-wang/PowerWeChat/v3/src/store/aftersale/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// GetAftersaleList 获取售后单列表
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_getaftersalelist.html
func (comp *Client) GetAftersaleList(ctx context.Context, data *request.RequestGetAftersaleList) (*response.ResponseGetAftersaleList, error) {
	result := &response.ResponseGetAftersaleList{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetAftersaleList, params, nil, nil, result)

	return result, err
}

// GetAftersaleOrder 获取售后单详情
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_getaftersaleorder.html
func (comp *Client) GetAftersaleOrder(ctx context.Context, data *request.RequestGetAftersaleOrder) (*response.ResponseGetAftersaleOrder, error) {
	result := &response.ResponseGetAftersaleOrder{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetAftersaleOrder, params, nil, nil, result)

	return result, err
}

// AcceptApply 同意售后
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_acceptapply.html
func (comp *Client) AcceptApply(ctx context.Context, data *request.RequestAcceptApply) (*response.ResponseAcceptApply, error) {
	result := &response.ResponseAcceptApply{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIAcceptApply, params, nil, nil, result)

	return result, err
}

// AcceptExchangeReship 换货发货
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_acceptexchangereship.html
func (comp *Client) AcceptExchangeReship(ctx context.Context, data *request.RequestAcceptExchangeReship) (*response.ResponseAcceptExchangeReship, error) {
	result := &response.ResponseAcceptExchangeReship{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIAcceptExchangeReship, params, nil, nil, result)

	return result, err
}

// GenAftersaleOrder 代用户发起售后
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_genaftersaleorder.html
func (comp *Client) GenAftersaleOrder(ctx context.Context, data *request.RequestGenAftersaleOrder) (*response.ResponseGenAftersaleOrder, error) {
	result := &response.ResponseGenAftersaleOrder{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGenAftersaleOrder, params, nil, nil, result)

	return result, err
}

// SearchGuaranteeOrder 商家获取保障单列表
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_searchguaranteeorder.html
func (comp *Client) SearchGuaranteeOrder(ctx context.Context, data *request.RequestSearchGuaranteeOrder) (*response.ResponseSearchGuaranteeOrder, error) {
	result := &response.ResponseSearchGuaranteeOrder{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APISearchGuaranteeOrder, params, nil, nil, result)

	return result, err
}

// GetGuaranteeOrder 获取保障单详情
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_getguaranteeorder.html
func (comp *Client) GetGuaranteeOrder(ctx context.Context, data *request.RequestGetGuaranteeOrder) (*response.ResponseGetGuaranteeOrder, error) {
	result := &response.ResponseGetGuaranteeOrder{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetGuaranteeOrder, params, nil, nil, result)

	return result, err
}

// MerchantAcceptGuarantee 商家同意保障单申请
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_merchantacceptguarantee.html
func (comp *Client) MerchantAcceptGuarantee(ctx context.Context, data *request.RequestMerchantAcceptGuarantee) (*response.ResponseMerchantAcceptGuarantee, error) {
	result := &response.ResponseMerchantAcceptGuarantee{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIMerchantAcceptGuarantee, params, nil, nil, result)

	return result, err
}

// MerchantModifyGuarantee 商家协商保障单
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_merchantmodifyguarantee.html
func (comp *Client) MerchantModifyGuarantee(ctx context.Context, data *request.RequestMerchantModifyGuarantee) (*response.ResponseMerchantModifyGuarantee, error) {
	result := &response.ResponseMerchantModifyGuarantee{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIMerchantModifyGuarantee, params, nil, nil, result)

	return result, err
}

// MerchantProofGuarantee 商家举证保障单
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_merchantproofguarantee.html
func (comp *Client) MerchantProofGuarantee(ctx context.Context, data *request.RequestMerchantProofGuarantee) (*response.ResponseMerchantProofGuarantee, error) {
	result := &response.ResponseMerchantProofGuarantee{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIMerchantProofGuarantee, params, nil, nil, result)

	return result, err
}

// MerchantRefuseGuarantee 商家拒绝保障单申请
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_merchantrefuseguarantee.html
func (comp *Client) MerchantRefuseGuarantee(ctx context.Context, data *request.RequestMerchantRefuseGuarantee) (*response.ResponseMerchantRefuseGuarantee, error) {
	result := &response.ResponseMerchantRefuseGuarantee{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIMerchantRefuseGuarantee, params, nil, nil, result)

	return result, err
}

// MerchantUpdateAftersale 商家协商
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_merchantupdateaftersale.html
func (comp *Client) MerchantUpdateAftersale(ctx context.Context, data *request.RequestMerchantUpdateAftersale) (*response.ResponseMerchantUpdateAftersale, error) {
	result := &response.ResponseMerchantUpdateAftersale{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIMerchantUpdateAftersale, params, nil, nil, result)

	return result, err
}

// GetAftersaleReason 获取全量售后原因
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_getaftersalereason.html
func (comp *Client) GetAftersaleReason(ctx context.Context, data *request.RequestGetAftersaleReason) (*response.ResponseGetAftersaleReason, error) {
	result := &response.ResponseGetAftersaleReason{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetAftersaleReason, params, nil, nil, result)

	return result, err
}

// RejectApply 拒绝售后
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_rejectapply.html
func (comp *Client) RejectApply(ctx context.Context, data *request.RequestRejectApply) (*response.ResponseRejectApply, error) {
	result := &response.ResponseRejectApply{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIRejectApply, params, nil, nil, result)

	return result, err
}

// RejectExchangeReship 换货拒绝发货
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_rejectexchangereship.html
func (comp *Client) RejectExchangeReship(ctx context.Context, data *request.RequestRejectExchangeReship) (*response.ResponseRejectExchangeReship, error) {
	result := &response.ResponseRejectExchangeReship{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIRejectExchangeReship, params, nil, nil, result)

	return result, err
}

// GetAftersaleRejectReason 获取拒绝售后原因
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_getaftersalerejectreason.html
func (comp *Client) GetAftersaleRejectReason(ctx context.Context, data *request.RequestGetAftersaleRejectReason) (*response.ResponseGetAftersaleRejectReason, error) {
	result := &response.ResponseGetAftersaleRejectReason{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetAftersaleRejectReason, params, nil, nil, result)

	return result, err
}

// UploadRefundCertificate 上传退款凭证
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_uploadrefundcertificate.html
func (comp *Client) UploadRefundCertificate(ctx context.Context, data *request.RequestUploadRefundCertificate) (*response.ResponseUploadRefundCertificate, error) {
	result := &response.ResponseUploadRefundCertificate{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIUploadRefundCertificate, params, nil, nil, result)

	return result, err
}

// RefundPriceDiff 代用户发起退差价
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_refundpricediff.html
func (comp *Client) RefundPriceDiff(ctx context.Context, data *request.RequestRefundPriceDiff) (*response.ResponseRefundPriceDiff, error) {
	result := &response.ResponseRefundPriceDiff{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIRefundPriceDiff, params, nil, nil, result)

	return result, err
}

// ApplyVirtualTelnum 售后单兑换虚拟号
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-aftersale/api_applyvirtualtelnum.html
func (comp *Client) ApplyVirtualTelnum(ctx context.Context, data *request.RequestApplyVirtualTelnum) (*response.ResponseApplyVirtualTelnum, error) {
	result := &response.ResponseApplyVirtualTelnum{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIApplyVirtualTelnum, params, nil, nil, result)

	return result, err
}
