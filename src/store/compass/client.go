package compass

import (
	"context"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel"
	"github.com/jinxin-wang/PowerWeChat/v3/src/store/compass/request"
	respCompass "github.com/jinxin-wang/PowerWeChat/v3/src/store/compass/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

func (comp *Client) GetShopFinderAuthorizationList(ctx context.Context, data *request.RequestGetShopFinderAuthorizationList) (*respCompass.ResponseGetShopFinderAuthorizationList, error) {
	result := &respCompass.ResponseGetShopFinderAuthorizationList{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetShopFinderAuthorizationList, params, nil, nil, result)

	return result, err
}

func (comp *Client) GetShopFinderList(ctx context.Context, data *request.RequestGetShopFinderList) (*respCompass.ResponseGetShopFinderList, error) {
	result := &respCompass.ResponseGetShopFinderList{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetShopFinderList, params, nil, nil, result)

	return result, err
}

func (comp *Client) GetShopFinderOverall(ctx context.Context, data *request.RequestGetShopFinderOverall) (*respCompass.ResponseGetShopFinderOverall, error) {
	result := &respCompass.ResponseGetShopFinderOverall{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetShopFinderOverall, params, nil, nil, result)

	return result, err
}

func (comp *Client) GetShopFinderProductList(ctx context.Context, data *request.RequestGetShopFinderProductList) (*respCompass.ResponseGetShopFinderProductList, error) {
	result := &respCompass.ResponseGetShopFinderProductList{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetShopFinderProductList, params, nil, nil, result)

	return result, err
}

func (comp *Client) GetShopFinderProductOverall(ctx context.Context, data *request.RequestGetShopFinderProductOverall) (*respCompass.ResponseGetShopFinderProductOverall, error) {
	result := &respCompass.ResponseGetShopFinderProductOverall{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetShopFinderProductOverall, params, nil, nil, result)

	return result, err
}

func (comp *Client) GetShopLiveList(ctx context.Context, data *request.RequestGetShopLiveList) (*respCompass.ResponseGetShopLiveList, error) {
	result := &respCompass.ResponseGetShopLiveList{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetShopLiveList, params, nil, nil, result)

	return result, err
}

func (comp *Client) GetShopOverall(ctx context.Context, data *request.RequestGetShopOverall) (*respCompass.ResponseGetShopOverall, error) {
	result := &respCompass.ResponseGetShopOverall{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetShopOverall, params, nil, nil, result)

	return result, err
}

func (comp *Client) GetShopProductData(ctx context.Context, data *request.RequestGetShopProductData) (*respCompass.ResponseGetShopProductData, error) {
	result := &respCompass.ResponseGetShopProductData{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetShopProductData, params, nil, nil, result)

	return result, err
}

func (comp *Client) GetShopProductList(ctx context.Context, data *request.RequestGetShopProductList) (*respCompass.ResponseGetShopProductList, error) {
	result := &respCompass.ResponseGetShopProductList{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetShopProductList, params, nil, nil, result)

	return result, err
}

func (comp *Client) GetShopSaleProfileData(ctx context.Context, data *request.RequestGetShopSaleProfileData) (*respCompass.ResponseGetShopSaleProfileData, error) {
	result := &respCompass.ResponseGetShopSaleProfileData{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetShopSaleProfileData, params, nil, nil, result)

	return result, err
}
