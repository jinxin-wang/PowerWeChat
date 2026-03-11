package complaint

import (
	"context"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel"
	"github.com/jinxin-wang/PowerWeChat/v3/src/store/complaint/request"
	respBase "github.com/jinxin-wang/PowerWeChat/v3/src/store/complaint/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// AddComplaintMaterial 商家补充纠纷单留言
// https://developers.weixin.qq.com/doc/store/shop/API/complaint/api_addcomplaintmaterial.html
func (comp *Client) AddComplaintMaterial(ctx context.Context, data *request.RequestAddComplaintMaterial) (*respBase.ResponseAddComplaintMaterial, error) {
	result := &respBase.ResponseAddComplaintMaterial{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIAddComplaintMaterial, params, nil, nil, result)

	return result, err
}

// AddComplaintProof 商家举证
// https://developers.weixin.qq.com/doc/store/shop/API/complaint/api_addcomplaintproof.html
func (comp *Client) AddComplaintProof(ctx context.Context, data *request.RequestAddComplaintProof) (*respBase.ResponseAddComplaintProof, error) {
	result := &respBase.ResponseAddComplaintProof{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIAddComplaintProof, params, nil, nil, result)

	return result, err
}

// GetComplaintOrder 获取纠纷单
// https://developers.weixin.qq.com/doc/store/shop/API/complaint/api_getcomplaintorder.html
func (comp *Client) GetComplaintOrder(ctx context.Context, data *request.RequestGetComplaintOrder) (*respBase.ResponseGetComplaintOrder, error) {
	result := &respBase.ResponseGetComplaintOrder{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetComplaintOrder, params, nil, nil, result)

	return result, err
}
