package logistics

const (
	// 地址管理
	APIAddAddress     = "channels/ec/logistics/address/add"
	APIGetAddressList = "channels/ec/logistics/address/list/get"
	APIGetAddress     = "channels/ec/logistics/address/get"
	APIUpdateAddress  = "channels/ec/logistics/address/update"
	APIDeleteAddress  = "channels/ec/logistics/address/delete"

	// 运费模板
	APIAddFreightTemplate       = "channels/ec/logistics/merchant/freighttemplate/add"
	APIGetFreightTemplateDetail = "channels/ec/logistics/merchant/freighttemplate/get"
	APIGetFreightTemplateList   = "channels/ec/logistics/merchant/freighttemplate/list/get"
	APIUpdateFreightTemplate    = "channels/ec/logistics/merchant/freighttemplate/update"

	// 电子面单
	APIEwaybillGetTemplateConfig = "channels/ec/logistics/ewaybill/get_template_config"
	APIEwaybillCreateTemplate    = "channels/ec/logistics/ewaybill/createtemplate"
	APIEwaybillDelTemplate       = "channels/ec/logistics/ewaybill/deltemplate"
	APIEwaybillUpdateTemplate    = "channels/ec/logistics/ewaybill/updatetemplate"
	APIEwaybillGetTemplate       = "channels/ec/logistics/ewaybill/gettemplate"
	APIEwaybillGetTemplateByID   = "channels/ec/logistics/ewaybill/gettemplatebyid"
	APIEwaybillGetAcct           = "channels/ec/logistics/ewaybill/getacct"
	APIEwaybillGetDeliveryList   = "channels/ec/logistics/ewaybill/getdeliverylist"
	APIEwaybillPrecreateOrder    = "channels/ec/logistics/ewaybill/precreateorder"
	APIEwaybillCreateOrder       = "channels/ec/logistics/ewaybill/createorder"
	APIEwaybillAddSubOrder       = "channels/ec/logistics/ewaybill/addsuborder"
	APIEwaybillCancelOrder       = "channels/ec/logistics/ewaybill/cancelorder"
	APIEwaybillGetOrder          = "channels/ec/logistics/ewaybill/getorder"
	APIEwaybillGetPrintContent   = "channels/ec/logistics/ewaybill/get_print_content"
	APIEwaybillPrintOrder        = "channels/ec/logistics/ewaybill/printorder"
	APIEwaybillBatchPrintOrder   = "channels/ec/logistics/ewaybill/batchprintorder"

	// 发货
	APISendDelivery              = "channels/ec/logistics/delivery/send"
	APIDeliveryCompensation      = "channels/ec/logistics/delivery/compensation"
	APIGetDeliveryCompanyListNew = "channels/ec/logistics/delivery/getcompanylist"
	APIGetDeliveryCompanyList    = "channels/ec/logistics/delivery/getcompanylistold"

	// 虚拟号码
	APIGetPrivateNumberPool = "channels/ec/logistics/phonenumber/pool/get"
	APIGetRealNumber        = "channels/ec/logistics/phonenumber/real/get"
	APIGetVirtualNumber     = "channels/ec/logistics/phonenumber/virtual/get"
)
