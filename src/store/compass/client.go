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

// GetShopFinderAuthorizationList 获取授权视频号列表
// https://developers.weixin.qq.com/doc/store/shop/API/compass/api_getshopfinderauthorizationlist.html
func (comp *Client) GetShopFinderAuthorizationList(ctx context.Context, data *request.RequestGetShopFinderAuthorizationList) (*respCompass.ResponseGetShopFinderAuthorizationList, error) {
	result := &respCompass.ResponseGetShopFinderAuthorizationList{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetShopFinderAuthorizationList, params, nil, nil, result)

	return result, err
}

// GetShopFinderList 获取带货达人列表
// https://developers.weixin.qq.com/doc/store/shop/API/compass/api_getshopfinderlist.html
func (comp *Client) GetShopFinderList(ctx context.Context, data *request.RequestGetShopFinderList) (*respCompass.ResponseGetShopFinderList, error) {
	result := &respCompass.ResponseGetShopFinderList{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetShopFinderList, params, nil, nil, result)

	return result, err
}

// GetShopFinderOverall 获取带货数据概览
// https://developers.weixin.qq.com/doc/store/shop/API/compass/api_getshopfinderoverall.html
func (comp *Client) GetShopFinderOverall(ctx context.Context, data *request.RequestGetShopFinderOverall) (*respCompass.ResponseGetShopFinderOverall, error) {
	result := &respCompass.ResponseGetShopFinderOverall{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetShopFinderOverall, params, nil, nil, result)

	return result, err
}

// GetShopFinderProductList 获取带货达人商品列表
// https://developers.weixin.qq.com/doc/store/shop/API/compass/api_getshopfinderproductlist.html
func (comp *Client) GetShopFinderProductList(ctx context.Context, data *request.RequestGetShopFinderProductList) (*respCompass.ResponseGetShopFinderProductList, error) {
	result := &respCompass.ResponseGetShopFinderProductList{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetShopFinderProductList, params, nil, nil, result)

	return result, err
}

// GetShopFinderProductOverall 获取带货达人详情
// https://developers.weixin.qq.com/doc/store/shop/API/compass/api_getshopfinderproductoverall.html
func (comp *Client) GetShopFinderProductOverall(ctx context.Context, data *request.RequestGetShopFinderProductOverall) (*respCompass.ResponseGetShopFinderProductOverall, error) {
	result := &respCompass.ResponseGetShopFinderProductOverall{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetShopFinderProductOverall, params, nil, nil, result)

	return result, err
}

// GetShopLiveList 获取店铺开播列表
// https://developers.weixin.qq.com/doc/store/shop/API/compass/api_getshoplivelist.html
func (comp *Client) GetShopLiveList(ctx context.Context, data *request.RequestGetShopLiveList) (*respCompass.ResponseGetShopLiveList, error) {
	result := &respCompass.ResponseGetShopLiveList{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetShopLiveList, params, nil, nil, result)

	return result, err
}

// GetShopOverall 获取电商数据概览
// https://developers.weixin.qq.com/doc/store/shop/API/compass/api_getshopoverall.html
func (comp *Client) GetShopOverall(ctx context.Context, data *request.RequestGetShopOverall) (*respCompass.ResponseGetShopOverall, error) {
	result := &respCompass.ResponseGetShopOverall{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetShopOverall, params, nil, nil, result)

	return result, err
}

// GetShopProductData 获取商品详细信息
// https://developers.weixin.qq.com/doc/store/shop/API/compass/api_getshopproductdata.html
func (comp *Client) GetShopProductData(ctx context.Context, data *request.RequestGetShopProductData) (*respCompass.ResponseGetShopProductData, error) {
	result := &respCompass.ResponseGetShopProductData{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetShopProductData, params, nil, nil, result)

	return result, err
}

// GetShopProductList 获取商品列表
// https://developers.weixin.qq.com/doc/store/shop/API/compass/api_getshopproductlist.html
func (comp *Client) GetShopProductList(ctx context.Context, data *request.RequestGetShopProductList) (*respCompass.ResponseGetShopProductList, error) {
	result := &respCompass.ResponseGetShopProductList{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetShopProductList, params, nil, nil, result)

	return result, err
}

// GetShopSaleProfileData 获取店铺人群数据
// https://developers.weixin.qq.com/doc/store/shop/API/compass/api_getshopsaleprofiledata.html
func (comp *Client) GetShopSaleProfileData(ctx context.Context, data *request.RequestGetShopSaleProfileData) (*respCompass.ResponseGetShopSaleProfileData, error) {
	result := &respCompass.ResponseGetShopSaleProfileData{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetShopSaleProfileData, params, nil, nil, result)

	return result, err
}
