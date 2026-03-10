package order

import (
	"context"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
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

// UploadFreshInsurance 上传生鲜商品质检信息
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_uploadfreshinsurance.html
func (comp *Client) UploadFreshInsurance(ctx context.Context, data *RequestUploadFreshInsurance) (*response.ResponseStore, error) {
	result := &response.ResponseStore{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIUploadFreshInsurance, params, nil, nil, result)

	return result, err
}

// AddGiftOrderNote 添加礼品订单备注
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_addgiftordernote.html
func (comp *Client) AddGiftOrderNote(ctx context.Context, data *RequestAddGiftOrderNote) (*response.ResponseStore, error) {
	result := &response.ResponseStore{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIAddGiftOrderNote, params, nil, nil, result)

	return result, err
}

// GetGiftOrderSubList 获取礼品订单子单列表
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_getgiftordersublist.html
func (comp *Client) GetGiftOrderSubList(ctx context.Context, data *RequestGetGiftOrderSubList) (*ResponseGetGiftOrderSubList, error) {
	result := &ResponseGetGiftOrderSubList{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetGiftOrderSubList, params, nil, nil, result)

	return result, err
}

// GetSKUChangeList 获取待发货SKU变更列表
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_getskuchangelist.html
func (comp *Client) GetSKUChangeList(ctx context.Context, data *RequestGetSKUChangeList) (*ResponseGetSKUChangeList, error) {
	result := &ResponseGetSKUChangeList{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetSKUChangeList, params, nil, nil, result)

	return result, err
}

// AcceptSKUChange 接受SKU变更请求
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_acceptskuchange.html
func (comp *Client) AcceptSKUChange(ctx context.Context, data *RequestAcceptSKUChange) (*response.ResponseStore, error) {
	result := &response.ResponseStore{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIAcceptSKUChange, params, nil, nil, result)

	return result, err
}

// RejectSKUChange 拒绝SKU变更请求
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_rejectskuchange.html
func (comp *Client) RejectSKUChange(ctx context.Context, data *RequestRejectSKUChange) (*response.ResponseStore, error) {
	result := &response.ResponseStore{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIRejectSKUChange, params, nil, nil, result)

	return result, err
}

// ApplyRealNumber 申请查看真实号码
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_applyrealnumber.html
func (comp *Client) ApplyRealNumber(ctx context.Context, data *RequestApplyRealNumber) (*ResponseApplyRealNumber, error) {
	result := &ResponseApplyRealNumber{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIApplyRealNumber, params, nil, nil, result)

	return result, err
}

// GetRealNumberStatus 查询真实号码审核状态
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_getrealnumberstatus.html
func (comp *Client) GetRealNumberStatus(ctx context.Context, data *RequestGetRealNumberStatus) (*ResponseGetRealNumberStatus, error) {
	result := &ResponseGetRealNumberStatus{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetRealNumberStatus, params, nil, nil, result)

	return result, err
}

// ReapplyVirtualNumber 重新申请虚拟号
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_reapplyvirtualnumber.html
func (comp *Client) ReapplyVirtualNumber(ctx context.Context, data *RequestReapplyVirtualNumber) (*response.ResponseStore, error) {
	result := &response.ResponseStore{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIReapplyVirtualNumber, params, nil, nil, result)

	return result, err
}

// DelayVirtualNumber 延期虚拟号有效期
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_delayvirtualnumber.html
func (comp *Client) DelayVirtualNumber(ctx context.Context, data *RequestDelayVirtualNumber) (*response.ResponseStore, error) {
	result := &response.ResponseStore{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIDelayVirtualNumber, params, nil, nil, result)

	return result, err
}

// AddPhoneVerifyCode 添加手机号用于核验
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_addphoneverifycode.html
func (comp *Client) AddPhoneVerifyCode(ctx context.Context, data *RequestAddPhoneVerifyCode) (*response.ResponseStore, error) {
	result := &response.ResponseStore{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIAddPhoneVerifyCode, params, nil, nil, result)

	return result, err
}

// SendPhoneVerifyCode 获取短信验证码
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_sendphoneverifycode.html
func (comp *Client) SendPhoneVerifyCode(ctx context.Context, data *RequestSendPhoneVerifyCode) (*response.ResponseStore, error) {
	result := &response.ResponseStore{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APISendPhoneVerifyCode, params, nil, nil, result)

	return result, err
}

// GetPhoneStatus 获取店铺手机号核验状态
// https://developers.weixin.qq.com/doc/store/shop/API/channels-shop-order/api_getphonestatus.html
func (comp *Client) GetPhoneStatus(ctx context.Context, data *RequestGetPhoneStatus) (*ResponseGetPhoneStatus, error) {
	result := &ResponseGetPhoneStatus{}

	params, err := object.StructToStringMap(data)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, APIGetPhoneStatus, params, nil, nil, result)

	return result, err
}
