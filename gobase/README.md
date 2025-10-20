# Go 语言基础学习模块

这个目录包含了 Go 语言的基础语法学习示例，帮助你系统地掌握 Go 语言的核心概念。

## 📚 学习顺序

建议按照以下顺序学习：

### 1. 变量和类型 (`01_variables_and_types.go`)
- 变量声明的多种方式
- Go 的基本数据类型
- 复合类型（数组、切片、映射）
- 零值概念
- 类型转换
- 常量定义

**运行命令：**
```bash
go run gobase/01_variables_and_types.go
```

### 2. 函数 (`02_functions.go`)
- 函数定义和调用
- 多返回值
- 命名返回值
- 可变参数
- 匿名函数和闭包
- 高阶函数
- 递归函数
- defer 语句

**运行命令：**
```bash
go run gobase/02_functions.go
```

### 3. 结构体和接口 (`03_structs_and_interfaces.go`)
- 结构体定义和使用
- 方法定义（值接收者 vs 指针接收者）
- 结构体嵌入（组合）
- 接口定义和实现
- 接口组合
- 类型断言
- 空接口

**运行命令：**
```bash
go run gobase/03_structs_and_interfaces.go
```

### 4. 并发编程 (`04_concurrency.go`)
- Goroutine 基础
- Channel 通信
- 带缓冲的 Channel
- Channel 方向
- Select 语句
- sync.WaitGroup
- sync.Mutex 互斥锁
- 工作池模式

**运行命令：**
```bash
go run gobase/04_concurrency.go
```

## 🎯 学习建议

1. **逐个运行示例**：每个文件都是独立的可执行程序，建议逐个运行并观察输出。

2. **修改代码实验**：在理解示例的基础上，尝试修改代码参数，观察不同的行为。

3. **编写练习**：为每个概念编写自己的小练习，加深理解。

4. **结合实际项目**：学习完基础概念后，结合 `cmd/api/main.go` 中的 Web 服务器代码，看看这些概念是如何在实际项目中应用的。

## 📖 扩展学习

完成基础学习后，可以继续学习：

- 错误处理模式
- 包管理和模块系统
- 测试编写
- 性能优化
- Web 开发框架（如 Gin、Echo）
- 数据库操作（database/sql、GORM）

## 🔗 相关资源

- [Go 官方文档](https://golang.org/doc/)
- [Go 语言之旅](https://tour.golang.org/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go 代码审查评论](https://github.com/golang/go/wiki/CodeReviewComments)