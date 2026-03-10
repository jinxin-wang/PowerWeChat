# IDENTITY.md

### 项目名称
- **PowerWeChat框架二次开发**: https://github.com/jinxin-wang/PowerWeChat
- 本项目基于 Golang 的微信全生态 SDK，从原PowerWeChat开源项目fork而来，并基于原项目PowerWeChat加入微信小店store的模块
- 整体结构：kernel（内核）→ basicService（基础服务）→ 业务模块（miniProgram / officialAccount / payment / store 等）。
- 模块约定：按微信官方业务域划分，store（微信小店）为独立模块。

### 技术要求
Golang、DI 容器、强类型结构体、HTTP 客户端封装、自动 Token 管理
- **语言**: Go 1.21+
- **WeChat SDK for Go, 微信SDK**: https://github.com/silenceper/wechat/
- **PowerWeChat框架**: https://github.com/ArtisanCloud/PowerWeChat  
封装公众号/小程序/微信支付/企业微信等能力
- **包管理**: Go modules

### 开发规范
- 使用 Go 官方推荐的代码风格 (gofmt)
- 优先使用标准库，减少外部依赖
- 错误处理：使用明确的错误类型和错误包装
- 命名：遵循 Go 惯例（camelCase 变量/函数，PascalCase 类型/接口）
- 并发：使用 goroutine 和 channel 时确保资源安全

### 项目结构

src/
  ├── basicSerice/   // 基础服务层，封装微信开放平台的跨业务通用基础能力，为其他模块提供公共服务
  ├── channels/      // 视频号，封装微信视频号（Channels）电商场景
  ├── kernel/        // 核心基础层，封装通用底层能力
  ├── miniProgram/      // 小程序
  ├── officialAccount/  // 公众号
  ├── openPlatform/      // 微信开放平台模块，封装微信开放平台（第三方服务商模式）能力
  ├── openWork/         // 企业微信
  ├── payment/          // 微信支付
  ├── work/         // 企业微信基础模块 / 历史模块，企业微信模块的早期封装或基础依赖
  └── store/            // 需要添加的微信小店模块

#### 模块目录结构解析

src/miniProgram/base/
  ├── request/
  ├── response/
  ├── client.go
  └── provider.go

示例代码地址 https://github.com/jinxin-wang/PowerWeChat/tree/release/3.4.0/src/miniProgram/base

| 名称| 核心作用|
|---|---|
| request/ | 存放小程序基础接口的**请求参数结构体**，所有基础接口入参封装为强类型Go结构体，编译期校验参数合法性，统一入参规范。 |
| response/ | 存放小程序基础接口的**响应数据结构体**，包含通用响应（errcode/errmsg）和专属响应模型，标准化JSON响应解析逻辑。|
| client.go | 小程序基础API客户端核心实现，封装HTTP请求底层逻辑，提供统一的API请求发送入口。 |
| provider.go | 服务提供器/依赖注入文件，负责初始化模块客户端实例，将能力注册到应用中，实现模块依赖管理与能力复用。 |

#### 核心设计逻辑
1. 遵循“请求/响应分离”原则：将入参/出参结构体拆分到独立目录，符合PowerWeChat强类型、高可维护性的设计规范；
2. 客户端与服务提供器解耦：client.go聚焦请求逻辑，provider.go聚焦实例注册，便于上层模块复用的基础调用能力；

### 常用命令
- `go build` - 编译项目
- `go test` - 运行测试
- `go vet` - 代码检查
- `golangci-lint run` - 代码 linting

## 服务器地址
共网IPv4: 43.139.131.42
内网IPv4: 10.1.20.13
IPv6: 2402:4e00:c011:3900:7ae6:d37a:ea98:0

