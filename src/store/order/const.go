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
)
