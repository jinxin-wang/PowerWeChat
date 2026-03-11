package vip

import (
	"context"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel"
	"github.com/jinxin-wang/PowerWeChat/v3/src/store/vip/request"
	respVip "github.com/jinxin-wang/PowerWeChat/v3/src/store/vip/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// GetVIPUserScore 获取用户积分
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/vip/api_getvipuserscore.html
func (comp *Client) GetVIPUserScore(ctx context.Context, data *request.RequestGetVIPUserScore) (*respVip.ResponseGetVIPUserScore, error) {
	result := &respVip.ResponseGetVIPUserScore{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetVIPUserScore, params, nil, nil, result)

	return result, err
}

// GetUserInfo 获取用户信息
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/vip/api_getuserinfo.html
func (comp *Client) GetUserInfo(ctx context.Context, data *request.RequestGetUserInfo) (*respVip.ResponseGetUserInfo, error) {
	result := &respVip.ResponseGetUserInfo{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetUserInfo, params, nil, nil, result)

	return result, err
}

// GetUserList 获取用户列表
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/vip/api_getuserlist.html
func (comp *Client) GetUserList(ctx context.Context, data *request.RequestGetUserList) (*respVip.ResponseGetUserList, error) {
	result := &respVip.ResponseGetUserList{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetUserList, params, nil, nil, result)

	return result, err
}

// GetUserScoreFlowRecord 获取用户积分流水
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/vip/api_getuserscoreflowrecord.html
func (comp *Client) GetUserScoreFlowRecord(ctx context.Context, data *request.RequestGetUserScoreFlowRecord) (*respVip.ResponseGetUserScoreFlowRecord, error) {
	result := &respVip.ResponseGetUserScoreFlowRecord{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetUserScoreFlowRecord, params, nil, nil, result)

	return result, err
}
