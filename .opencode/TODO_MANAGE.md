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

