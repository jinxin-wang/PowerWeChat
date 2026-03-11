# store/league 优选联盟模块开发任务

## 模块概述
- **模块名称**: league (优选联盟)
- **API总数**: 11个
- **子模块**: 达人操作(5个) + 商品操作(6个)
- **API路径前缀**: `channels/ec/league/`

---

## 准备工作
- [ ] 1. [src/store/league/doc.go] 包文档注释
- [ ] 2. [src/store/league/const.go] 11个API URL常量定义
- [ ] 3. [src/store/league/request.go] 请求结构体
- [ ] 4. [src/store/league/response.go] 响应结构体

---

## 达人操作 (5个API)

### API 1: 新增达人
- [ ] 5. [request] RequestAddPromoter - 请求参数：finder_id, promoter_type
- [ ] 6. [response] ResponseAddPromoter - 响应字段：promoter_id
- [ ] 7. [client] AddPromoter 方法 - 调用 channels/ec/league/promoter/add
- [ ] 8. [test] 编写单元测试

### API 2: 删除达人
- [ ] 9. [request] RequestDeletePromoter - 请求参数：promoter_id
- [ ] 10. [response] ResponseDeletePromoter - 基础响应
- [ ] 11. [client] DeletePromoter 方法 - 调用 channels/ec/league/promoter/delete
- [ ] 12. [test] 编写单元测试

### API 3: 获取达人详情信息
- [ ] 13. [request] RequestGetPromoter - 请求参数：promoter_id
- [ ] 14. [response] ResponseGetPromoter - 响应字段：promoter_info
- [ ] 15. [client] GetPromoter 方法 - 调用 channels/ec/league/promoter/get
- [ ] 16. [test] 编写单元测试

### API 4: 获取商店达人列表
- [ ] 17. [request] RequestGetPromoterList - 请求参数：page_size, next_key
- [ ] 18. [response] ResponseGetPromoterList - 响应字段：promoter_list, next_key, has_more
- [ ] 19. [client] GetPromoterList 方法 - 调用 channels/ec/league/promoter/list/get
- [ ] 20. [test] 编写单元测试

### API 5: 编辑达人
- [ ] 21. [request] RequestUpdatePromoter - 请求参数：promoter_id, promoter_type
- [ ] 22. [response] ResponseUpdatePromoter - 基础响应
- [ ] 23. [client] UpdatePromoter 方法 - 调用 channels/ec/league/promoter/update
- [ ] 24. [test] 编写单元测试

---

## 商品操作 (6个API)

### API 6: 批量新增联盟商品
- [ ] 25. [request] RequestBatchAddItem - 请求参数：product_ids
- [ ] 26. [response] ResponseBatchAddItem - 响应字段：result_list
- [ ] 27. [client] BatchAddItem 方法 - 调用 channels/ec/league/item/batchadd
- [ ] 28. [test] 编写单元测试

### API 7: 删除联盟商品
- [ ] 29. [request] RequestDeleteItem - 请求参数：product_id
- [ ] 30. [response] ResponseDeleteItem - 基础响应
- [ ] 31. [client] DeleteItem 方法 - 调用 channels/ec/league/item/delete
- [ ] 32. [test] 编写单元测试

### API 8: 获取联盟商品详情
- [ ] 33. [request] RequestGetItem - 请求参数：product_id
- [ ] 34. [response] ResponseGetItem - 响应字段：item_info
- [ ] 35. [client] GetItem 方法 - 调用 channels/ec/league/item/get
- [ ] 36. [test] 编写单元测试

### API 9: 批量新增联盟机构推广
- [ ] 37. [request] RequestBatchAddHeadSupplierItem - 请求参数：product_ids, supplier_info
- [ ] 38. [response] ResponseBatchAddHeadSupplierItem - 响应字段：result_list
- [ ] 39. [client] BatchAddHeadSupplierItem 方法 - 调用 channels/ec/league/item/batchaddheadsupplier
- [ ] 40. [test] 编写单元测试

### API 10: 获取联盟商品推广列表
- [ ] 41. [request] RequestGetItemList - 请求参数：page_size, next_key, status
- [ ] 42. [response] ResponseGetItemList - 响应字段：item_list, next_key, has_more
- [ ] 43. [client] GetItemList 方法 - 调用 channels/ec/league/item/list/get
- [ ] 44. [test] 编写单元测试

### API 11: 更新联盟商品信息
- [ ] 45. [request] RequestUpdateItem - 请求参数：product_id, item_info
- [ ] 46. [response] ResponseUpdateItem - 基础响应
- [ ] 47. [client] UpdateItem 方法 - 调用 channels/ec/league/item/update
- [ ] 48. [test] 编写单元测试

---

## Provider 注册
- [ ] 49. [src/store/league/provider.go] IOC 注册
- [ ] 50. [src/store/application.go] 注册 League 到 Store 应用

---

## 验证
- [ ] 51. go build ./src/store/... - 编译验证
- [ ] 52. go vet ./src/store/... - 代码检查
- [ ] 53. go test ./src/store/league/... - 运行单元测试
