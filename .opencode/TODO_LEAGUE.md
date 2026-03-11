# store/league 优选联盟模块开发任务

## 模块概述
- **模块名称**: league (优选联盟)
- **API总数**: 11个
- **子模块**: 达人操作(5个) + 商品操作(6个)
- **API路径前缀**: `channels/ec/league/`

---

## 准备工作
- [x] 1. [src/store/league/doc.go] 包文档注释
- [x] 2. [src/store/league/const.go] 11个API URL常量定义
- [x] 3. [src/store/league/request.go] 请求结构体
- [x] 4. [src/store/league/response.go] 响应结构体

---

## 达人操作 (5个API)

### API 1: 新增达人
- [x] 5. [request] RequestAddPromoter - 请求参数：finder_id, promoter_type
- [x] 6. [response] ResponseAddPromoter - 响应字段：promoter_id
- [x] 7. [client] AddPromoter 方法 - 调用 channels/ec/league/promoter/add
- [x] 8. [test] 编写单元测试

### API 2: 删除达人
- [x] 9. [request] RequestDeletePromoter - 请求参数：promoter_id
- [x] 10. [response] ResponseDeletePromoter - 基础响应
- [x] 11. [client] DeletePromoter 方法 - 调用 channels/ec/league/promoter/delete
- [x] 12. [test] 编写单元测试

### API 3: 获取达人详情信息
- [x] 13. [request] RequestGetPromoter - 请求参数：promoter_id
- [x] 14. [response] ResponseGetPromoter - 响应字段：promoter_info
- [x] 15. [client] GetPromoter 方法 - 调用 channels/ec/league/promoter/get
- [x] 16. [test] 编写单元测试

### API 4: 获取商店达人列表
- [x] 17. [request] RequestGetPromoterList - 请求参数：page_size, next_key
- [x] 18. [response] ResponseGetPromoterList - 响应字段：promoter_list, next_key, has_more
- [x] 19. [client] GetPromoterList 方法 - 调用 channels/ec/league/promoter/list/get
- [x] 20. [test] 编写单元测试

### API 5: 编辑达人
- [x] 21. [request] RequestUpdatePromoter - 请求参数：promoter_id, promoter_type
- [x] 22. [response] ResponseUpdatePromoter - 基础响应
- [x] 23. [client] UpdatePromoter 方法 - 调用 channels/ec/league/promoter/update
- [x] 24. [test] 编写单元测试

---

## 商品操作 (6个API)

### API 6: 批量新增联盟商品
- [x] 25. [request] RequestBatchAddItem - 请求参数：product_ids
- [x] 26. [response] ResponseBatchAddItem - 响应字段：result_list
- [x] 27. [client] BatchAddItem 方法 - 调用 channels/ec/league/item/batchadd
- [x] 28. [test] 编写单元测试

### API 7: 删除联盟商品
- [x] 29. [request] RequestDeleteItem - 请求参数：product_id
- [x] 30. [response] ResponseDeleteItem - 基础响应
- [x] 31. [client] DeleteItem 方法 - 调用 channels/ec/league/item/delete
- [x] 32. [test] 编写单元测试

### API 8: 获取联盟商品详情
- [x] 33. [request] RequestGetItem - 请求参数：product_id
- [x] 34. [response] ResponseGetItem - 响应字段：item_info
- [x] 35. [client] GetItem 方法 - 调用 channels/ec/league/item/get
- [x] 36. [test] 编写单元测试

### API 9: 批量新增联盟机构推广
- [x] 37. [request] RequestBatchAddHeadSupplierItem - 请求参数：product_ids, supplier_info
- [x] 38. [response] ResponseBatchAddHeadSupplierItem - 响应字段：result_list
- [x] 39. [client] BatchAddHeadSupplierItem 方法 - 调用 channels/ec/league/item/batchaddheadsupplier
- [x] 40. [test] 编写单元测试

### API 10: 获取联盟商品推广列表
- [x] 41. [request] RequestGetItemList - 请求参数：page_size, next_key, status
- [x] 42. [response] ResponseGetItemList - 响应字段：item_list, next_key, has_more
- [x] 43. [client] GetItemList 方法 - 调用 channels/ec/league/item/list/get
- [x] 44. [test] 编写单元测试

### API 11: 更新联盟商品信息
- [x] 45. [request] RequestUpdateItem - 请求参数：product_id, item_info
- [x] 46. [response] ResponseUpdateItem - 基础响应
- [x] 47. [client] UpdateItem 方法 - 调用 channels/ec/league/item/update
- [x] 48. [test] 编写单元测试

---

## Provider 注册
- [x] 49. [src/store/league/provider.go] IOC 注册
- [x] 50. [src/store/application.go] 注册 League 到 Store 应用

---

## 验证
- [x] 51. go build ./src/store/... - 编译验证
- [x] 52. go vet ./src/store/... - 代码检查
- [x] 53. go test ./src/store/league/... - 运行单元测试

---

## Bug修复
- [x] 54. 修复 response/responseItem.go 缺少 response 包导入
- [x] 55. 修复 client.go 未使用的 respKernel 导入
