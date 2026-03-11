# store/kf + store/complaint 客服与纠纷管理模块开发任务

## 模块概述
- **模块名称**: kf (客服) + complaint (纠纷管理)
- **API总数**: 5个 (kf: 2个, complaint: 3个)
- **API路径前缀**: 
  - kf: `channels/ec/kf/`
  - complaint: `channels/ec/complaint/`

---

## 准备工作
- [ ] 1. [src/store/kf/doc.go] 客服模块包文档注释
- [ ] 2. [src/store/kf/const.go] 客服API URL常量定义
- [ ] 3. [src/store/kf/request.go] 客服请求结构体
- [ ] 4. [src/store/kf/response.go] 客服响应结构体
- [ ] 5. [src/store/complaint/doc.go] 纠纷模块包文档注释
- [ ] 6. [src/store/complaint/const.go] 纠纷API URL常量定义
- [ ] 7. [src/store/complaint/request.go] 纠纷请求结构体
- [ ] 8. [src/store/complaint/response.go] 纠纷响应结构体

---

## kf 客服模块

### API 1: 上传多媒体资源
- [ ] 9. [request] RequestCOSUpload - 请求参数：media_type, media_data
- [ ] 10. [response] ResponseCOSUpload - 响应字段：media_id, url
- [ ] 11. [client] COSUpload 方法 - 调用 channels/ec/kf/cosupload
- [ ] 12. [test] 编写单元测试

### API 2: 发送消息
- [ ] 13. [request] RequestSendMsg - 请求参数：openid, msg_type, content
- [ ] 14. [response] ResponseSendMsg - 基础响应
- [ ] 15. [client] SendMsg 方法 - 调用 channels/ec/kf/sendmsg
- [ ] 16. [test] 编写单元测试

---

## complaint 纠纷管理模块

### API 1: 商家补充纠纷单留言
- [ ] 17. [request] RequestAddComplaintMaterial - 请求参数：complaint_id, content
- [ ] 18. [response] ResponseAddComplaintMaterial - 基础响应
- [ ] 19. [client] AddComplaintMaterial 方法 - 调用 channels/ec/complaint/addcomplaintmaterial
- [ ] 20. [test] 编写单元测试

### API 2: 商家举证
- [ ] 21. [request] RequestAddComplaintProof - 请求参数：complaint_id, proof_info
- [ ] 22. [response] ResponseAddComplaintProof - 基础响应
- [ ] 23. [client] AddComplaintProof 方法 - 调用 channels/ec/complaint/addcomplaintproof
- [ ] 24. [test] 编写单元测试

### API 3: 获取纠纷单
- [ ] 25. [request] RequestGetComplaintOrder - 请求参数：complaint_id
- [ ] 26. [response] ResponseGetComplaintOrder - 响应字段：complaint_info
- [ ] 27. [client] GetComplaintOrder 方法 - 调用 channels/ec/complaint/getcomplaintorder
- [ ] 28. [test] 编写单元测试

---

## Provider 注册
- [ ] 29. [src/store/kf/provider.go] kf IOC 注册
- [ ] 30. [src/store/complaint/provider.go] complaint IOC 注册
- [ ] 31. [src/store/application.go] 注册 KF 和 Complaint 到 Store 应用

---

## 验证
- [ ] 32. go build ./src/store/... - 编译验证
- [ ] 33. go vet ./src/store/... - 代码检查
- [ ] 34. go test ./src/store/kf/... - 运行kf单元测试
- [ ] 35. go test ./src/store/complaint/... - 运行complaint单元测试
