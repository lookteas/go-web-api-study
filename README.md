# go-web-api-study

📚 Go 语言 Web 开发与接口开发学习笔记与实践项目集合

> 记录使用 Go (Golang) 进行 Web 开发和 API 接口开发过程中的学习笔记、代码示例、实战项目与最佳实践。

---

## 🌱 项目目标

本仓库旨在系统性地学习和掌握 Go 语言在后端 Web 开发和 RESTful API 开发中的应用，涵盖从基础语法到框架使用、项目结构设计、数据库操作、中间件开发、接口测试等完整技能链。

适合：
- Go 初学者入门 Web 开发
- 想深入理解 Go 后端工程实践的开发者
- 寻找可复用代码示例的学习者

---

## 📚 学习内容概览

| 主题 | 内容 |
|------|------|
| ✅ 基础语法回顾 | 变量、函数、结构体、接口、并发等 |
| ✅ HTTP 服务开发 | 使用 `net/http` 构建基础 Web 服务 |
| ✅ 路由控制 | 使用 `gorilla/mux` 或 `gin` 实现路由 |
| ✅ RESTful API 设计 | 请求处理、JSON 序列化、状态码规范 |
| ✅ 中间件开发 | 日志、认证、CORS、限流等 |
| ✅ 数据库操作 | 使用 `database/sql` 与 `GORM` 操作 PostgreSQL/MySQL |
| ✅ 用户认证 | JWT、Session、OAuth2 基础实现 |
| ✅ 配置管理 | 使用 `Viper` 管理配置文件 |
| ✅ 错误处理 | 统一错误响应结构 |
| ✅ 接口测试 | 使用 `testing` 包进行单元测试与集成测试 |
| ✅ 项目结构设计 | 推荐的 Go 项目分层结构（如 api, service, model, handler） |
| ✅ 部署实践 | Docker 打包、简单部署流程 |

---

## 📁 项目结构示例（建议）

```bash
.
├── cmd/                   # 主程序入口
│   └── api/
│       └── main.go
├── internal/              # 内部逻辑代码
│   ├── handler/           # HTTP 请求处理器
│   ├── service/           # 业务逻辑
│   ├── model/             # 数据结构与数据库模型
│   ├── middleware/        # 自定义中间件
│   └── config/            # 配置加载
├── pkg/                   # 可复用工具包
├── config/                # 配置文件（如 config.yaml）
├── api/                   # OpenAPI/Swagger 文档（可选）
├── scripts/               # 脚本（如数据库迁移）
├── tests/                 # 测试代码
├── go.mod
├── go.sum
└── README.md


# 克隆项目
git clone https://github.com/你的用户名/go-web-api-study.git
cd go-web-api-study
