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
- [x] 11. [request] 创建 GetStableAccessToken.go - 请求参数：grant_type, appid, secret
- [x] 12. [response] 添加 ResponseGetStableAccessToken - 响应字段：access_token, expires_in
- [x] 13. [client] 实现 GetStableAccessToken 方法 - 调用 cgi-bin/stable_token
- [ ] 14. [test] 编写单元测试

#### API 2: 查询API调用额度
- [x] 15. [request] 创建 GetAPIQuota.go - 请求参数：cgi_path（可选）
- [x] 16. [response] 添加 ResponseGetAPIQuota - 响应字段：quota, quota_limit
- [x] 17. [client] 实现 GetAPIQuota 方法 - 调用 wxa/api_getapiquota
- [ ] 18. [test] 编写单元测试

#### API 3: 重置指定API调用次数
- [x] 19. [request] 创建 ClearAPIQuota.go - 请求参数：cgi_path
- [x] 20. [response] 使用基础响应
- [x] 21. [client] 实现 ClearAPIQuota 方法 - 调用 wxa/api_clearapiquota
- [ ] 22. [test] 编写单元测试

#### API 4: 重置API调用次数
- [x] 23. [request] 无需请求结构体
- [x] 24. [response] 使用基础响应
- [x] 25. [client] 实现 ClearQuota 方法 - 调用 wxa/api_clearquota
- [ ] 26. [test] 编写单元测试

#### API 5: 使用AppSecret重置API调用次数
- [x] 27. [request] 创建 ClearQuotaByAppSecret.go - 请求参数：secret
- [x] 28. [response] 使用基础响应
- [x] 29. [client] 实现 ClearQuotaByAppSecret 方法 - 调用 wxa/api_clearquotabyappsecret
- [ ] 30. [test] 编写单元测试

#### API 6: 网络通信检测
- [x] 31. [request] 创建 CallbackCheck.go - 请求参数：action, operator
- [x] 32. [response] 添加 ResponseCallbackCheck - 响应字段：operator, result
- [x] 33. [client] 实现 CallbackCheck 方法 - 调用 wxa/api_callbackcheck
- [ ] 34. [test] 编写单元测试

#### API 7: 获取微信API服务器IP
- [x] 35. [request] 无需请求结构体
- [x] 36. [response] 添加 ResponseGetAPIDomainIP - 响应字段：ip_list
- [x] 37. [client] 实现 GetAPIDomainIP 方法 - 调用 cgi-bin/get_api_domain_ip
- [ ] 38. [test] 编写单元测试

#### API 8: 获取微信推送服务器IP
- [x] 39. [request] 无需请求结构体
- [x] 40. [response] 添加 ResponseGetCallbackIP - 响应字段：ip_list
- [x] 41. [client] 实现 GetCallbackIP 方法 - 调用 cgi-bin/getcallbackip
- [ ] 42. [test] 编写单元测试

#### API 9: 查询rid信息
- [x] 43. [request] 创建 GetRidInfo.go - 请求参数：rid
- [x] 44. [response] 添加 ResponseGetRidInfo - 响应字段：request_info, request_msg
- [x] 45. [client] 实现 GetRidInfo 方法 - 调用 wxa/api_getridinfo
- [ ] 46. [test] 编写单元测试

#### API 10: 通过mediaid获取数据
- [x] 47. [request] 创建 GetDataByMediaID.go - 请求参数：media_id
- [x] 48. [response] 添加 ResponseGetDataByMediaID - 响应字段：data
- [x] 49. [client] 实现 GetDataByMediaID 方法 - 调用 wxa/api_getdatabymediaid
- [ ] 50. [test] 编写单元测试

#### Provider 注册
- [x] 51. [src/store/application.go] 注册 Base 到 Store 应用

### Phase 4: 验证
- [ ] 52. go build ./... - 编译验证
- [ ] 53. go vet ./... - 代码检查
- [ ] 54. go test ./src/store/... - 运行单元测试
