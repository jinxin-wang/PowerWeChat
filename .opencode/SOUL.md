# SOUL.md

## 核心原则

- 职责单一：内核不碰业务，业务不碰底层
- 依赖向下：业务只依赖 kernel
- 高内聚低耦合：模块/子模块边界清晰
- 代码即文档：命名与注释对齐官方

1. 分层复用：通用逻辑下沉 kernel，业务只封装领域语义
2. 领域驱动：严格对齐微信官方业务域划分
3. 约定优于配置：统一结构降低心智负担
4. 强类型安全：编译期校验，拒绝运行时玄学

## 代码风格哲学

- 所有源文件使用 Go
- 遵循 gofmt 和 go vet
- 在 CI/CD 中启用静态分析（golangci-lint）

## Formatting 格式化规则

- 使用 gofmt 自动格式化（默认使用 tab 缩进）
- 导出函数必须有文档注释
- 保持代码简洁，避免不必要的嵌套
- 优先使用单行 if 简短判断

## Testing 测试理念

- 用例命名：Test接口名_场景
- 测试文件命名为 `*_test.go`
- 测试放在源文件同目录
- 使用描述性测试名称：`func TestFetchItem_Success(t *testing.T)`
- 遵循 AAA 模式：Setup, Exercise, Teardown
- 使用 testify 框架进行断言
- Mock 外部依赖（HTTP 客户端、数据库等）
- 基准测试使用 `func Benchmark*` 命名
- 单元测试覆盖核心接口
- Mock HTTP，不依赖真实接口


## Documentation 文档规范

- 公共 API 使用 GoDoc 注释
- 记录复杂算法和业务逻辑
- 包文档放在 doc.go 中
- 保持注释与代码更改同步
- 示例测试放在 `*_example_test.go` 中
