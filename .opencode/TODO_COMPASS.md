# store/compass 罗盘商家版模块开发任务

## 模块概述
- **模块名称**: compass (罗盘商家版)
- **API总数**: 11个
- **API路径前缀**: `channels/ec/compass/`

---

## 准备工作
- [x] 1. [src/store/compass/doc.go] 包文档注释
- [x] 2. [src/store/compass/const.go] 11个API URL常量定义
- [x] 3. [src/store/compass/request] 请求结构体
- [x] 4. [src/store/compass/response] 响应结构体

---

## API 1: 获取授权视频号列表
- [x] 5. [request] RequestGetShopFinderAuthorizationList - 请求参数：page_size, next_key
- [x] 6. [response] ResponseGetShopFinderAuthorizationList - 响应字段：finder_list, next_key, has_more
- [x] 7. [client] GetShopFinderAuthorizationList 方法 - 调用 channels/ec/compass/shop/finder/authorization/list
- [x] 8. [test] 编写单元测试

## API 2: 获取带货达人列表
- [x] 9. [request] RequestGetShopFinderList - 请求参数：page_size, next_key, start_time, end_time
- [x] 10. [response] ResponseGetShopFinderList - 响应字段：finder_list, next_key, has_more
- [x] 11. [client] GetShopFinderList 方法 - 调用 channels/ec/compass/shop/finder/list
- [x] 12. [test] 编写单元测试

## API 3: 获取带货数据概览
- [x] 13. [request] RequestGetShopFinderOverall - 请求参数：finder_id, start_time, end_time
- [x] 14. [response] ResponseGetShopFinderOverall - 响应字段：overall_data
- [x] 15. [client] GetShopFinderOverall 方法 - 调用 channels/ec/compass/shop/finder/overall
- [x] 16. [test] 编写单元测试

## API 4: 获取带货达人商品列表
- [x] 17. [request] RequestGetShopFinderProductList - 请求参数：finder_id, page_size, next_key
- [x] 18. [response] ResponseGetShopFinderProductList - 响应字段：product_list, next_key, has_more
- [x] 19. [client] GetShopFinderProductList 方法 - 调用 channels/ec/compass/shop/finder/product/list
- [x] 20. [test] 编写单元测试

## API 5: 获取带货达人详情
- [x] 21. [request] RequestGetShopFinderProductOverall - 请求参数：finder_id, product_id, start_time, end_time
- [x] 22. [response] ResponseGetShopFinderProductOverall - 响应字段：product_overall
- [x] 23. [client] GetShopFinderProductOverall 方法 - 调用 channels/ec/compass/shop/finder/product/overall
- [x] 24. [test] 编写单元测试

## API 6: 获取店铺开播列表
- [x] 25. [request] RequestGetShopLiveList - 请求参数：page_size, next_key, start_time, end_time
- [x] 26. [response] ResponseGetShopLiveList - 响应字段：live_list, next_key, has_more
- [x] 27. [client] GetShopLiveList 方法 - 调用 channels/ec/compass/shop/live/list
- [x] 28. [test] 编写单元测试

## API 7: 获取电商数据概览
- [x] 29. [request] RequestGetShopOverall - 请求参数：start_time, end_time
- [x] 30. [response] ResponseGetShopOverall - 响应字段：overall_data
- [x] 31. [client] GetShopOverall 方法 - 调用 channels/ec/compass/shop/overall
- [x] 32. [test] 编写单元测试

## API 8: 获取商品详细信息
- [x] 33. [request] RequestGetShopProductData - 请求参数：product_id, start_time, end_time
- [x] 34. [response] ResponseGetShopProductData - 响应字段：product_data
- [x] 35. [client] GetShopProductData 方法 - 调用 channels/ec/compass/shop/product/data
- [x] 36. [test] 编写单元测试

## API 9: 获取商品列表
- [x] 37. [request] RequestGetShopProductList - 请求参数：page_size, next_key, start_time, end_time
- [x] 38. [response] ResponseGetShopProductList - 响应字段：product_list, next_key, has_more
- [x] 39. [client] GetShopProductList 方法 - 调用 channels/ec/compass/shop/product/list
- [x] 40. [test] 编写单元测试

## API 10: 获取店铺人群数据
- [x] 41. [request] RequestGetShopSaleProfileData - 请求参数：start_time, end_time, type
- [x] 42. [response] ResponseGetShopSaleProfileData - 响应字段：profile_data
- [x] 43. [client] GetShopSaleProfileData 方法 - 调用 channels/ec/compass/shop/sale/profile
- [x] 44. [test] 编写单元测试

---

## Provider 注册
- [x] 45. [src/store/compass/provider.go] IOC 注册
- [x] 46. [src/store/application.go] 注册 Compass 到 Store 应用

---

## 验证
- [x] 47. go build ./src/store/... - 编译验证
- [x] 48. go vet ./src/store/... - 代码检查
- [x] 49. go test ./src/store/compass/... - 运行单元测试 (20个测试用例全部通过)

---

## 备注
- ✅ 所有任务已完成
- 包含10个请求结构体测试和10个响应结构体测试
- 测试覆盖率：请求/响应结构体字段赋值验证
