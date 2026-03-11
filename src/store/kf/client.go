package kf

import (
	"context"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel/response"
	"github.com/jinxin-wang/PowerWeChat/v3/src/store/kf/request"
	respBase "github.com/jinxin-wang/PowerWeChat/v3/src/store/kf/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// COSUpload 上传多媒体资源
// https://developers.weixin.qq.com/doc/store/shop/API/kf/api_cosupload.html
func (comp *Client) COSUpload(ctx context.Context, data *request.RequestCOSUpload) (*respBase.ResponseCOSUpload, error) {
	result := &respBase.ResponseCOSUpload{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APICOSUpload, params, nil, nil, result)

	return result, err
}

// SendMsg 发送消息
// https://developers.weixin.qq.com/doc/store/shop/API/kf/api_sendmsg.html
func (comp *Client) SendMsg(ctx context.Context, data *request.RequestSendMsg) (*response.ResponseStore, error) {
	result := &response.ResponseStore{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APISendMsg, params, nil, nil, result)

	return result, err
}
