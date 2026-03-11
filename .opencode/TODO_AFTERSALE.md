# store/aftersale 售后管理模块开发任务

## 模块概述
- **模块名称**: aftersale (售后管理)
- **API总数**: 17个
- **API路径前缀**: `channels/ec/aftersale/`

---

## 准备工作
- [ ] 1. [src/store/aftersale/doc.go] 包文档注释
- [ ] 2. [src/store/aftersale/const.go] 17个API URL常量定义
- [ ] 3. [src/store/aftersale/request.go] 请求结构体
- [ ] 4. [src/store/aftersale/response.go] 响应结构体

---

## API 1: 获取售后单列表
- [ ] 5. [request] RequestGetAftersaleList - 请求参数：page_size, next_key, status, start_time, end_time
- [ ] 6. [response] ResponseGetAftersaleList - 响应字段：aftersale_list, next_key, has_more
- [ ] 7. [client] GetAftersaleList 方法 - 调用 channels/ec/aftersale/getaftersalelist
- [ ] 8. [test] 编写单元测试

## API 2: 获取售后单详情
- [ ] 9. [request] RequestGetAftersaleOrder - 请求参数：aftersale_id
- [ ] 10. [response] ResponseGetAftersaleOrder - 响应字段：aftersale_info
- [ ] 11. [client] GetAftersaleOrder 方法 - 调用 channels/ec/aftersale/getaftersaleorder
- [ ] 12. [test] 编写单元测试

## API 3: 同意售后
- [ ] 13. [request] RequestAcceptApply - 请求参数：aftersale_id, address_id
- [ ] 14. [response] ResponseAcceptApply - 基础响应
- [ ] 15. [client] AcceptApply 方法 - 调用 channels/ec/aftersale/acceptapply
- [ ] 16. [test] 编写单元测试

## API 4: 换货发货
- [ ] 17. [request] RequestAcceptExchangeReship - 请求参数：aftersale_id, delivery_id, waybill_id
- [ ] 18. [response] ResponseAcceptExchangeReship - 基础响应
- [ ] 19. [client] AcceptExchangeReship 方法 - 调用 channels/ec/aftersale/acceptexchangereship
- [ ] 20. [test] 编写单元测试

## API 5: 代用户发起售后
- [ ] 21. [request] RequestGenAftersaleOrder - 请求参数：order_id, product_id, type, reason
- [ ] 22. [response] ResponseGenAftersaleOrder - 响应字段：aftersale_id
- [ ] 23. [client] GenAftersaleOrder 方法 - 调用 channels/ec/aftersale/genaftersaleorder
- [ ] 24. [test] 编写单元测试

## API 6: 商家获取保障单列表
- [ ] 25. [request] RequestSearchGuaranteeOrder - 请求参数：page_size, next_key, status
- [ ] 26. [response] ResponseSearchGuaranteeOrder - 响应字段：guarantee_list, next_key, has_more
- [ ] 27. [client] SearchGuaranteeOrder 方法 - 调用 channels/ec/aftersale/searchguaranteeorder
- [ ] 28. [test] 编写单元测试

## API 7: 获取保障单详情
- [ ] 29. [request] RequestGetGuaranteeOrder - 请求参数：guarantee_id
- [ ] 30. [response] ResponseGetGuaranteeOrder - 响应字段：guarantee_info
- [ ] 31. [client] GetGuaranteeOrder 方法 - 调用 channels/ec/aftersale/getguaranteeorder
- [ ] 32. [test] 编写单元测试

## API 8: 商家同意保障单申请
- [ ] 33. [request] RequestMerchantAcceptGuarantee - 请求参数：guarantee_id
- [ ] 34. [response] ResponseMerchantAcceptGuarantee - 基础响应
- [ ] 35. [client] MerchantAcceptGuarantee 方法 - 调用 channels/ec/aftersale/merchantacceptguarantee
- [ ] 36. [test] 编写单元测试

## API 9: 商家协商保障单
- [ ] 37. [request] RequestMerchantModifyGuarantee - 请求参数：guarantee_id, amount
- [ ] 38. [response] ResponseMerchantModifyGuarantee - 基础响应
- [ ] 39. [client] MerchantModifyGuarantee 方法 - 调用 channels/ec/aftersale/merchantmodifyguarantee
- [ ] 40. [test] 编写单元测试

## API 10: 商家举证保障单
- [ ] 41. [request] RequestMerchantProofGuarantee - 请求参数：guarantee_id, proof_info
- [ ] 42. [response] ResponseMerchantProofGuarantee - 基础响应
- [ ] 43. [client] MerchantProofGuarantee 方法 - 调用 channels/ec/aftersale/merchantproofguarantee
- [ ] 44. [test] 编写单元测试

## API 11: 商家拒绝保障单申请
- [ ] 45. [request] RequestMerchantRefuseGuarantee - 请求参数：guarantee_id, reason
- [ ] 46. [response] ResponseMerchantRefuseGuarantee - 基础响应
- [ ] 47. [client] MerchantRefuseGuarantee 方法 - 调用 channels/ec/aftersale/merchantrefuseguarantee
- [ ] 48. [test] 编写单元测试

## API 12: 商家协商
- [ ] 49. [request] RequestMerchantUpdateAftersale - 请求参数：aftersale_id, refund_amount
- [ ] 50. [response] ResponseMerchantUpdateAftersale - 基础响应
- [ ] 51. [client] MerchantUpdateAftersale 方法 - 调用 channels/ec/aftersale/merchantupdateaftersale
- [ ] 52. [test] 编写单元测试

## API 13: 获取全量售后原因
- [ ] 53. [request] RequestGetAftersaleReason - 请求参数：无
- [ ] 54. [response] ResponseGetAftersaleReason - 响应字段：reason_list
- [ ] 55. [client] GetAftersaleReason 方法 - 调用 channels/ec/aftersale/getaftersalereason
- [ ] 56. [test] 编写单元测试

## API 14: 拒绝售后
- [ ] 57. [request] RequestRejectApply - 请求参数：aftersale_id, reason
- [ ] 58. [response] ResponseRejectApply - 基础响应
- [ ] 59. [client] RejectApply 方法 - 调用 channels/ec/aftersale/rejectapply
- [ ] 60. [test] 编写单元测试

## API 15: 换货拒绝发货
- [ ] 61. [request] RequestRejectExchangeReship - 请求参数：aftersale_id, reason
- [ ] 62. [response] ResponseRejectExchangeReship - 基础响应
- [ ] 63. [client] RejectExchangeReship 方法 - 调用 channels/ec/aftersale/rejectexchangereship
- [ ] 64. [test] 编写单元测试

## API 16: 获取拒绝售后原因
- [ ] 65. [request] RequestGetAftersaleRejectReason - 请求参数：无
- [ ] 66. [response] ResponseGetAftersaleRejectReason - 响应字段：reason_list
- [ ] 67. [client] GetAftersaleRejectReason 方法 - 调用 channels/ec/aftersale/getaftersalerejectreason
- [ ] 68. [test] 编写单元测试

## API 17: 上传退款凭证
- [ ] 69. [request] RequestUploadRefundCertificate - 请求参数：aftersale_id, certificate_info
- [ ] 70. [response] ResponseUploadRefundCertificate - 基础响应
- [ ] 71. [client] UploadRefundCertificate 方法 - 调用 channels/ec/aftersale/uploadrefundcertificate
- [ ] 72. [test] 编写单元测试

## API 18: 代用户发起退差价
- [ ] 73. [request] RequestRefundPriceDiff - 请求参数：order_id, price_diff_amount
- [ ] 74. [response] ResponseRefundPriceDiff - 响应字段：aftersale_id
- [ ] 75. [client] RefundPriceDiff 方法 - 调用 channels/ec/aftersale/refundpricediff
- [ ] 76. [test] 编写单元测试

## API 19: 售后单兑换虚拟号
- [ ] 77. [request] RequestApplyVirtualTelnum - 请求参数：aftersale_id
- [ ] 78. [response] ResponseApplyVirtualTelnum - 基础响应
- [ ] 79. [client] ApplyVirtualTelnum 方法 - 调用 channels/ec/aftersale/applyvirtualtelnum
- [ ] 80. [test] 编写单元测试

---

## Provider 注册
- [ ] 81. [src/store/aftersale/provider.go] IOC 注册
- [ ] 82. [src/store/application.go] 注册 Aftersale 到 Store 应用

---

## 验证
- [ ] 83. go build ./src/store/... - 编译验证
- [ ] 84. go vet ./src/store/... - 代码检查
- [ ] 85. go test ./src/store/aftersale/... - 运行单元测试
