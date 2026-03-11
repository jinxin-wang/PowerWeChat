package manage

import (
	"context"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel"
	"github.com/jinxin-wang/PowerWeChat/v3/src/store/manage/request"
	"github.com/jinxin-wang/PowerWeChat/v3/src/store/manage/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// GetShopBasicInfo 获取店铺基本信息
// https://developers.weixin.qq.com/doc/store/shop/API/storemanage/api_mmecapi_basicinfo.html
func (comp *Client) GetShopBasicInfo(ctx context.Context) (*response.ResponseGetShopBasicInfo, error) {
	result := &response.ResponseGetShopBasicInfo{}

	_, err := comp.BaseClient.HttpGet(ctx, APIGetShopBasicInfo, nil, nil, result)

	return result, err
}

// GetShopH5URL 获取店铺H5链接
// https://developers.weixin.qq.com/doc/store/shop/API/storemanage/api_getshoph5url.html
func (comp *Client) GetShopH5URL(ctx context.Context) (*response.ResponseGetShopH5URL, error) {
	result := &response.ResponseGetShopH5URL{}

	_, err := comp.BaseClient.HttpGet(ctx, APIGetShopH5URL, nil, nil, result)

	return result, err
}

// GetShopQRCode 获取店铺二维码
// https://developers.weixin.qq.com/doc/store/shop/API/storemanage/api_getshopqrcode.html
func (comp *Client) GetShopQRCode(ctx context.Context, data *request.RequestGetShopQRCode) (*response.ResponseGetShopQRCode, error) {
	result := &response.ResponseGetShopQRCode{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetShopQRCode, params, nil, nil, result)

	return result, err
}

// GetShopTagLink 获取店铺口令
// https://developers.weixin.qq.com/doc/store/shop/API/storemanage/api_getshoptaglink.html
func (comp *Client) GetShopTagLink(ctx context.Context, data *request.RequestGetShopTagLink) (*response.ResponseGetShopTagLink, error) {
	result := &response.ResponseGetShopTagLink{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetShopTagLink, params, nil, nil, result)

	return result, err
}
