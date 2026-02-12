# Hoppscotch 转 Markdown 生成器

一个用于将 Hoppscotch JSON 集合转换为 Markdown API 文档的 Go CLI 工具。

## 功能特性

- 解析 Hoppscotch JSON 集合文件
- 生成格式良好的 Markdown 文档
- 支持多级文件夹结构
- 执行 GET 请求并将实际响应包含在文档中
- 支持：
  - 目录索引
  - HTTP 方法（带视觉徽章：🟢 GET、🟡 POST、🔴 DELETE）
  - 请求头
  - 查询参数
  - 请求体（格式化的 JSON）
  - 响应数据（状态码、响应头、响应体）
  - 认证信息
  - 完整的描述支持

## 安装

### 使用 Make（推荐）

```bash
# 克隆仓库
git clone https://github.com/dovics/hoppscotch-gen-doc.git
cd hoppscotch-gen-doc

# 使用 make 构建
make build

# 或直接安装到 GOPATH/bin
make install
```

### 从源码安装

```bash
# 构建
go build -o hoppscotch-gen-doc

# 安装到 GOPATH/bin
go install
```

### 使用 Go install

```bash
go install github.com/dovics/hoppscotch-gen-doc@latest
```

## 开发

### Make 命令

```bash
# 显示所有可用命令
make help

# 构建应用
make build

# 安装到 GOPATH/bin
make install

# 清理构建文件
make clean

# 运行测试
make test

# 运行测试并生成覆盖率报告
make test-coverage

# 格式化代码
make fmt

# 整理 go modules
make tidy

# 运行代码检查（需要安装 golangci-lint）
make lint

# 构建多平台发布版本
make release

# 生成示例文档
make generate
```

## 使用方法

### 查看帮助

```bash
# 查看主命令帮助
hoppscotch-gen-doc --help

# 查看子命令帮助
hoppscotch-gen-doc generate --help
```

### 生成文档

```bash
# 生成到文件（不执行请求）
hoppscotch-gen-doc generate -i example.json -o API.md

# 文档替换模式：替换文档中显示的服务器 URL
hoppscotch-gen-doc generate -i example.json --server https://api.example.com -o API.md

# 请求替换模式：只在执行请求时替换服务器 URL，文档中仍显示原始 URL
hoppscotch-gen-doc generate -i example.json --target-server https://api.example.com -x -o API.md

# 生成到文件并执行 GET 请求
hoppscotch-gen-doc generate -i example.json -o API.md -x

# 替换服务器地址并执行 GET 请求
hoppscotch-gen-doc generate -i example.json --server https://api.example.com -x -o API.md

# 输出到终端
hoppscotch-gen-doc generate -i example.json

# 使用自定义超时时间执行 GET 请求（默认：10 秒）
hoppscotch-gen-doc generate -i example.json -x -t 30 -o API.md

# 使用完整参数名
hoppscotch-gen-doc generate --input example.json --output API.md --execute

# 使用 make
make generate
```

### 命令行参数

| 参数 | 简写 | 描述 | 必需 |
|------|------|------|------|
| `--input` | `-i` | Hoppscotch JSON 文件路径 | 是 |
| `--output` | `-o` | 输出 Markdown 文件路径（可选，默认输出到 stdout） | 否 |
| `--server` | | 只在文档中替换服务器 URL（请求仍发送到原始 URL） | 否 |
| `--target-server` | | 只在执行请求时替换服务器 URL（文档中显示原始 URL） | 否 |
| `--execute` | `-x` | 执行 GET 请求并将响应包含在文档中 | 否 |
| `--timeout` | `-t` | 请求超时时间（秒，默认 10） | 否 |
| `--help` | `-h` | 显示帮助信息 | 否 |

## 使用示例

### 基础文档生成

给定一个 Hoppscotch JSON 文件 `example.json`，运行：

```bash
hoppscotch-gen-doc generate -i example.json -o API.md
```

将生成包含以下内容的 Markdown 文件：

- API 集合名称作为标题
- 分层目录（按文件夹组织）
- 文件夹分组及描述
- 每个请求的详细文档，包括：
  - HTTP 方法及视觉徽章
  - 端点 URL
  - 描述信息
  - 请求头表格
  - 查询参数表格
  - 请求体（格式化的 JSON）
  - 认证详情

### 替换服务器地址

工具支持两种服务器 URL 替换模式：

#### 模式 1：文档替换（`--server`）

只在文档中替换服务器 URL，请求仍发送到原始 URL：

```bash
hoppscotch-gen-doc generate -i example.json --server https://api.example.com -o API.md
```

**示例：**

如果你的原始端点是：
- `http://localhost:8080/api/v1/health`

使用 `--server https://api.example.com` 后：
- 文档中显示：`https://api.example.com/api/v1/health`
- 请求发送到：`http://localhost:8080/api/v1/health`（原始 URL）

这在以下场景中非常有用：
- Hoppscotch 集合使用的是开发服务器 URL
- 你想在文档中显示生产服务器 URL
- 但请求仍然发送到开发服务器
- 或结合 `--execute` 使用时，可以对不同服务器执行请求

#### 模式 2：请求替换（`--target-server`）

只在执行请求时替换服务器 URL，文档中保持原始 URL：

```bash
hoppscotch-gen-doc generate -i example.json --target-server https://api.example.com -x -o API.md
```

**示例：**

如果你的原始端点是：
- `http://localhost:8080/api/v1/health`

使用 `--target-server https://api.example.com` 后：
- 文档中显示：`http://localhost:8080/api/v1/health`（原始 URL）
- 请求发送到：`https://api.example.com/api/v1/health`（替换后的 URL）

这在以下场景中非常有用：
- Hoppscotch 集合使用的是开发服务器 URL
- 你想在文档中保留原始 URL
- 但需要针对不同的服务器（如生产服务器）执行请求并获取响应
- 你需要针对不同环境测试 API，而不改变文档内容

### 执行 GET 请求

要在文档中包含实际的 API 响应，使用 `--execute` 标志：

```bash
hoppscotch-gen-doc generate -i example.json -x -o API.md
```

这将执行所有 GET 请求并包含：

- **响应状态码**：HTTP 状态码和消息
- **响应头**：响应头表格
- **响应体**：格式化的 JSON 或文本响应

示例输出：

```markdown
### Health

**🟢 GET**

**Endpoint:** `https://api.example.com/health`

#### Response

**Status Code:** 200 200 OK

**Response Headers:**

| Key | Value |
|-----|-------|
| Content-Type | application/json |
| Server | nginx |

**Response Body:**

```json
{
  "status": "healthy"
}
```
```

## 项目结构

```
.
├── cmd/
│   ├── root.go         # 根命令
│   └── generate.go     # generate 子命令
├── internal/
│   └── generator/
│       └── generator.go # 文档生成逻辑
├── main.go             # 入口文件
├── Makefile            # 构建自动化
├── go.mod
├── go.sum
├── README.md
├── README_zh.md
├── example.json        # 示例输入文件
└── .gitignore
```

## 完整示例

### 输入 JSON (example.json)

```json
{
  "v": 11,
  "name": "Paas",
  "folders": [
    {
      "v": 11,
      "name": "Postgres",
      "folders": [],
      "requests": [
        {
          "v": "17",
          "name": "Create PostgreSQL Clusters",
          "method": "POST",
          "endpoint": "https://operator.insightst.com/api/v1/clusters",
          "body": {
            "contentType": "application/json",
            "body": "{\"database\": \"postgresql\", \"name\": \"my-postgres\"}"
          }
        }
      ]
    }
  ],
  "requests": []
}
```

### 输出 Markdown (API.md)

```markdown
# Paas

## Table of Contents

- [Postgres](#postgres)
  - [Create PostgreSQL Clusters](#create-postgresql-clusters)

## Postgres

### Create PostgreSQL Clusters

**🟡 POST**

**Endpoint:** `https://operator.insightst.com/api/v1/clusters`

#### Request Body

**Content-Type:** application/json

```json
{
  "database": "postgresql",
  "name": "my-postgres"
}
```
```

## 贡献

欢迎提交 Issue 和 Pull Request！

## 许可证

MIT License

## 使用的技术

- [Cobra](https://github.com/spf13/cobra) - 强大的 Go CLI 应用程序框架
- Go 标准库 - encoding/json, fmt, strings, os, net/http
