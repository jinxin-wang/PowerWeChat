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

func (comp *Client) AddComplaintMaterial(ctx context.Context, data *request.RequestAddComplaintMaterial) (*respBase.ResponseAddComplaintMaterial, error) {
	result := &respBase.ResponseAddComplaintMaterial{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIAddComplaintMaterial, params, nil, nil, result)

	return result, err
}

func (comp *Client) AddComplaintProof(ctx context.Context, data *request.RequestAddComplaintProof) (*respBase.ResponseAddComplaintProof, error) {
	result := &respBase.ResponseAddComplaintProof{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIAddComplaintProof, params, nil, nil, result)

	return result, err
}

func (comp *Client) GetComplaintOrder(ctx context.Context, data *request.RequestGetComplaintOrder) (*respBase.ResponseGetComplaintOrder, error) {
	result := &respBase.ResponseGetComplaintOrder{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetComplaintOrder, params, nil, nil, result)

	return result, err
}
