package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// serveSourceCode 返回一个处理器函数，用于显示源代码
func serveSourceCode(filePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		content, err := os.ReadFile(filePath)
		if err != nil {
			http.Error(w, "文件不存在: "+filePath, http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `
		<html>
		<head>
			<title>%s - 源代码查看</title>
			<style>
				body { font-family: 'Courier New', monospace; margin: 20px; }
				.header { background: #f0f0f0; padding: 10px; margin-bottom: 20px; }
				.code { background: #f8f8f8; padding: 15px; border: 1px solid #ddd; white-space: pre-wrap; }
				.nav { margin-bottom: 20px; }
				.nav a { margin-right: 10px; color: #0066cc; text-decoration: none; }
				.run-btn { background: #4CAF50; color: white; padding: 8px 16px; border: none; cursor: pointer; margin: 10px 0; }
			</style>
		</head>
		<body>
			<div class="nav">
				<a href="/">🏠 首页</a>
				<a href="/exercises">📚 练习列表</a>
				<a href="/gobase">🔧 基础模块</a>
			</div>
			<div class="header">
				<h2>📄 %s</h2>
				<p>💡 命令行运行: <code>go run %s</code></p>
			</div>
			<div class="code">%s</div>
		</body>
		</html>
		`, filepath.Base(filePath), filepath.Base(filePath), filePath, string(content))
	}
}

func main() {
	// 健康检查端点
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"status": "ok", "message": "Go Web API Study Server is running"}`)
	})

	// 欢迎页面
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `
		<html>
		<head>
			<title>Go Web API 学习项目</title>
			<style>
				body { font-family: Arial, sans-serif; margin: 40px; background: #f5f5f5; }
				.container { max-width: 800px; margin: 0 auto; background: white; padding: 30px; border-radius: 10px; box-shadow: 0 2px 10px rgba(0,0,0,0.1); }
				h1 { color: #333; text-align: center; }
				.section { margin: 20px 0; padding: 15px; background: #f9f9f9; border-radius: 5px; }
				.link-grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(200px, 1fr)); gap: 15px; margin: 20px 0; }
				.link-card { background: #007acc; color: white; padding: 15px; text-align: center; border-radius: 5px; text-decoration: none; }
				.link-card:hover { background: #005a9e; }
			</style>
		</head>
		<body>
			<div class="container">
				<h1>🚀 Go Web API 学习项目</h1>
				<div class="section">
					<h2>📚 学习资源</h2>
					<div class="link-grid">
						<a href="/exercises" class="link-card">📝 每日练习</a>
						<a href="/gobase" class="link-card">🔧 Go基础模块</a>
						<a href="/api/hello" class="link-card">🌐 API示例</a>
						<a href="/health" class="link-card">💚 健康检查</a>
					</div>
				</div>
				<div class="section">
					<h3>🎯 学习目标</h3>
					<ul>
						<li>掌握Go语言基础语法</li>
						<li>学习HTTP服务开发</li>
						<li>构建RESTful API</li>
						<li>数据库操作与ORM</li>
						<li>中间件和路由</li>
					</ul>
				</div>
			</div>
		</body>
		</html>
		`)
	})

	// API示例端点
	http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"message": "Hello from Go Web API!", "timestamp": "%s"}`, "2024-01-01T00:00:00Z")
	})

	// 练习目录页面
	http.HandleFunc("/exercises", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `
		<html>
		<head>
			<title>每日练习 - Go学习</title>
			<style>
				body { font-family: Arial, sans-serif; margin: 40px; }
				.nav { margin-bottom: 20px; }
				.nav a { margin-right: 10px; color: #0066cc; text-decoration: none; }
				.exercise-list { display: grid; grid-template-columns: repeat(auto-fit, minmax(250px, 1fr)); gap: 15px; }
				.exercise-card { background: #f0f8ff; padding: 15px; border-radius: 5px; border: 1px solid #ddd; }
				.exercise-card h3 { margin-top: 0; color: #333; }
			</style>
		</head>
		<body>
			<div class="nav">
				<a href="/">🏠 首页</a>
				<a href="/gobase">🔧 基础模块</a>
			</div>
			<h1>📝 每日练习</h1>
			<div class="exercise-list">
				<div class="exercise-card">
					<h3>Day 01 - Hello World</h3>
					<p>基础输出和变量练习</p>
					<a href="/exercises/day01">查看源代码</a>
				</div>
				<div class="exercise-card">
					<h3>Day 02 - 变量练习</h3>
					<p>变量声明和类型转换</p>
					<a href="/exercises/day02">查看源代码</a>
				</div>
			</div>
		</body>
		</html>
		`)
	})

	// Go基础模块页面
	http.HandleFunc("/gobase", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `
		<html>
		<head>
			<title>Go基础学习模块</title>
			<style>
				body { font-family: Arial, sans-serif; margin: 40px; }
				.nav { margin-bottom: 20px; }
				.nav a { margin-right: 10px; color: #0066cc; text-decoration: none; }
				.module-list { display: grid; grid-template-columns: repeat(auto-fit, minmax(300px, 1fr)); gap: 15px; }
				.module-card { background: #fff8dc; padding: 15px; border-radius: 5px; border: 1px solid #ddd; }
				.module-card h3 { margin-top: 0; color: #333; }
				.advanced { background: #f0f8ff; }
				.http { background: #f5f5dc; }
				.database { background: #f0fff0; }
			</style>
		</head>
		<body>
			<div class="nav">
				<a href="/">🏠 首页</a>
				<a href="/exercises">📝 练习</a>
			</div>
			<h1>🔧 Go基础学习模块</h1>
			<div class="module-list">
				<div class="module-card">
					<h3>01 - 变量和类型</h3>
					<p>变量声明、基本类型、零值、类型转换</p>
					<a href="/gobase/01">查看源代码</a>
				</div>
				<div class="module-card">
					<h3>02 - 切片与映射 CRUD</h3>
					<p>切片/映射的创建、读取、更新、删除及清空技巧</p>
					<a href="/gobase/02">查看源代码</a>
				</div>
				<div class="module-card">
					<h3>03 - 函数</h3>
					<p>函数定义、参数、返回值、闭包、defer</p>
					<a href="/gobase/03">查看源代码</a>
				</div>
				<div class="module-card">
					<h3>04 - 结构体和接口</h3>
					<p>结构体、方法、接口、类型断言</p>
					<a href="/gobase/04">查看源代码</a>
				</div>
				<div class="module-card">
					<h3>05 - 并发编程</h3>
					<p>goroutine、channel、select、同步</p>
					<a href="/gobase/05">查看源代码</a>
				</div>
				<div class="module-card http">
					<h3>06 - HTTP基础</h3>
					<p>HTTP服务器、处理器、中间件、客户端</p>
					<a href="/gobase/06">查看源代码</a>
				</div>
				<div class="module-card http">
					<h3>07 - API开发</h3>
					<p>RESTful API、CRUD操作、JSON处理、错误处理</p>
					<a href="/gobase/07">查看源代码</a>
				</div>
				<div class="module-card database">
					<h3>08 - 数据库操作</h3>
					<p>SQL操作、连接池、事务、ORM基础</p>
					<a href="/gobase/08">查看源代码</a>
				</div>
				<div class="module-card advanced">
					<h3>09 - 高级特性</h3>
					<p>并发模式、反射、泛型、性能优化、设计模式</p>
					<a href="/gobase/09">查看源代码</a>
				</div>
			</div>
		</body>
		</html>
		`)
	})

	// 练习文件源代码查看
	http.HandleFunc("/exercises/day01", serveSourceCode("exercises/day01/hello_world.go"))
	http.HandleFunc("/exercises/day02", serveSourceCode("exercises/day02/variables_practice.go"))

	// Go基础模块源代码查看
	http.HandleFunc("/gobase/01", serveSourceCode("gobase/01_variables_and_types.go"))
	http.HandleFunc("/gobase/02", serveSourceCode("gobase/02_slices_maps.go"))
	http.HandleFunc("/gobase/03", serveSourceCode("gobase/02_functions.go"))
	http.HandleFunc("/gobase/04", serveSourceCode("gobase/03_structs_and_interfaces.go"))
	http.HandleFunc("/gobase/05", serveSourceCode("gobase/04_concurrency.go"))
	http.HandleFunc("/gobase/06", serveSourceCode("gobase/05_http_basics.go"))
	http.HandleFunc("/gobase/07", serveSourceCode("gobase/06_api_development.go"))
	http.HandleFunc("/gobase/08", serveSourceCode("gobase/07_database_basics.go"))
	http.HandleFunc("/gobase/09", serveSourceCode("gobase/08_advanced_features.go"))

	fmt.Println("🚀 Go Web API 学习服务器启动成功!")
	fmt.Println("📱 访问地址: http://localhost:8080")
	fmt.Println("📚 练习目录: http://localhost:8080/exercises")
	fmt.Println("🔧 基础模块: http://localhost:8080/gobase")
	fmt.Println("💚 健康检查: http://localhost:8080/health")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
