# store/compass 罗盘商家版模块开发任务

## 模块概述
- **模块名称**: compass (罗盘商家版)
- **API总数**: 11个
- **API路径前缀**: `channels/ec/compass/`

---

## 准备工作
- [ ] 1. [src/store/compass/doc.go] 包文档注释
- [ ] 2. [src/store/compass/const.go] 11个API URL常量定义
- [ ] 3. [src/store/compass/request.go] 请求结构体
- [ ] 4. [src/store/compass/response.go] 响应结构体

---

## API 1: 获取授权视频号列表
- [ ] 5. [request] RequestGetShopFinderAuthorizationList - 请求参数：page_size, next_key
- [ ] 6. [response] ResponseGetShopFinderAuthorizationList - 响应字段：finder_list, next_key, has_more
- [ ] 7. [client] GetShopFinderAuthorizationList 方法 - 调用 channels/ec/compass/shop/finder/authorization/list
- [ ] 8. [test] 编写单元测试

## API 2: 获取带货达人列表
- [ ] 9. [request] RequestGetShopFinderList - 请求参数：page_size, next_key, start_time, end_time
- [ ] 10. [response] ResponseGetShopFinderList - 响应字段：finder_list, next_key, has_more
- [ ] 11. [client] GetShopFinderList 方法 - 调用 channels/ec/compass/shop/finder/list
- [ ] 12. [test] 编写单元测试

## API 3: 获取带货数据概览
- [ ] 13. [request] RequestGetShopFinderOverall - 请求参数：finder_id, start_time, end_time
- [ ] 14. [response] ResponseGetShopFinderOverall - 响应字段：overall_data
- [ ] 15. [client] GetShopFinderOverall 方法 - 调用 channels/ec/compass/shop/finder/overall
- [ ] 16. [test] 编写单元测试

## API 4: 获取带货达人商品列表
- [ ] 17. [request] RequestGetShopFinderProductList - 请求参数：finder_id, page_size, next_key
- [ ] 18. [response] ResponseGetShopFinderProductList - 响应字段：product_list, next_key, has_more
- [ ] 19. [client] GetShopFinderProductList 方法 - 调用 channels/ec/compass/shop/finder/product/list
- [ ] 20. [test] 编写单元测试

## API 5: 获取带货达人详情
- [ ] 21. [request] RequestGetShopFinderProductOverall - 请求参数：finder_id, product_id, start_time, end_time
- [ ] 22. [response] ResponseGetShopFinderProductOverall - 响应字段：product_overall
- [ ] 23. [client] GetShopFinderProductOverall 方法 - 调用 channels/ec/compass/shop/finder/product/overall
- [ ] 24. [test] 编写单元测试

## API 6: 获取店铺开播列表
- [ ] 25. [request] RequestGetShopLiveList - 请求参数：page_size, next_key, start_time, end_time
- [ ] 26. [response] ResponseGetShopLiveList - 响应字段：live_list, next_key, has_more
- [ ] 27. [client] GetShopLiveList 方法 - 调用 channels/ec/compass/shop/live/list
- [ ] 28. [test] 编写单元测试

## API 7: 获取电商数据概览
- [ ] 29. [request] RequestGetShopOverall - 请求参数：start_time, end_time
- [ ] 30. [response] ResponseGetShopOverall - 响应字段：overall_data
- [ ] 31. [client] GetShopOverall 方法 - 调用 channels/ec/compass/shop/overall
- [ ] 32. [test] 编写单元测试

## API 8: 获取商品详细信息
- [ ] 33. [request] RequestGetShopProductData - 请求参数：product_id, start_time, end_time
- [ ] 34. [response] ResponseGetShopProductData - 响应字段：product_data
- [ ] 35. [client] GetShopProductData 方法 - 调用 channels/ec/compass/shop/product/data
- [ ] 36. [test] 编写单元测试

## API 9: 获取商品列表
- [ ] 37. [request] RequestGetShopProductList - 请求参数：page_size, next_key, start_time, end_time
- [ ] 38. [response] ResponseGetShopProductList - 响应字段：product_list, next_key, has_more
- [ ] 39. [client] GetShopProductList 方法 - 调用 channels/ec/compass/shop/product/list
- [ ] 40. [test] 编写单元测试

## API 10: 获取店铺人群数据
- [ ] 41. [request] RequestGetShopSaleProfileData - 请求参数：start_time, end_time, type
- [ ] 42. [response] ResponseGetShopSaleProfileData - 响应字段：profile_data
- [ ] 43. [client] GetShopSaleProfileData 方法 - 调用 channels/ec/compass/shop/sale/profile
- [ ] 44. [test] 编写单元测试

---

## Provider 注册
- [ ] 45. [src/store/compass/provider.go] IOC 注册
- [ ] 46. [src/store/application.go] 注册 Compass 到 Store 应用

---

## 验证
- [ ] 47. go build ./src/store/... - 编译验证
- [ ] 48. go vet ./src/store/... - 代码检查
- [ ] 49. go test ./src/store/compass/... - 运行单元测试
