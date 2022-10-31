<h1 align="center">backend</h1>
<p align="center">简历在线生成、在线预览服务</p>

## 命令行工具

### 新建 Model

```bash
# example
# go run main.go make model resume
go run main.go make model [modelName]
```

### 新建数据库迁移

```bash
# example
# go run main.go make migration add_resumes_table
go run main.go make migration [migrationFileName]

# 修改 database/migrations/xxxx_xx_xx_xxxxxx_add_resumes_table.go
# 修改 Struct 后执行迁移
go run main.go migrate up

# 发现写的不对？可以执行回滚
go run main.go migrate down
```

### 生成请求验证器

```bash
# example
# go run main.go make request resume
go run main.go make request [fileName]
```

### 生成 API controller

```bash
# example
# go run main.go make api-controller v1/resume
go run main.go make api-controller [apiControllerName]
```
