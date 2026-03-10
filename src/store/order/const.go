package order

const (
	APIGetOrderList             = "channels/ec/order/list/get"
	APIGetOrder                 = "channels/ec/order/get"
	APISearchOrder              = "channels/ec/order/search"
	APIUpdateOrderPrice         = "channels/ec/order/price/update"
	APIUpdateOrderMerchantNote  = "channels/ec/order/merchantnote/update"
	APIUpdateOrderAddress       = "channels/ec/order/address/update"
	APIUpdateOrderDelivery      = "channels/ec/order/delivery/update"
	APIAcceptOrderAddressModify = "channels/ec/order/address/accept"
	APIRejectOrderAddressModify = "channels/ec/order/address/reject"
	APIDecodeSensitiveInfo      = "channels/ec/order/sensitiveinfo/decode"

	// Batch 2 APIs
	APIUploadFreshInsurance = "channels/ec/order/fresh/insurance"
	APIAddGiftOrderNote     = "channels/ec/order/present/note"
	APIGetGiftOrderSubList  = "channels/ec/order/present/sublist"
	APIGetSKUChangeList     = "channels/ec/order/sku/change/list"
	APIAcceptSKUChange      = "channels/ec/order/sku/change/accept"
	APIRejectSKUChange      = "channels/ec/order/sku/change/reject"
	APIApplyRealNumber      = "channels/ec/order/realnumber/apply"
	APIGetRealNumberStatus  = "channels/ec/order/realnumber/status"
	APIReapplyVirtualNumber = "channels/ec/order/virtualnumber/apply"
	APIDelayVirtualNumber   = "channels/ec/order/virtualnumber/delay"
	APIAddPhoneVerifyCode   = "channels/ec/order/phone/verifycode/add"
	APISendPhoneVerifyCode  = "channels/ec/order/phone/verifycode/send"
	APIGetPhoneStatus       = "channels/ec/order/phone/status"
)
