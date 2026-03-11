# store/vip 小店会员模块开发任务

## 模块概述
- **模块名称**: vip (小店会员)
- **API总数**: 4个
- **API路径前缀**: `channels/ec/vip/`

---

## 准备工作
- [ ] 1. [src/store/vip/doc.go] 包文档注释
- [ ] 2. [src/store/vip/const.go] 4个API URL常量定义
- [ ] 3. [src/store/vip/request.go] 请求结构体
- [ ] 4. [src/store/vip/response.go] 响应结构体

---

## API 1: 获取用户积分
- [ ] 5. [request] RequestGetVIPUserScore - 请求参数：openid
- [ ] 6. [response] ResponseGetVIPUserScore - 响应字段：score, total_score
- [ ] 7. [client] GetVIPUserScore 方法 - 调用 channels/ec/vip/score/get
- [ ] 8. [test] 编写单元测试

## API 2: 获取用户信息
- [ ] 9. [request] RequestGetUserInfo - 请求参数：openid
- [ ] 10. [response] ResponseGetUserInfo - 响应字段：user_info
- [ ] 11. [client] GetUserInfo 方法 - 调用 channels/ec/vip/user/info/get
- [ ] 12. [test] 编写单元测试

## API 3: 获取用户列表
- [ ] 13. [request] RequestGetUserList - 请求参数：page_size, next_key
- [ ] 14. [response] ResponseGetUserList - 响应字段：user_list, next_key, has_more
- [ ] 15. [client] GetUserList 方法 - 调用 channels/ec/vip/user/list/get
- [ ] 16. [test] 编写单元测试

## API 4: 获取用户积分流水
- [ ] 17. [request] RequestGetUserScoreFlowRecord - 请求参数：openid, page_size, next_key
- [ ] 18. [response] ResponseGetUserScoreFlowRecord - 响应字段：flow_list, next_key, has_more
- [ ] 19. [client] GetUserScoreFlowRecord 方法 - 调用 channels/ec/vip/score/flow/get
- [ ] 20. [test] 编写单元测试

---

## Provider 注册
- [ ] 21. [src/store/vip/provider.go] IOC 注册
- [ ] 22. [src/store/application.go] 注册 VIP 到 Store 应用

---

## 验证
- [ ] 23. go build ./src/store/... - 编译验证
- [ ] 24. go vet ./src/store/... - 代码检查
- [ ] 25. go test ./src/store/vip/... - 运行单元测试
