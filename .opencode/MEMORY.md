# MEMORY.md

## 目录结构规范
src/{模块}/
├── client.go      # 模块客户端
├── provider.go    # IOC 注册
├── const.go       # URL/错误码常量
├── doc.go         # 包文档
├── {子模块}/
│   ├── client.go
│   ├── request/xxx.go
│   └── response/responseXXX.go

## 代码规范

### Imports 导入规则

- 使用标准化的导入路径
- 导入顺序：标准库 → 第三方库 → 项目内部包
- 分组：空行分隔不同类型的导入
- 匿名导入（仅使用包内的副作用）：使用 `_`
- 点导入（仅用于测试）：使用 `.`

```go
// Good
import (
    "context"
    "encoding/json"
    "fmt"

    "github.com/mark3labs/mcp-go/server"
    "github.com/pkg/errors"

    "open-wx-mcp/internal/client"
    "open-wx-mcp/internal/types"
)

// Avoid
import "fmt"
import "encoding/json"
```

### Types 类型规则

- 函数参数和返回值必须显式声明类型
- 使用结构体定义对象形状，使用接口定义行为
- 避免使用 `interface{}`，使用泛型（Go 1.18+）
- 使用指针传递大型结构体或需要修改的值

```go
// Good
type WxItem struct {
    ID    string  `json:"id"`
    Title string  `json:"title"`
    Price float64 `json:"price"`
}

func FetchItem(ctx context.Context, id string) (*WxItem, error) {
    // ...
}

// Avoid
func FetchItem(id string) interface{} {
    // ...
}
```

### Naming Conventions 命名约定

- **模块名**：全小写，store
- **请求文件**：大驼峰 CreateProduct.go
- **响应文件**：responseCreateProduct.go
- **类型/接口**: PascalCase（如 `WXClient`）
- **函数/变量**: camelCase
- **常量**: PascalCase 或 MixedCase（根据上下文）
- **结构体**：CreateProductRequest / CreateProductResponse
- **接口**: 以 er 结尾表示行为（如 `Reader`, `Writer`）
- **缩写词**: 全部大写或全小写（如 `URL` 或 `url`，不要 `Url`）

```go
// Good
const MaxRetries = 3

type Reader interface {
    Read(p []byte) (n int, err error)
}

var defaultClient *Client

// Avoid
const max_retries = 3
type iReader interface {}
```

### 基于原项目的二次开发
- 必须复用 kernel.BaseClient
- 响应必须嵌入 response.ResponseBase
- 禁止 map，全强类型结构体
- 错误统一包装：errors.NewXXXError
- 字段必须加 json tag 和官方含义注释

### Error Handling 错误处理

- 始终检查错误，不要忽略（使用 `_` 明确忽略）
- 使用错误包装提供上下文（`fmt.Errorf` with `%w`）
- 定义 sentinel errors 用于特定错误情况
- 避免在错误路径中分配内存
- 按 resp.CheckError() 判断
- 错误必须带 errcode / errmsg

```go
// Good
var (
    ErrNotFound   = errors.New("resource not found")
    ErrUnauthorized = errors.New("unauthorized")
)

func FetchItem(ctx context.Context, id string) (*WxItem, error) {
    resp, err := client.Get(ctx, "/item/"+id)
    if err != nil {
        return nil, fmt.Errorf("fetch item %s: %w", id, err)
    }
    // Handle specific error
    if resp.StatusCode == 404 {
        return nil, ErrNotFound
    }
    // ...
}

// Avoid
func FetchItem(id string) (*WxItem, error) {
    resp, _ := client.Get("/item/"+id)  // 忽略错误
    // ...
}
```

## Concurrency 并发规范

- 使用 `context` 传播取消信号
- 使用 `sync.WaitGroup` 等待一组 goroutine
- 使用 `sync.Mutex` 或 `sync.RWMutex` 保护共享资源
- 优先使用 channel 进行通信而非共享内存
- 关闭 channel 时确保只关闭一次

```go
// Good
func ProcessItems(ctx context.Context, items []Item) error {
    resultCh := make(chan Result, len(items))
    
    var wg sync.WaitGroup
    for _, item := range items {
        wg.Add(1)
        go func(i Item) {
            defer wg.Done()
            result, err := process(i)
            if err != nil {
                result = Result{Error: err}
            }
            resultCh <- result
        }(item)
    }
    
    go func() {
        wg.Close()
        close(resultCh)
    }()
    
    for result := range resultCh {
        // Handle result
    }
    return nil
}
```

## 任务执行规范

任务开始前使用命令`git status` 确认仓库状态，确保工作区干净，无未提交的临时修改，避免操作冲突。
出现非预期修改、代码报错、规范不符时，必须立即执行 `/undo` 撤销全部变更，禁止手动叠加修改
完成任务后将代码变更提交至Git仓库
完成任务模块后调用/compact命令来压缩session上下文，从而节省token
