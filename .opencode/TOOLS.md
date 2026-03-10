# TOOLS.md

## 工具命令

## Build 构建

# TOOLS.md

```bash
go build ./...              # 编译所有模块（核心：验证SDK所有包可编译）
go build -o bin/powerwechat # 编译SDK核心包到指定目录
go build -ldflags "-s -w" ./...   # 编译所有包并移除调试信息（减小SDK体积）
go build -race ./...        # 编译并启用数据竞争检测（SDK并发场景必备）
```

## Lint 代码检查

```bash
go vet ./...                # 运行 go vet 检查
golangci-lint run           # 运行完整静态分析
golangci-lint run --fix     # 自动修复问题
golangci-lint run ./src/store/... # 仅检查指定模块（如store小店模块）
```

## Test 测试

```bash
go test ./...                   # 运行所有模块单元测试
go test -v ./src/store/...      # 详细输出指定模块（如store）的测试结果
go test -run TestCreateProduct  # 运行单个测试用例（精准测试SDK接口）
go test -cover ./src/...        # 查看SDK核心模块测试覆盖率
go test -coverprofile=coverage.out ./... # 生成覆盖率报告（便于核对80%覆盖率要求）
go test -bench ./src/kernel/... # 基准测试内核模块（HTTP/Token性能）
```

## Type Check 类型检查

```bash
go build ./...         # 类型检查是编译的一部分
go vet ./...           # 静态分析检查
```

## Development 开发

```bash

go mod tidy            # 整理依赖（添加缺失/移除未用，SDK依赖管理核心）
go mod vendor          # 生成vendor目录（SDK发布时锁定依赖版本）
go install ./...       # 安装SDK到GOPATH（便于本地项目引用测试）
go doc ./src/store/    # 查看模块文档（验证doc.go注释是否生效）
go fmt ./...           # 统一代码格式（补充：保证SDK代码风格一致）
go install             # 安装到 $GOPATH/bin
```

## opencode 常用命令

/undo         # 撤销上一次 AI 操作带来的所有文件修改、命令执行结果，强依赖 Git 仓库实现回滚
/redo         # 重做已撤销的操作，恢复被回滚的文件修改与命令结果
/open [文件名/路径]    # 模糊搜索并打开指定文件，预览文件内容
/review       # 审查当前未提交的代码变更，支持指定文件审查，校验编码规范、安全漏洞、逻辑合理性
/compact      # 压缩上下文，特别是在完成一个任务模块后，压缩上下文有助于节省token开始
/mcp          # 管理 Model Context Protocol（MCP）的开关与配置，控制外部工具调用权限

git status    # 查询仓库状态
git pull      # 拉取远程分支最新代码并合并到本地
git commit -m "提交说明" # 将暂存区文件提交到版本库
git push      # 推送本地最新提交到远程分支
git add <文件名> # 	将指定文件加入暂存区

