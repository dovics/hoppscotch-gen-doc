# Hoppscotch 转 Markdown 生成器

一个用于将 Hoppscotch JSON 集合转换为 Markdown API 文档的 Go CLI 工具。

## 功能特性

- 解析 Hoppscotch JSON 集合文件
- 生成格式良好的 Markdown 文档
- 支持多级文件夹结构
- 支持：
  - 目录索引
  - HTTP 方法（带视觉徽章：🟢 GET、🟡 POST、🔴 DELETE）
  - 请求头
  - 查询参数
  - 请求体（格式化的 JSON）
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
# 生成到文件
hoppscotch-gen-doc generate -i example.json -o API.md

# 输出到终端
hoppscotch-gen-doc generate -i example.json

# 使用完整参数名
hoppscotch-gen-doc generate --input example.json --output API.md

# 使用 make
make generate
```

### 命令行参数

| 参数 | 简写 | 描述 | 必需 |
|------|------|------|------|
| `--input` | `-i` | Hoppscotch JSON 文件路径 | 是 |
| `--output` | `-o` | 输出 Markdown 文件路径（可选，默认输出到 stdout） | 否 |
| `--help` | `-h` | 显示帮助信息 | 否 |

## 使用示例

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
- Go 标准库 - encoding/json, fmt, strings, os
