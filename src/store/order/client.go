package order

import (
	"context"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// GetOrderList 获取订单列表
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_getorderlist.html
func (comp *Client) GetOrderList(ctx context.Context, data *RequestGetOrderList) (*ResponseGetOrderList, error) {
	result := &ResponseGetOrderList{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetOrderList, params, nil, nil, result)

	return result, err
}

// GetOrder 获取订单详情
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_getorder.html
func (comp *Client) GetOrder(ctx context.Context, data *RequestGetOrder) (*ResponseGetOrder, error) {
	result := &ResponseGetOrder{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetOrder, params, nil, nil, result)

	return result, err
}

// SearchOrder 订单搜索
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_searchorder.html
func (comp *Client) SearchOrder(ctx context.Context, data *RequestSearchOrder) (*ResponseSearchOrder, error) {
	result := &ResponseSearchOrder{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APISearchOrder, params, nil, nil, result)

	return result, err
}

// UpdateOrderPrice 修改订单价格
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_changeorderprice.html
func (comp *Client) UpdateOrderPrice(ctx context.Context, data *RequestUpdateOrderPrice) (*ResponseUpdateOrder, error) {
	result := &ResponseUpdateOrder{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIUpdateOrderPrice, params, nil, nil, result)

	return result, err
}

// UpdateOrderMerchantNote 修改订单备注
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_changemerchantnotes.html
func (comp *Client) UpdateOrderMerchantNote(ctx context.Context, data *RequestUpdateOrderMerchantNote) (*ResponseUpdateOrder, error) {
	result := &ResponseUpdateOrder{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIUpdateOrderMerchantNote, params, nil, nil, result)

	return result, err
}

// UpdateOrderAddress 修改订单地址
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_changeorderaddress.html
func (comp *Client) UpdateOrderAddress(ctx context.Context, data *RequestUpdateOrderAddress) (*ResponseUpdateOrder, error) {
	result := &ResponseUpdateOrder{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIUpdateOrderAddress, params, nil, nil, result)

	return result, err
}

// UpdateOrderDelivery 修改物流信息
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_changedeliveryinfo.html
func (comp *Client) UpdateOrderDelivery(ctx context.Context, data *RequestUpdateOrderDelivery) (*ResponseUpdateOrder, error) {
	result := &ResponseUpdateOrder{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIUpdateOrderDelivery, params, nil, nil, result)

	return result, err
}

// AcceptOrderAddressModify 同意用户修改收货地址申请
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_acceptorderaddressmodifyapply.html
func (comp *Client) AcceptOrderAddressModify(ctx context.Context, data *RequestAcceptOrderAddressModify) (*ResponseUpdateOrder, error) {
	result := &ResponseUpdateOrder{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIAcceptOrderAddressModify, params, nil, nil, result)

	return result, err
}

// RejectOrderAddressModify 拒绝用户修改收货地址申请
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_rejectorderaddressmodifyapply.html
func (comp *Client) RejectOrderAddressModify(ctx context.Context, data *RequestRejectOrderAddressModify) (*ResponseUpdateOrder, error) {
	result := &ResponseUpdateOrder{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIRejectOrderAddressModify, params, nil, nil, result)

	return result, err
}

// DecodeSensitiveInfo 解密订单中的详细收货信息
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_decodesensitiveinfo.html
func (comp *Client) DecodeSensitiveInfo(ctx context.Context, data *RequestDecodeSensitiveInfo) (*ResponseDecodeSensitiveInfo, error) {
	result := &ResponseDecodeSensitiveInfo{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIDecodeSensitiveInfo, params, nil, nil, result)

	return result, err
}
