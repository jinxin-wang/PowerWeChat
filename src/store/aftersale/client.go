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

func (comp *Client) GetAftersaleList(ctx context.Context, data *request.RequestGetAftersaleList) (*response.ResponseGetAftersaleList, error) {
	result := &response.ResponseGetAftersaleList{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetAftersaleList, params, nil, nil, result)

	return result, err
}

func (comp *Client) GetAftersaleOrder(ctx context.Context, data *request.RequestGetAftersaleOrder) (*response.ResponseGetAftersaleOrder, error) {
	result := &response.ResponseGetAftersaleOrder{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetAftersaleOrder, params, nil, nil, result)

	return result, err
}

func (comp *Client) AcceptApply(ctx context.Context, data *request.RequestAcceptApply) (*response.ResponseAcceptApply, error) {
	result := &response.ResponseAcceptApply{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIAcceptApply, params, nil, nil, result)

	return result, err
}

func (comp *Client) AcceptExchangeReship(ctx context.Context, data *request.RequestAcceptExchangeReship) (*response.ResponseAcceptExchangeReship, error) {
	result := &response.ResponseAcceptExchangeReship{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIAcceptExchangeReship, params, nil, nil, result)

	return result, err
}

func (comp *Client) GenAftersaleOrder(ctx context.Context, data *request.RequestGenAftersaleOrder) (*response.ResponseGenAftersaleOrder, error) {
	result := &response.ResponseGenAftersaleOrder{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGenAftersaleOrder, params, nil, nil, result)

	return result, err
}

func (comp *Client) SearchGuaranteeOrder(ctx context.Context, data *request.RequestSearchGuaranteeOrder) (*response.ResponseSearchGuaranteeOrder, error) {
	result := &response.ResponseSearchGuaranteeOrder{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APISearchGuaranteeOrder, params, nil, nil, result)

	return result, err
}

func (comp *Client) GetGuaranteeOrder(ctx context.Context, data *request.RequestGetGuaranteeOrder) (*response.ResponseGetGuaranteeOrder, error) {
	result := &response.ResponseGetGuaranteeOrder{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetGuaranteeOrder, params, nil, nil, result)

	return result, err
}

func (comp *Client) MerchantAcceptGuarantee(ctx context.Context, data *request.RequestMerchantAcceptGuarantee) (*response.ResponseMerchantAcceptGuarantee, error) {
	result := &response.ResponseMerchantAcceptGuarantee{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIMerchantAcceptGuarantee, params, nil, nil, result)

	return result, err
}

func (comp *Client) MerchantModifyGuarantee(ctx context.Context, data *request.RequestMerchantModifyGuarantee) (*response.ResponseMerchantModifyGuarantee, error) {
	result := &response.ResponseMerchantModifyGuarantee{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIMerchantModifyGuarantee, params, nil, nil, result)

	return result, err
}

func (comp *Client) MerchantProofGuarantee(ctx context.Context, data *request.RequestMerchantProofGuarantee) (*response.ResponseMerchantProofGuarantee, error) {
	result := &response.ResponseMerchantProofGuarantee{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIMerchantProofGuarantee, params, nil, nil, result)

	return result, err
}

func (comp *Client) MerchantRefuseGuarantee(ctx context.Context, data *request.RequestMerchantRefuseGuarantee) (*response.ResponseMerchantRefuseGuarantee, error) {
	result := &response.ResponseMerchantRefuseGuarantee{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIMerchantRefuseGuarantee, params, nil, nil, result)

	return result, err
}

func (comp *Client) MerchantUpdateAftersale(ctx context.Context, data *request.RequestMerchantUpdateAftersale) (*response.ResponseMerchantUpdateAftersale, error) {
	result := &response.ResponseMerchantUpdateAftersale{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIMerchantUpdateAftersale, params, nil, nil, result)

	return result, err
}

func (comp *Client) GetAftersaleReason(ctx context.Context, data *request.RequestGetAftersaleReason) (*response.ResponseGetAftersaleReason, error) {
	result := &response.ResponseGetAftersaleReason{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetAftersaleReason, params, nil, nil, result)

	return result, err
}

func (comp *Client) RejectApply(ctx context.Context, data *request.RequestRejectApply) (*response.ResponseRejectApply, error) {
	result := &response.ResponseRejectApply{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIRejectApply, params, nil, nil, result)

	return result, err
}

func (comp *Client) RejectExchangeReship(ctx context.Context, data *request.RequestRejectExchangeReship) (*response.ResponseRejectExchangeReship, error) {
	result := &response.ResponseRejectExchangeReship{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIRejectExchangeReship, params, nil, nil, result)

	return result, err
}

func (comp *Client) GetAftersaleRejectReason(ctx context.Context, data *request.RequestGetAftersaleRejectReason) (*response.ResponseGetAftersaleRejectReason, error) {
	result := &response.ResponseGetAftersaleRejectReason{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetAftersaleRejectReason, params, nil, nil, result)

	return result, err
}

func (comp *Client) UploadRefundCertificate(ctx context.Context, data *request.RequestUploadRefundCertificate) (*response.ResponseUploadRefundCertificate, error) {
	result := &response.ResponseUploadRefundCertificate{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIUploadRefundCertificate, params, nil, nil, result)

	return result, err
}

func (comp *Client) RefundPriceDiff(ctx context.Context, data *request.RequestRefundPriceDiff) (*response.ResponseRefundPriceDiff, error) {
	result := &response.ResponseRefundPriceDiff{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIRefundPriceDiff, params, nil, nil, result)

	return result, err
}

func (comp *Client) ApplyVirtualTelnum(ctx context.Context, data *request.RequestApplyVirtualTelnum) (*response.ResponseApplyVirtualTelnum, error) {
	result := &response.ResponseApplyVirtualTelnum{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIApplyVirtualTelnum, params, nil, nil, result)

	return result, err
}
