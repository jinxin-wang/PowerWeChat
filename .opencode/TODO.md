# TODO.md

## store 通用模块开发

### Phase 1: 准备工作
- [x] 1. [kernel/response/wx.go] 新增 ResponseStore 类型
- [x] 2. [PROJECT.md] 更新 store 模块入口

### Phase 2: store 应用入口
- [x] 3. [src/store/application.go] 创建应用入口文件结构
- [x] 4. [src/store/application.go] 实现 NewStore 函数
- [x] 5. [src/store/application.go] 实现 MapUserConfig 函数
- [x] 6. [src/store/application.go] 实现 GetContainer/GetAccessToken/GetConfig/GetComponent

### Phase 3: base 通用模块
#### 基础设施
- [x] 7. [src/store/base/doc.go] 包文档注释
- [x] 8. [src/store/base/const.go] API URL 常量定义
- [x] 9. [src/store/base/provider.go] IOC 注册
- [x] 10. [src/store/base/response/response.go] 响应结构体定义

#### API 1: 获取稳定版接口调用凭据
- [x] 11. [request] 创建 GetStableAccessToken - 请求参数：grant_type, appid, secret
- [x] 12. [response] 添加 ResponseGetStableAccessToken - 响应字段：access_token, expires_in
- [x] 13. [client] 实现 GetStableAccessToken 方法 - 调用 cgi-bin/stable_token
- [x] 14. [test] 编写单元测试

#### API 2: 查询API调用额度
- [x] 15. [request] 创建 GetAPIQuota - 请求参数：cgi_path（可选）
- [x] 16. [response] 添加 ResponseGetAPIQuota - 响应字段：quota, quota_limit
- [x] 17. [client] 实现 GetAPIQuota 方法 - 调用 wxa/api_getapiquota
- [x] 18. [test] 编写单元测试

#### API 3: 重置指定API调用次数
- [x] 19. [request] 创建 ClearAPIQuota - 请求参数：cgi_path
- [x] 20. [response] 使用基础响应
- [x] 21. [client] 实现 ClearAPIQuota 方法 - 调用 wxa/api_clearapiquota
- [x] 22. [test] 编写单元测试

#### API 4: 重置API调用次数
- [x] 23. [request] 无需请求结构体
- [x] 24. [response] 使用基础响应
- [x] 25. [client] 实现 ClearQuota 方法 - 调用 wxa/api_clearquota
- [x] 26. [test] 编写单元测试

#### API 5: 使用AppSecret重置API调用次数
- [x] 27. [request] 创建 ClearQuotaByAppSecret - 请求参数：secret
- [x] 28. [response] 使用基础响应
- [x] 29. [client] 实现 ClearQuotaByAppSecret 方法 - 调用 wxa/api_clearquotabyappsecret
- [x] 30. [test] 编写单元测试

#### API 6: 网络通信检测
- [x] 31. [request] 创建 CallbackCheck - 请求参数：action, operator
- [x] 32. [response] 添加 ResponseCallbackCheck - 响应字段：operator, result
- [x] 33. [client] 实现 CallbackCheck 方法 - 调用 wxa/api_callbackcheck
- [x] 34. [test] 编写单元测试

#### API 7: 获取微信API服务器IP
- [x] 35. [request] 无需请求结构体
- [x] 36. [response] 添加 ResponseGetAPIDomainIP - 响应字段：ip_list
- [x] 37. [client] 实现 GetAPIDomainIP 方法 - 调用 cgi-bin/get_api_domain_ip
- [x] 38. [test] 编写单元测试

#### API 8: 获取微信推送服务器IP
- [x] 39. [request] 无需请求结构体
- [x] 40. [response] 添加 ResponseGetCallbackIP - 响应字段：ip_list
- [x] 41. [client] 实现 GetCallbackIP 方法 - 调用 cgi-bin/getcallbackip
- [x] 42. [test] 编写单元测试

#### API 9: 查询rid信息
- [x] 43. [request] 创建 GetRidInfo - 请求参数：rid
- [x] 44. [response] 添加 ResponseGetRidInfo - 响应字段：request_info, request_msg
- [x] 45. [client] 实现 GetRidInfo 方法 - 调用 wxa/api_getridinfo
- [x] 46. [test] 编写单元测试

#### API 10: 通过mediaid获取数据
- [x] 47. [request] 创建 GetDataByMediaID - 请求参数：media_id
- [x] 48. [response] 添加 ResponseGetDataByMediaID - 响应字段：data
- [x] 49. [client] 实现 GetDataByMediaID 方法 - 调用 wxa/api_getdatabymediaid
- [x] 50. [test] 编写单元测试

#### Provider 注册
- [x] 51. [src/store/application.go] 注册 Base 到 Store 应用

### Phase 4: 验证
- [x] 52. go build ./... - 编译验证
- [x] 53. go vet ./... - 代码检查
- [x] 54. go test ./src/store/... - 运行单元测试 (15个测试全部通过)

---

## store/manage 店铺管理模块开发

### 准备工作
- [x] 1. [src/store/manage/doc.go] 包文档注释
- [x] 2. [src/store/manage/const.go] API URL 常量定义

### API 1: 获取店铺基本信息
- [x] 3. [response] 添加 ResponseGetShopBasicInfo - 响应字段：nickname, headimg_url, subject_type, status, username
- [x] 4. [client] 实现 GetShopBasicInfo 方法 - 调用 channels/ec/basics/info/get
- [x] 5. [test] 编写单元测试

### API 2: 获取店铺H5链接
- [x] 6. [response] 添加 ResponseGetShopH5URL - 响应字段：h5_url
- [x] 7. [client] 实现 GetShopH5URL 方法 - 调用 channels/ec/basics/shopurl/get
- [x] 8. [test] 编写单元测试

### API 3: 获取店铺二维码
- [x] 9. [request] 添加 RequestGetShopQRCode - 请求参数：wecom_corp_id, wecom_user_id
- [x] 10. [response] 添加 ResponseGetShopQRCode - 响应字段：shop_qrcode
- [x] 11. [client] 实现 GetShopQRCode 方法 - 调用 channels/ec/basics/shop/qrcode/get
- [x] 12. [test] 编写单元测试

### API 4: 获取店铺口令
- [x] 13. [request] 添加 RequestGetShopTagLink - 请求参数：wecom_corp_id, wecom_user_id
- [x] 14. [response] 添加 ResponseGetShopTagLink - 响应字段：tag_link
- [x] 15. [client] 实现 GetShopTagLink 方法 - 调用 channels/ec/basics/shoptaglink/get
- [x] 16. [test] 编写单元测试

### Provider 注册
- [x] 17. [src/store/manage/provider.go] IOC 注册
- [x] 18. [src/store/application.go] 注册 Manage 到 Store 应用

### 验证
- [x] 19. go build ./src/store/... - 编译验证
- [x] 20. go test ./src/store/manage/... - 运行单元测试 (9个测试全部通过)

---

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
