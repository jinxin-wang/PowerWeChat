package league

import (
	"context"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/jinxin-wang/PowerWeChat/v3/src/kernel"
	"github.com/jinxin-wang/PowerWeChat/v3/src/store/league/request"
	"github.com/jinxin-wang/PowerWeChat/v3/src/store/league/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// AddPromoter 新增达人
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/promoter/api_addpromoter.html
func (comp *Client) AddPromoter(ctx context.Context, data *request.RequestAddPromoter) (*response.ResponseAddPromoter, error) {
	result := &response.ResponseAddPromoter{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIAddPromoter, params, nil, nil, result)

	return result, err
}

// DeletePromoter 删除达人
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/promoter/api_deletepromoter.html
func (comp *Client) DeletePromoter(ctx context.Context, data *request.RequestDeletePromoter) (*response.ResponseDeletePromoter, error) {
	result := &response.ResponseDeletePromoter{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIDeletePromoter, params, nil, nil, result)

	return result, err
}

// GetPromoter 获取达人详情信息
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/promoter/api_getpromoter.html
func (comp *Client) GetPromoter(ctx context.Context, data *request.RequestGetPromoter) (*response.ResponseGetPromoter, error) {
	result := &response.ResponseGetPromoter{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetPromoter, params, nil, nil, result)

	return result, err
}

// GetPromoterList 获取商店达人列表
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/promoter/api_getpromoterlist.html
func (comp *Client) GetPromoterList(ctx context.Context, data *request.RequestGetPromoterList) (*response.ResponseGetPromoterList, error) {
	result := &response.ResponseGetPromoterList{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetPromoterList, params, nil, nil, result)

	return result, err
}

// UpdatePromoter 编辑达人
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/promoter/api_updpromoter.html
func (comp *Client) UpdatePromoter(ctx context.Context, data *request.RequestUpdatePromoter) (*response.ResponseUpdatePromoter, error) {
	result := &response.ResponseUpdatePromoter{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIUpdatePromoter, params, nil, nil, result)

	return result, err
}

// BatchAddItem 批量新增联盟商品
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/item/api_batchadditem.html
func (comp *Client) BatchAddItem(ctx context.Context, data *request.RequestBatchAddItem) (*response.ResponseBatchAddItem, error) {
	result := &response.ResponseBatchAddItem{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIBatchAddItem, params, nil, nil, result)

	return result, err
}

// DeleteItem 删除联盟商品
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/item/api_deleteitem.html
func (comp *Client) DeleteItem(ctx context.Context, data *request.RequestDeleteItem) (*response.ResponseDeleteItem, error) {
	result := &response.ResponseDeleteItem{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIDeleteItem, params, nil, nil, result)

	return result, err
}

// GetItem 获取联盟商品详情
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/item/api_getitem.html
func (comp *Client) GetItem(ctx context.Context, data *request.RequestGetItem) (*response.ResponseGetItem, error) {
	result := &response.ResponseGetItem{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetItem, params, nil, nil, result)

	return result, err
}

// BatchAddHeadSupplierItem 批量新增联盟机构推广
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/item/api_batchaddheadsupplieritem.html
func (comp *Client) BatchAddHeadSupplierItem(ctx context.Context, data *request.RequestBatchAddHeadSupplierItem) (*response.ResponseBatchAddHeadSupplierItem, error) {
	result := &response.ResponseBatchAddHeadSupplierItem{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIBatchAddHeadSupplierItem, params, nil, nil, result)

	return result, err
}

// GetItemList 获取联盟商品推广列表
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/item/api_getitemlist.html
func (comp *Client) GetItemList(ctx context.Context, data *request.RequestGetItemList) (*response.ResponseGetItemList, error) {
	result := &response.ResponseGetItemList{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetItemList, params, nil, nil, result)

	return result, err
}

// UpdateItem 更新联盟商品信息
// 官方文档: https://developers.weixin.qq.com/doc/store/shop/API/league/item/api_upditem.html
func (comp *Client) UpdateItem(ctx context.Context, data *request.RequestUpdateItem) (*response.ResponseUpdateItem, error) {
	result := &response.ResponseUpdateItem{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIUpdateItem, params, nil, nil, result)

	return result, err
}
