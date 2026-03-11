package base

import (
	"context"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
	"github.com/jinxin-wang/PowerWeChat/v3/src/store/base/request"
	respBase "github.com/jinxin-wang/PowerWeChat/v3/src/store/base/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// GetStableAccessToken 获取稳定版接口调用凭据
// https://developers.weixin.qq.com/doc/store/shop/API/apimgnt/api_getstableaccesstoken.html
func (comp *Client) GetStableAccessToken(ctx context.Context, data *request.RequestGetStableAccessToken) (*respBase.ResponseGetStableAccessToken, error) {
	result := &respBase.ResponseGetStableAccessToken{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpGet(ctx, APIGetStableAccessToken, params, nil, result)

	return result, err
}

// GetAPIQuota 查询API调用额度
// https://developers.weixin.qq.com/doc/store/shop/API/apimgnt/api_getapiquota.html
func (comp *Client) GetAPIQuota(ctx context.Context, data *request.RequestGetAPIQuota) (*respBase.ResponseGetAPIQuota, error) {
	result := &respBase.ResponseGetAPIQuota{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetAPIQuota, params, nil, nil, result)

	return result, err
}

// ClearAPIQuota 重置指定API调用次数
// https://developers.weixin.qq.com/doc/store/shop/API/apimgnt/api_clearapiquota.html
func (comp *Client) ClearAPIQuota(ctx context.Context, data *request.RequestClearAPIQuota) (*response.ResponseStore, error) {
	result := &response.ResponseStore{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIClearAPIQuota, params, nil, nil, result)

	return result, err
}

// ClearQuota 重置API调用次数
// https://developers.weixin.qq.com/doc/store/shop/API/apimgnt/api_clearquota.html
func (comp *Client) ClearQuota(ctx context.Context) (*response.ResponseStore, error) {
	result := &response.ResponseStore{}

	_, err := comp.BaseClient.HttpPost(ctx, APIClearQuota, nil, nil, result)

	return result, err
}

// ClearQuotaByAppSecret 使用AppSecret重置API调用次数
// https://developers.weixin.qq.com/doc/store/shop/API/apimgnt/api_clearquotabyappsecret.html
func (comp *Client) ClearQuotaByAppSecret(ctx context.Context, data *request.RequestClearQuotaByAppSecret) (*response.ResponseStore, error) {
	result := &response.ResponseStore{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIClearQuotaByAppSecret, params, nil, nil, result)

	return result, err
}

// CallbackCheck 网络通信检测
// https://developers.weixin.qq.com/doc/store/shop/API/apimgnt/api_callbackcheck.html
func (comp *Client) CallbackCheck(ctx context.Context, data *request.RequestCallbackCheck) (*respBase.ResponseCallbackCheck, error) {
	result := &respBase.ResponseCallbackCheck{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APICallbackCheck, params, nil, nil, result)

	return result, err
}

// GetAPIDomainIP 获取微信API服务器IP
// https://developers.weixin.qq.com/doc/store/shop/API/apimgnt/api_getapidomainip.html
func (comp *Client) GetAPIDomainIP(ctx context.Context) (*respBase.ResponseGetAPIDomainIP, error) {
	result := &respBase.ResponseGetAPIDomainIP{}

	_, err := comp.BaseClient.HttpGet(ctx, APIGetAPIDomainIP, nil, nil, result)

	return result, err
}

// GetCallbackIP 获取微信推送服务器IP
// https://developers.weixin.qq.com/doc/store/shop/API/apimgnt/api_getcallbackip.html
func (comp *Client) GetCallbackIP(ctx context.Context) (*respBase.ResponseGetCallbackIP, error) {
	result := &respBase.ResponseGetCallbackIP{}

	_, err := comp.BaseClient.HttpGet(ctx, APIGetCallbackIP, nil, nil, result)

	return result, err
}

// GetRidInfo 查询rid信息
// https://developers.weixin.qq.com/doc/store/shop/API/apimgnt/api_getridinfo.html
func (comp *Client) GetRidInfo(ctx context.Context, data *request.RequestGetRidInfo) (*respBase.ResponseGetRidInfo, error) {
	result := &respBase.ResponseGetRidInfo{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpGet(ctx, APIGetRidInfo, params, nil, result)

	return result, err
}

// GetDataByMediaID 通过mediaid获取数据
// https://developers.weixin.qq.com/doc/store/shop/API/apimgnt/api_getdatabymediaid.html
func (comp *Client) GetDataByMediaID(ctx context.Context, data *request.RequestGetDataByMediaID) (*respBase.ResponseGetDataByMediaID, error) {
	result := &respBase.ResponseGetDataByMediaID{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpGet(ctx, APIGetDataByMediaID, params, nil, result)

	return result, err
}
