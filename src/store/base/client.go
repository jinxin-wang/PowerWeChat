package base

import (
	"context"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/store/base/request"
	respBase "github.com/ArtisanCloud/PowerWeChat/v3/src/store/base/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

func (comp *Client) GetStableAccessToken(ctx context.Context, data *request.RequestGetStableAccessToken) (*respBase.ResponseGetStableAccessToken, error) {
	result := &respBase.ResponseGetStableAccessToken{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpGet(ctx, APIGetStableAccessToken, params, nil, result)

	return result, err
}

func (comp *Client) GetAPIQuota(ctx context.Context, data *request.RequestGetAPIQuota) (*respBase.ResponseGetAPIQuota, error) {
	result := &respBase.ResponseGetAPIQuota{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetAPIQuota, params, nil, nil, result)

	return result, err
}

func (comp *Client) ClearAPIQuota(ctx context.Context, data *request.RequestClearAPIQuota) (*response.ResponseStore, error) {
	result := &response.ResponseStore{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIClearAPIQuota, params, nil, nil, result)

	return result, err
}

func (comp *Client) ClearQuota(ctx context.Context) (*response.ResponseStore, error) {
	result := &response.ResponseStore{}

	_, err := comp.BaseClient.HttpPost(ctx, APIClearQuota, nil, nil, result)

	return result, err
}

func (comp *Client) ClearQuotaByAppSecret(ctx context.Context, data *request.RequestClearQuotaByAppSecret) (*response.ResponseStore, error) {
	result := &response.ResponseStore{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIClearQuotaByAppSecret, params, nil, nil, result)

	return result, err
}

func (comp *Client) CallbackCheck(ctx context.Context, data *request.RequestCallbackCheck) (*respBase.ResponseCallbackCheck, error) {
	result := &respBase.ResponseCallbackCheck{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APICallbackCheck, params, nil, nil, result)

	return result, err
}

func (comp *Client) GetAPIDomainIP(ctx context.Context) (*respBase.ResponseGetAPIDomainIP, error) {
	result := &respBase.ResponseGetAPIDomainIP{}

	_, err := comp.BaseClient.HttpGet(ctx, APIGetAPIDomainIP, nil, nil, result)

	return result, err
}

func (comp *Client) GetCallbackIP(ctx context.Context) (*respBase.ResponseGetCallbackIP, error) {
	result := &respBase.ResponseGetCallbackIP{}

	_, err := comp.BaseClient.HttpGet(ctx, APIGetCallbackIP, nil, nil, result)

	return result, err
}

func (comp *Client) GetRidInfo(ctx context.Context, data *request.RequestGetRidInfo) (*respBase.ResponseGetRidInfo, error) {
	result := &respBase.ResponseGetRidInfo{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpGet(ctx, APIGetRidInfo, params, nil, result)

	return result, err
}

func (comp *Client) GetDataByMediaID(ctx context.Context, data *request.RequestGetDataByMediaID) (*respBase.ResponseGetDataByMediaID, error) {
	result := &respBase.ResponseGetDataByMediaID{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpGet(ctx, APIGetDataByMediaID, params, nil, result)

	return result, err
}
