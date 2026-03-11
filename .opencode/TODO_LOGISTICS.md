# store/logistics 物流发货模块开发任务

## 模块概述
- **模块名称**: logistics (物流发货)
- **API总数**: 31个
- **子模块**: 地址管理(5个) + 运费模板(4个) + 电子面单(15个) + 发货(4个) + 虚拟号码(3个)
- **API路径前缀**: `channels/ec/logistics/`
- **代码结构规范**: 使用 `request/` 和 `response/` 文件夹，每个结构体一个文件

---

## 准备工作
- [x] 1. [src/store/logistics/doc.go] 包文档注释
- [x] 2. [src/store/logistics/const.go] 31个API URL常量定义
- [x] 3. [src/store/logistics/request/] 创建请求结构体文件夹
- [x] 4. [src/store/logistics/response/] 创建响应结构体文件夹
- [x] 5. [src/store/logistics/client.go] 客户端框架
- [x] 6. [src/store/logistics/provider.go] IOC注册

---

## 地址管理 (5个API)

### API 1: 添加地址
- [x] 7. [request/requestAddAddress.go] RequestAddAddress - 请求参数：receiver_name, tel, province, city, district, detail
- [x] 8. [response/responseAddAddress.go] ResponseAddAddress - 响应字段：address_id
- [x] 9. [client.go] AddAddress 方法 - 调用 channels/ec/logistics/address/add
- [x] 10. [client_test.go] 编写单元测试

### API 2: 获取地址列表
- [x] 11. [request/requestGetAddressList.go] RequestGetAddressList - 请求参数：page_size, next_key
- [x] 12. [response/responseGetAddressList.go] ResponseGetAddressList - 响应字段：address_list, next_key, has_more
- [x] 13. [client.go] GetAddressList 方法 - 调用 channels/ec/logistics/address/list/get
- [x] 14. [client_test.go] 编写单元测试

### API 3: 获取地址详情
- [x] 15. [request/requestGetAddress.go] RequestGetAddress - 请求参数：address_id
- [x] 16. [response/responseGetAddress.go] ResponseGetAddress - 响应字段：address_info
- [x] 17. [client.go] GetAddress 方法 - 调用 channels/ec/logistics/address/get
- [x] 18. [client_test.go] 编写单元测试

### API 4: 更新地址
- [x] 19. [request/requestUpdateAddress.go] RequestUpdateAddress - 请求参数：address_id, receiver_name, tel等
- [x] 20. [response/responseUpdateAddress.go] ResponseUpdateAddress - 基础响应
- [x] 21. [client.go] UpdateAddress 方法 - 调用 channels/ec/logistics/address/update
- [x] 22. [client_test.go] 编写单元测试

### API 5: 删除地址
- [x] 23. [request/requestDeleteAddress.go] RequestDeleteAddress - 请求参数：address_id
- [x] 24. [response/responseDeleteAddress.go] ResponseDeleteAddress - 基础响应
- [x] 25. [client.go] DeleteAddress 方法 - 调用 channels/ec/logistics/address/delete
- [x] 26. [client_test.go] 编写单元测试

---

## 运费模板 (4个API)

### API 6: 增加运费模版
- [x] 27. [request/requestAddFreightTemplate.go] RequestAddFreightTemplate - 请求参数：name, freight_info
- [x] 28. [response/responseAddFreightTemplate.go] ResponseAddFreightTemplate - 响应字段：template_id
- [x] 29. [client.go] AddFreightTemplate 方法 - 调用 channels/ec/logistics/merchant/freighttemplate/add
- [x] 30. [client_test.go] 编写单元测试

### API 7: 查询运费模版
- [x] 31. [request/requestGetFreightTemplateDetail.go] RequestGetFreightTemplateDetail - 请求参数：template_id
- [x] 32. [response/responseGetFreightTemplateDetail.go] ResponseGetFreightTemplateDetail - 响应字段：template_info
- [x] 33. [client.go] GetFreightTemplateDetail 方法 - 调用 channels/ec/logistics/merchant/freighttemplate/get
- [x] 34. [client_test.go] 编写单元测试

### API 8: 获取运费模板列表
- [x] 35. [request/requestGetFreightTemplateList.go] RequestGetFreightTemplateList - 请求参数：page_size, next_key
- [x] 36. [response/responseGetFreightTemplateList.go] ResponseGetFreightTemplateList - 响应字段：template_list, next_key, has_more
- [x] 37. [client.go] GetFreightTemplateList 方法 - 调用 channels/ec/logistics/merchant/freighttemplate/list/get
- [x] 38. [client_test.go] 编写单元测试

### API 9: 更新运费模版
- [x] 39. [request/requestUpdateFreightTemplate.go] RequestUpdateFreightTemplate - 请求参数：template_id, name, freight_info
- [x] 40. [response/responseUpdateFreightTemplate.go] ResponseUpdateFreightTemplate - 基础响应
- [x] 41. [client.go] UpdateFreightTemplate 方法 - 调用 channels/ec/logistics/merchant/freighttemplate/update
- [x] 42. [client_test.go] 编写单元测试

---

## 电子面单 (15个API)

### API 10: 获取面单标准模板
- [x] 43. [request/requestEwaybillGetTemplateConfig.go] RequestEwaybillGetTemplateConfig - 请求参数：无
- [x] 44. [response/responseEwaybillGetTemplateConfig.go] ResponseEwaybillGetTemplateConfig - 响应字段：template_configs
- [x] 45. [client.go] EwaybillGetTemplateConfig 方法 - 调用 channels/ec/logistics/ewaybill/get_template_config
- [x] 46. [client_test.go] 编写单元测试

### API 11: 新增面单模板
- [x] 47. [request/requestEwaybillCreateTemplate.go] RequestEwaybillCreateTemplate - 请求参数：config_info
- [x] 48. [response/responseEwaybillCreateTemplate.go] ResponseEwaybillCreateTemplate - 响应字段：template_id
- [x] 49. [client.go] EwaybillCreateTemplate 方法 - 调用 channels/ec/logistics/ewaybill/createtemplate
- [x] 50. [client_test.go] 编写单元测试

### API 12: 删除面单模版
- [x] 51. [request/requestEwaybillDelTemplate.go] RequestEwaybillDelTemplate - 请求参数：template_id
- [x] 52. [response/responseEwaybillDelTemplate.go] ResponseEwaybillDelTemplate - 基础响应
- [x] 53. [client.go] EwaybillDelTemplate 方法 - 调用 channels/ec/logistics/ewaybill/deltemplate
- [x] 54. [client_test.go] 编写单元测试

### API 13: 更新面单模版
- [x] 55. [request/requestEwaybillUpdateTemplate.go] RequestEwaybillUpdateTemplate - 请求参数：template_id, config_info
- [x] 56. [response/responseEwaybillUpdateTemplate.go] ResponseEwaybillUpdateTemplate - 基础响应
- [x] 57. [client.go] EwaybillUpdateTemplate 方法 - 调用 channels/ec/logistics/ewaybill/updatetemplate
- [x] 58. [client_test.go] 编写单元测试

### API 14: 获取面单模板信息
- [x] 59. [request/requestEwaybillGetTemplate.go] RequestEwaybillGetTemplate - 请求参数：template_code
- [x] 60. [response/responseEwaybillGetTemplate.go] ResponseEwaybillGetTemplate - 响应字段：template_info
- [x] 61. [client.go] EwaybillGetTemplate 方法 - 调用 channels/ec/logistics/ewaybill/gettemplate
- [x] 62. [client_test.go] 编写单元测试

### API 15: 根据模板ID获取面单模板信息
- [x] 63. [request/requestEwaybillGetTemplateByID.go] RequestEwaybillGetTemplateByID - 请求参数：template_id
- [x] 64. [response/responseEwaybillGetTemplateByID.go] ResponseEwaybillGetTemplateByID - 响应字段：template_info
- [x] 65. [client.go] EwaybillGetTemplateByID 方法 - 调用 channels/ec/logistics/ewaybill/gettemplatebyid
- [x] 66. [client_test.go] 编写单元测试

### API 16: 查询开通的电子面单网点/账号信息
- [x] 67. [request/requestEwaybillGetAcct.go] RequestEwaybillGetAcct - 请求参数：无
- [x] 68. [response/responseEwaybillGetAcct.go] ResponseEwaybillGetAcct - 响应字段：acct_list
- [x] 69. [client.go] EwaybillGetAcct 方法 - 调用 channels/ec/logistics/ewaybill/getacct
- [x] 70. [client_test.go] 编写单元测试

### API 17: 查询开通的快递公司列表
- [x] 71. [request/requestEwaybillGetDeliveryList.go] RequestEwaybillGetDeliveryList - 请求参数：无
- [x] 72. [response/responseEwaybillGetDeliveryList.go] ResponseEwaybillGetDeliveryList - 响应字段：delivery_list
- [x] 73. [client.go] EwaybillGetDeliveryList 方法 - 调用 channels/ec/logistics/ewaybill/getdeliverylist
- [x] 74. [client_test.go] 编写单元测试

### API 18: 电子面单预取号
- [x] 75. [request/requestEwaybillPrecreateOrder.go] RequestEwaybillPrecreateOrder - 请求参数：delivery_id, waybill_info
- [x] 76. [response/responseEwaybillPrecreateOrder.go] ResponseEwaybillPrecreateOrder - 响应字段：waybill_id
- [x] 77. [client.go] EwaybillPrecreateOrder 方法 - 调用 channels/ec/logistics/ewaybill/precreateorder
- [x] 78. [client_test.go] 编写单元测试

### API 19: 电子面单取号
- [x] 79. [request/requestEwaybillCreateOrder.go] RequestEwaybillCreateOrder - 请求参数：delivery_id, order_id, sender_info, receiver_info
- [x] 80. [response/responseEwaybillCreateOrder.go] ResponseEwaybillCreateOrder - 响应字段：waybill_id, waybill_data
- [x] 81. [client.go] EwaybillCreateOrder 方法 - 调用 channels/ec/logistics/ewaybill/createorder
- [x] 82. [client_test.go] 编写单元测试

### API 20: 电子面单子件追加
- [x] 83. [request/requestEwaybillAddSubOrder.go] RequestEwaybillAddSubOrder - 请求参数：waybill_id, sub_order_info
- [x] 84. [response/responseEwaybillAddSubOrder.go] ResponseEwaybillAddSubOrder - 基础响应
- [x] 85. [client.go] EwaybillAddSubOrder 方法 - 调用 channels/ec/logistics/ewaybill/addsuborder
- [x] 86. [client_test.go] 编写单元测试

### API 21: 电子面单取消下单
- [x] 87. [request/requestEwaybillCancelOrder.go] RequestEwaybillCancelOrder - 请求参数：waybill_id, order_id
- [x] 88. [response/responseEwaybillCancelOrder.go] ResponseEwaybillCancelOrder - 基础响应
- [x] 89. [client.go] EwaybillCancelOrder 方法 - 调用 channels/ec/logistics/ewaybill/cancelorder
- [x] 90. [client_test.go] 编写单元测试

### API 22: 查询面单详情
- [x] 91. [request/requestEwaybillGetOrder.go] RequestEwaybillGetOrder - 请求参数：waybill_id
- [x] 92. [response/responseEwaybillGetOrder.go] ResponseEwaybillGetOrder - 响应字段：waybill_info
- [x] 93. [client.go] EwaybillGetOrder 方法 - 调用 channels/ec/logistics/ewaybill/getorder
- [x] 94. [client_test.go] 编写单元测试

### API 23: 获取打印报文
- [x] 95. [request/requestEwaybillGetPrintContent.go] RequestEwaybillGetPrintContent - 请求参数：waybill_ids
- [x] 96. [response/responseEwaybillGetPrintContent.go] ResponseEwaybillGetPrintContent - 响应字段：print_content_list
- [x] 97. [client.go] EwaybillGetPrintContent 方法 - 调用 channels/ec/logistics/ewaybill/get_print_content
- [x] 98. [client_test.go] 编写单元测试

### API 24: 打印成功通知
- [x] 99. [request/requestEwaybillPrintOrder.go] RequestEwaybillPrintOrder - 请求参数：waybill_id
- [x] 100. [response/responseEwaybillPrintOrder.go] ResponseEwaybillPrintOrder - 基础响应
- [x] 101. [client.go] EwaybillPrintOrder 方法 - 调用 channels/ec/logistics/ewaybill/printorder
- [x] 102. [client_test.go] 编写单元测试

### API 25: 批量打印通知
- [x] 103. [request/requestEwaybillBatchPrintOrder.go] RequestEwaybillBatchPrintOrder - 请求参数：waybill_ids
- [x] 104. [response/responseEwaybillBatchPrintOrder.go] ResponseEwaybillBatchPrintOrder - 基础响应
- [x] 105. [client.go] EwaybillBatchPrintOrder 方法 - 调用 channels/ec/logistics/ewaybill/batchprintorder
- [x] 106. [client_test.go] 编写单元测试

---

## 发货 (4个API)

### API 26: 订单发货
- [x] 107. [request/requestSendDelivery.go] RequestSendDelivery - 请求参数：order_id, delivery_list
- [x] 108. [response/responseSendDelivery.go] ResponseSendDelivery - 基础响应
- [x] 109. [client.go] SendDelivery 方法 - 调用 channels/ec/logistics/delivery/send
- [x] 110. [client_test.go] 编写单元测试

### API 27: 订单补发货
- [x] 111. [request/requestDeliveryCompensation.go] RequestDeliveryCompensation - 请求参数：order_id, delivery_list
- [x] 112. [response/responseDeliveryCompensation.go] ResponseDeliveryCompensation - 基础响应
- [x] 113. [client.go] DeliveryCompensation 方法 - 调用 channels/ec/logistics/delivery/compensation
- [x] 114. [client_test.go] 编写单元测试

### API 28: 获取快递公司列表
- [x] 115. [request/requestGetDeliveryCompanyListNew.go] RequestGetDeliveryCompanyListNew - 请求参数：无
- [x] 116. [response/responseGetDeliveryCompanyListNew.go] ResponseGetDeliveryCompanyListNew - 响应字段：company_list
- [x] 117. [client.go] GetDeliveryCompanyListNew 方法 - 调用 channels/ec/logistics/delivery/getcompanylist
- [x] 118. [client_test.go] 编写单元测试

### API 29: 获取快递公司列表-旧
- [x] 119. [request/requestGetDeliveryCompanyList.go] RequestGetDeliveryCompanyList - 请求参数：无
- [x] 120. [response/responseGetDeliveryCompanyList.go] ResponseGetDeliveryCompanyList - 响应字段：company_list
- [x] 121. [client.go] GetDeliveryCompanyList 方法 - 调用 channels/ec/logistics/delivery/getcompanylistold
- [x] 122. [client_test.go] 编写单元测试

---

## 物流公司虚拟号码 (3个API)

### API 30: 获取虚拟号码池
- [x] 123. [request/requestGetPrivateNumberPool.go] RequestGetPrivateNumberPool - 请求参数：无
- [x] 124. [response/responseGetPrivateNumberPool.go] ResponseGetPrivateNumberPool - 响应字段：number_pool
- [x] 125. [client.go] GetPrivateNumberPool 方法 - 调用 channels/ec/logistics/phonenumber/pool/get
- [x] 126. [client_test.go] 编写单元测试

### API 31: 根据运单号获取真实手机号
- [x] 127. [request/requestGetRealNumber.go] RequestGetRealNumber - 请求参数：waybill_id
- [x] 128. [response/responseGetRealNumber.go] ResponseGetRealNumber - 响应字段：phone_number
- [x] 129. [client.go] GetRealNumber 方法 - 调用 channels/ec/logistics/phonenumber/real/get
- [x] 130. [client_test.go] 编写单元测试

### API 32: 根据运单号获取虚拟手机号
- [x] 131. [request/requestGetVirtualNumber.go] RequestGetVirtualNumber - 请求参数：waybill_id
- [x] 132. [response/responseGetVirtualNumber.go] ResponseGetVirtualNumber - 响应字段：virtual_number
- [x] 133. [client.go] GetVirtualNumber 方法 - 调用 channels/ec/logistics/phonenumber/virtual/get
- [x] 134. [client_test.go] 编写单元测试

---

## Provider 注册
- [x] 135. [src/store/logistics/provider.go] IOC 注册（已有准备工作6）
- [x] 136. [src/store/application.go] 注册 Logistics 到 Store 应用

---

## 验证
- [x] 137. go build ./src/store/... - 编译验证
- [x] 138. go vet ./src/store/... - 代码检查
- [x] 139. go test ./src/store/logistics/... - 运行单元测试
