package base

import (
	"context"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

func (comp *Client) GetStableAccessToken(ctx context.Context, data *RequestGetStableAccessToken) (*ResponseGetStableAccessToken, error) {
	result := &ResponseGetStableAccessToken{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpGet(ctx, APIGetStableAccessToken, params, nil, result)

	return result, err
}

func (comp *Client) GetAPIQuota(ctx context.Context, data *RequestGetAPIQuota) (*ResponseGetAPIQuota, error) {
	result := &ResponseGetAPIQuota{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetAPIQuota, params, nil, nil, result)

	return result, err
}

func (comp *Client) ClearAPIQuota(ctx context.Context, data *RequestClearAPIQuota) (*response.ResponseStore, error) {
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

func (comp *Client) ClearQuotaByAppSecret(ctx context.Context, data *RequestClearQuotaByAppSecret) (*response.ResponseStore, error) {
	result := &response.ResponseStore{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIClearQuotaByAppSecret, params, nil, nil, result)

	return result, err
}

func (comp *Client) CallbackCheck(ctx context.Context, data *RequestCallbackCheck) (*ResponseCallbackCheck, error) {
	result := &ResponseCallbackCheck{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APICallbackCheck, params, nil, nil, result)

	return result, err
}

func (comp *Client) GetAPIDomainIP(ctx context.Context) (*ResponseGetAPIDomainIP, error) {
	result := &ResponseGetAPIDomainIP{}

	_, err := comp.BaseClient.HttpGet(ctx, APIGetAPIDomainIP, nil, nil, result)

	return result, err
}

func (comp *Client) GetCallbackIP(ctx context.Context) (*ResponseGetCallbackIP, error) {
	result := &ResponseGetCallbackIP{}

	_, err := comp.BaseClient.HttpGet(ctx, APIGetCallbackIP, nil, nil, result)

	return result, err
}

func (comp *Client) GetRidInfo(ctx context.Context, data *RequestGetRidInfo) (*ResponseGetRidInfo, error) {
	result := &ResponseGetRidInfo{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpGet(ctx, APIGetRidInfo, params, nil, result)

	return result, err
}

func (comp *Client) GetDataByMediaID(ctx context.Context, data *RequestGetDataByMediaID) (*ResponseGetDataByMediaID, error) {
	result := &ResponseGetDataByMediaID{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpGet(ctx, APIGetDataByMediaID, params, nil, result)

	return result, err
}
