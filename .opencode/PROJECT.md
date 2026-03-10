# 项目目标
提供一套 Go 语言标准化、高复用、易扩展的微信生态 SDK，让业务开发者只关注业务，不处理签名/加密/Token 等底层。

# 整体架构
- 内核层(kernel)：HTTP、Token、加解密、日志、错误
- 基础服务层(basicService)：跨业务公共能力
- 业务模块 / payment / store层：miniProgram 等

# store(微信小店)设计
- 独立模块，不嵌套进小程序
- 完全复用内核能力，不重复造轮子
- 子模块拆分：通用接口，店铺管理，订单管理，售后管理，商家客服，纠纷管理，物流发货，优选联盟，罗盘商家，小店会员
- 子模块列表和接口文档在[DOCUMENTS.md](.opencode/DOCUMENTS.md)中

# store 模块入口

## 应用初始化
```go
func NewStore(config *UserConfig) (*Store, error)
```

## 配置结构
```go
type UserConfig struct {
    AppID  string
    Secret string
    StableTokenMode   bool
    ForceRefresh      bool
    RefreshToken      string
    ComponentAppID    string
    ComponentAppToken string
    Token             string
    AESKey            string
    ResponseType       string
    Log               Log
    OAuth             OAuth
    Cache             kernel.CacheInterface
    Http              Http
    HttpDebug         bool
    Debug             bool
    NotifyURL         string
}

type Http struct {
    Timeout   float64
    BaseURI   string
    ProxyURI  string
    Transport http.RoundTripper
}

type Log struct {
    Driver contract.LoggerInterface
    Level  string
    File   string
    Error  string
    ENV    string
    Stdout bool
}

type OAuth struct {
    Callback string
    Scopes   []string
}
```

## 子模块
- **Base**: 通用接口（10个API）
  - 获取稳定版接口调用凭据
  - 查询API调用额度
  - 重置指定API调用次数
  - 重置API调用次数
  - 使用AppSecret重置API调用次数
  - 网络通信检测
  - 获取微信API服务器IP
  - 获取微信推送服务器IP
  - 查询rid信息
  - 通过mediaid获取数据
- **Store**: 店铺管理（待开发）
- **Order**: 订单管理（待开发）
- **AfterSale**: 售后管理（待开发）
- **CustomerService**: 商家客服（待开发）
- **Complaint**: 纠纷管理（待开发）
- **Delivery**: 物流发货（待开发）
- **Alliance**: 优选联盟（待开发）
- **Compass**: 罗盘商家（待开发）
- **Member**: 小店会员（待开发）

## 使用示例
```go
app, err := store.NewStore(&store.UserConfig{
    AppID:  "your_app_id",
    Secret: "your_secret",
    StableTokenMode: true,
})
if err != nil {
    panic(err)
}

// 调用通用接口
resp, err := app.Base.GetStableAccessToken(ctx, &request.GetStableAccessToken{
    GrantType: "client_credential",
    AppID:     "your_app_id",
    Secret:    "your_secret",
})
```
