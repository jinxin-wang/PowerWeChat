# PowerWeChat Store 模块开发任务总览

微信小店store模块分为多个子模块进行开发，每个子模块的任务列表被记录到一个指定的文件中，命名规则为TODO_[子模块名].md。
计划进行构建的子模块需要加载相应的任务列表文件，没有被计划构建的子模块就不需要加载其任务列表文件。

## 模块列表

| 状态 | 文件 | 模块 | 描述 | API数量 |
|------|------|------|------|---------|
| [x] | [TODO_BASE.md](.opencode/TODO_BASE.md) | base | store 通用模块 - 获取稳定版接口凭据、API调用额度等 | 2个 |
| [x] | [TODO_ORDER.md](.opencode/TODO_ORDER.md) | order | 订单管理模块 - 订单列表、详情、搜索、价格修改等 | 9个 |
| [x] | [TODO_MANAGE.md](.opencode/TODO_MANAGE.md) | manage | 店铺管理模块 - 店铺基本信息、H5链接、二维码、口令等 | 4个 |
| [x] | [TODO_AFTERSALE.md](.opencode/TODO_AFTERSALE.md) | aftersale | 售后管理模块 - 售后单列表、详情、同意/拒绝售后等 | 19个 |
| [x] | [TODO_COMPASS.md](.opencode/TODO_COMPASS.md) | compass | 罗盘商家版模块 - 授权视频号列表、带货达人列表、数据概览等 | 11个 |
| [ ] | [TODO_KF_COMPLAINT.md](.opencode/TODO_KF_COMPLAINT.md) | kf + complaint | 客服与纠纷管理模块 - 客服上传多媒体、纠纷列表/详情/调解等 | 5个 |
| [x] | [TODO_LEAGUE.md](.opencode/TODO_LEAGUE.md) | league | 优选联盟模块 - 达人操作(5个) + 商品操作(6个) | 11个 |
| [x] | [TODO_LOGISTICS.md](.opencode/TODO_LOGISTICS.md) | logistics | 物流发货模块 - 地址管理(5) + 运费模板(4) + 电子面单(15) + 发货(4) + 虚拟号码(3) | 31个 |
| [x] | [TODO_VIP.md](.opencode/TODO_VIP.md) | vip | 小店会员模块 - 用户积分、用户信息、用户列表、积分明细等 | 4个 |

