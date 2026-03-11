## store/order 订单管理模块开发（第一批 - 9个核心API）

### 准备工作
- [x] 1. [src/store/order/doc.go] 包文档注释
- [x] 2. [src/store/order/const.go] 9个API URL常量定义
- [x] 3. [src/store/order/request.go] 请求结构体
- [x] 4. [src/store/order/response.go] 响应结构体

### API 1: 获取订单列表
- [x] 5. [request] RequestGetOrderList - 请求参数：page_size, next_key, status, start_time, end_time
- [x] 6. [response] ResponseGetOrderList - 响应字段：order_list, next_key, has_more
- [x] 7. [client] GetOrderList 方法 - 调用 channels/ec/order/list/get
- [x] 8. [test] 编写单元测试

### API 2: 获取订单详情
- [x] 9. [request] RequestGetOrder - 请求参数：order_id
- [x] 10. [response] ResponseGetOrder - 响应字段：order
- [x] 11. [client] GetOrder 方法 - 调用 channels/ec/order/get
- [x] 12. [test] 编写单元测试

### API 3: 订单搜索
- [x] 13. [request] RequestSearchOrder - 请求参数：keyword, page_size, next_key
- [x] 14. [response] ResponseSearchOrder - 响应字段：order_list, next_key, has_more
- [x] 15. [client] SearchOrder 方法 - 调用 channels/ec/order/search
- [x] 16. [test] 编写单元测试

### API 4: 修改订单价格
- [x] 17. [request] RequestUpdateOrderPrice - 请求参数：order_id, change_type, price, remark
- [x] 18. [response] ResponseUpdateOrder - 基础响应
- [x] 19. [client] UpdateOrderPrice 方法 - 调用 channels/ec/order/price/update
- [x] 20. [test] 编写单元测试

### API 5: 修改订单备注
- [x] 21. [request] RequestUpdateOrderMerchantNote - 请求参数：order_id, note
- [x] 22. [response] ResponseUpdateOrder - 基础响应
- [x] 23. [client] UpdateOrderMerchantNote 方法 - 调用 channels/ec/order/merchantnote/update
- [x] 24. [test] 编写单元测试

### API 6: 修改订单地址
- [x] 25. [request] RequestUpdateOrderAddress - 请求参数：order_id, receiver_name, detail_info等
- [x] 26. [response] ResponseUpdateOrder - 基础响应
- [x] 27. [client] UpdateOrderAddress 方法 - 调用 channels/ec/order/address/update
- [x] 28. [test] 编写单元测试

### API 7: 修改物流信息
- [x] 29. [request] RequestUpdateOrderDelivery - 请求参数：order_id, delivery_id, waybill_id等
- [x] 30. [response] ResponseUpdateOrder - 基础响应
- [x] 31. [client] UpdateOrderDelivery 方法 - 调用 channels/ec/order/delivery/update
- [x] 32. [test] 编写单元测试

### API 8: 同意/拒绝用户修改收货地址申请
- [x] 33. [request] RequestAcceptOrderAddressModify - 请求参数：order_id
- [x] 34. [response] ResponseUpdateOrder - 基础响应
- [x] 35. [client] AcceptOrderAddressModify 方法 - 调用 channels/ec/order/address/accept
- [x] 36. [test] 编写单元测试
- [x] 37. [request] RequestRejectOrderAddressModify - 请求参数：order_id, reason
- [x] 38. [response] ResponseUpdateOrder - 基础响应
- [x] 39. [client] RejectOrderAddressModify 方法 - 调用 channels/ec/order/address/reject
- [x] 40. [test] 编写单元测试

### API 9: 解密订单中的详细收货信息
- [x] 41. [request] RequestDecodeSensitiveInfo - 请求参数：order_id, encrypted_data
- [x] 42. [response] ResponseDecodeSensitiveInfo - 响应字段：receiver_name, tel_number, detail_info
- [x] 43. [client] DecodeSensitiveInfo 方法 - 调用 channels/ec/order/sensitiveinfo/decode
- [x] 44. [test] 编写单元测试

### Provider 注册
- [x] 45. [src/store/order/provider.go] IOC 注册
- [x] 46. [src/store/application.go] 注册 Order 到 Store 应用

### 验证
- [x] 47. go build ./src/store/... - 编译验证
- [x] 48. go test ./src/store/order/... - 运行单元测试

---

## store/order 订单管理模块开发（第二批 - 13个API）

### 准备工作
- [x] 1. [src/store/order/const.go] 添加13个API URL常量
- [x] 2. [src/store/order/request.go] 请求结构体
- [x] 3. [src/store/order/response.go] 响应结构体

### API 1: 上传生鲜商品质检信息
- [x] 4. [request] RequestUploadFreshInsurance - 请求参数：order_id, insurance_info
- [x] 5. [response] ResponseUploadFreshInsurance - 基础响应
- [x] 6. [client] UploadFreshInsurance 方法 - 调用 channels/ec/order/fresh/insurance
- [x] 7. [test] 编写单元测试

### API 2: 添加礼品订单备注
- [x] 8. [request] RequestAddGiftOrderNote - 请求参数：order_id, note
- [x] 9. [response] ResponseAddGiftOrderNote - 基础响应
- [x] 10. [client] AddGiftOrderNote 方法 - 调用 channels/ec/order/present/note
- [x] 11. [test] 编写单元测试

### API 3: 获取礼品订单子单列表
- [x] 12. [request] RequestGetGiftOrderSubList - 请求参数：order_id, page_size, next_key
- [x] 13. [response] ResponseGetGiftOrderSubList - 响应字段：sub_order_list, next_key, has_more
- [x] 14. [client] GetGiftOrderSubList 方法 - 调用 channels/ec/order/present/sublist
- [x] 15. [test] 编写单元测试

### API 4: 获取待发货SKU变更列表
- [x] 16. [request] RequestGetSKUChangeList - 请求参数：order_id, page_size, next_key
- [x] 17. [response] ResponseGetSKUChangeList - 响应字段：sku_change_list, next_key, has_more
- [x] 18. [client] GetSKUChangeList 方法 - 调用 channels/ec/order/sku/change/list
- [x] 19. [test] 编写单元测试

### API 5: 接受SKU变更请求
- [x] 20. [request] RequestAcceptSKUChange - 请求参数：order_id, sku_change_id
- [x] 21. [response] ResponseAcceptSKUChange - 基础响应
- [x] 22. [client] AcceptSKUChange 方法 - 调用 channels/ec/order/sku/change/accept
- [x] 23. [test] 编写单元测试

### API 6: 拒绝SKU变更请求
- [x] 24. [request] RequestRejectSKUChange - 请求参数：order_id, sku_change_id, reason
- [x] 25. [response] ResponseRejectSKUChange - 基础响应
- [x] 26. [client] RejectSKUChange 方法 - 调用 channels/ec/order/sku/change/reject
- [x] 27. [test] 编写单元测试

### API 7: 申请查看真实号码
- [x] 28. [request] RequestApplyRealNumber - 请求参数：order_id, reason
- [x] 29. [response] ResponseApplyRealNumber - 响应字段：apply_id
- [x] 30. [client] ApplyRealNumber 方法 - 调用 channels/ec/order/realnumber/apply
- [x] 31. [test] 编写单元测试

### API 8: 查询真实号码审核状态
- [x] 32. [request] RequestGetRealNumberStatus - 请求参数：apply_id
- [x] 33. [response] ResponseGetRealNumberStatus - 响应字段：status, phone_number, audit_result
- [x] 34. [client] GetRealNumberStatus 方法 - 调用 channels/ec/order/realnumber/status
- [x] 35. [test] 编写单元测试

### API 9: 重新申请虚拟号
- [x] 36. [request] RequestReapplyVirtualNumber - 请求参数：order_id
- [x] 37. [response] ResponseReapplyVirtualNumber - 基础响应
- [x] 38. [client] ReapplyVirtualNumber 方法 - 调用 channels/ec/order/virtualnumber/apply
- [x] 39. [test] 编写单元测试

### API 10: 延期虚拟号有效期
- [x] 40. [request] RequestDelayVirtualNumber - 请求参数：order_id, delay_days
- [x] 41. [response] ResponseDelayVirtualNumber - 基础响应
- [x] 42. [client] DelayVirtualNumber 方法 - 调用 channels/ec/order/virtualnumber/delay
- [x] 43. [test] 编写单元测试

### API 11: 添加手机号用于核验
- [x] 44. [request] RequestAddPhoneVerifyCode - 请求参数：order_id, phone
- [x] 45. [response] ResponseAddPhoneVerifyCode - 基础响应
- [x] 46. [client] AddPhoneVerifyCode 方法 - 调用 channels/ec/order/phone/verifycode/add
- [x] 47. [test] 编写单元测试

### API 12: 获取短信验证码
- [x] 48. [request] RequestSendPhoneVerifyCode - 请求参数：order_id
- [x] 49. [response] ResponseSendPhoneVerifyCode - 基础响应
- [x] 50. [client] SendPhoneVerifyCode 方法 - 调用 channels/ec/order/phone/verifycode/send
- [x] 51. [test] 编写单元测试

### API 13: 获取店铺手机号核验状态
- [x] 52. [request] RequestGetPhoneStatus - 请求参数：order_id
- [x] 53. [response] ResponseGetPhoneStatus - 响应字段：status, phone
- [x] 54. [client] GetPhoneStatus 方法 - 调用 channels/ec/order/phone/status
- [x] 55. [test] 编写单元测试

### 验证
- [x] 56. go build ./src/store/... - 编译验证
- [x] 57. go vet ./src/store/... - 代码检查
- [x] 58. go test ./src/store/order/... - 运行单元测试（18个新测试全部通过）
