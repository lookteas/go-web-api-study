// HTTP基础 - Go语言Web开发基础
// 学习目标：掌握Go语言HTTP编程基础概念和实践

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

// User 用户结构体
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func main() {
	fmt.Println("=== Go HTTP基础学习 ===")
	
	// 1. 基础HTTP服务器
	fmt.Println("\n1. 基础HTTP服务器")
	demonstrateBasicServer()
	
	// 2. HTTP处理器函数
	fmt.Println("\n2. HTTP处理器函数")
	demonstrateHandlers()
	
	// 3. HTTP方法处理
	fmt.Println("\n3. HTTP方法处理")
	demonstrateHTTPMethods()
	
	// 4. JSON处理
	fmt.Println("\n4. JSON处理")
	demonstrateJSONHandling()
	
	// 5. 中间件概念
	fmt.Println("\n5. 中间件概念")
	demonstrateMiddleware()
	
	// 6. HTTP客户端
	fmt.Println("\n6. HTTP客户端")
	demonstrateHTTPClient()
	
	fmt.Println("\n=== HTTP基础学习完成 ===")
}

// 1. 基础HTTP服务器演示
func demonstrateBasicServer() {
	fmt.Println("基础HTTP服务器创建方法：")
	
	// 创建一个简单的处理器
	handler := func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World! 当前时间: %s", time.Now().Format("2006-01-02 15:04:05"))
	}
	
	fmt.Println("http.HandleFunc(\"/hello\", handler)")
	fmt.Println("log.Fatal(http.ListenAndServe(\":8080\", nil))")
	fmt.Println("访问: http://localhost:8080/hello")
}

// 2. HTTP处理器函数演示
func demonstrateHandlers() {
	fmt.Println("不同类型的处理器函数：")
	
	// 函数类型处理器
	simpleHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "这是一个简单的处理器函数")
	}
	
	// 带参数的处理器
	paramHandler := func(name string) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello, %s!", name)
		}
	}
	
	fmt.Printf("简单处理器: %T\n", simpleHandler)
	fmt.Printf("参数化处理器: %T\n", paramHandler("张三"))
	
	// 演示路径参数解析
	fmt.Println("\n路径参数解析示例：")
	fmt.Println("URL: /user/123")
	fmt.Println("解析: strings.TrimPrefix(r.URL.Path, \"/user/\")")
}

// 3. HTTP方法处理演示
func demonstrateHTTPMethods() {
	fmt.Println("HTTP方法处理：")
	
	methodHandler := func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fmt.Fprintf(w, "处理GET请求")
		case http.MethodPost:
			fmt.Fprintf(w, "处理POST请求")
		case http.MethodPut:
			fmt.Fprintf(w, "处理PUT请求")
		case http.MethodDelete:
			fmt.Fprintf(w, "处理DELETE请求")
		default:
			http.Error(w, "不支持的HTTP方法", http.StatusMethodNotAllowed)
		}
	}
	
	fmt.Printf("方法处理器类型: %T\n", methodHandler)
	fmt.Println("支持的方法: GET, POST, PUT, DELETE")
	
	// 演示请求体读取
	fmt.Println("\n请求体读取示例：")
	fmt.Println("body, err := io.ReadAll(r.Body)")
	fmt.Println("defer r.Body.Close()")
}

// 4. JSON处理演示
func demonstrateJSONHandling() {
	fmt.Println("JSON数据处理：")
	
	// 创建示例用户
	user := User{
		ID:   1,
		Name: "张三",
		Age:  25,
	}
	
	// JSON编码
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Printf("JSON编码错误: %v", err)
		return
	}
	fmt.Printf("JSON编码结果: %s\n", jsonData)
	
	// JSON解码
	var decodedUser User
	err = json.Unmarshal(jsonData, &decodedUser)
	if err != nil {
		log.Printf("JSON解码错误: %v", err)
		return
	}
	fmt.Printf("JSON解码结果: %+v\n", decodedUser)
	
	// 响应结构示例
	response := Response{
		Code:    200,
		Message: "成功",
		Data:    user,
	}
	
	responseJSON, _ := json.Marshal(response)
	fmt.Printf("响应JSON: %s\n", responseJSON)
	
	// JSON处理器示例
	fmt.Println("\nJSON处理器模式：")
	fmt.Println("w.Header().Set(\"Content-Type\", \"application/json\")")
	fmt.Println("json.NewEncoder(w).Encode(response)")
}

// 5. 中间件概念演示
func demonstrateMiddleware() {
	fmt.Println("中间件概念和实现：")
	
	// 日志中间件
	loggingMiddleware := func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			fmt.Printf("请求开始: %s %s\n", r.Method, r.URL.Path)
			
			next(w, r)
			
			duration := time.Since(start)
			fmt.Printf("请求完成: 耗时 %v\n", duration)
		}
	}
	
	// 认证中间件
	authMiddleware := func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			if token == "" {
				http.Error(w, "缺少认证令牌", http.StatusUnauthorized)
				return
			}
			
			fmt.Printf("认证令牌: %s\n", token)
			next(w, r)
		}
	}
	
	fmt.Printf("日志中间件类型: %T\n", loggingMiddleware)
	fmt.Printf("认证中间件类型: %T\n", authMiddleware)
	
	fmt.Println("\n中间件链式调用：")
	fmt.Println("handler = loggingMiddleware(authMiddleware(actualHandler))")
}

// 6. HTTP客户端演示
func demonstrateHTTPClient() {
	fmt.Println("HTTP客户端使用：")
	
	// 基础GET请求
	fmt.Println("\n基础GET请求：")
	fmt.Println("resp, err := http.Get(\"https://api.example.com/users\")")
	fmt.Println("defer resp.Body.Close()")
	fmt.Println("body, err := io.ReadAll(resp.Body)")
	
	// 自定义客户端
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	fmt.Printf("自定义客户端: %+v\n", client)
	
	// POST请求示例
	fmt.Println("\nPOST请求示例：")
	user := User{Name: "李四", Age: 30}
	jsonData, _ := json.Marshal(user)
	
	fmt.Printf("请求数据: %s\n", jsonData)
	fmt.Println("req, err := http.NewRequest(\"POST\", url, bytes.NewBuffer(jsonData))")
	fmt.Println("req.Header.Set(\"Content-Type\", \"application/json\")")
	fmt.Println("resp, err := client.Do(req)")
	
	// 请求头设置
	fmt.Println("\n常用请求头：")
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer token123",
		"User-Agent":    "Go-HTTP-Client/1.0",
		"Accept":        "application/json",
	}
	
	for key, value := range headers {
		fmt.Printf("%s: %s\n", key, value)
	}
}

// 实用工具函数

// WriteJSONResponse 写入JSON响应
func WriteJSONResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	
	response := Response{
		Code:    code,
		Message: getStatusMessage(code),
		Data:    data,
	}
	
	json.NewEncoder(w).Encode(response)
}

// WriteErrorResponse 写入错误响应
func WriteErrorResponse(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	
	response := Response{
		Code:    code,
		Message: message,
	}
	
	json.NewEncoder(w).Encode(response)
}

// getStatusMessage 获取状态码对应的消息
func getStatusMessage(code int) string {
	switch code {
	case 200:
		return "成功"
	case 201:
		return "创建成功"
	case 400:
		return "请求错误"
	case 401:
		return "未授权"
	case 404:
		return "未找到"
	case 500:
		return "服务器错误"
	default:
		return "未知状态"
	}
}

// ParseJSONBody 解析JSON请求体
func ParseJSONBody(r *http.Request, v interface{}) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	
	return json.Unmarshal(body, v)
}

// ExtractPathParam 提取路径参数
func ExtractPathParam(path, prefix string) string {
	return strings.TrimPrefix(path, prefix)
}

/*
学习要点总结：

1. HTTP服务器基础
   - http.HandleFunc() 注册路由
   - http.ListenAndServe() 启动服务器
   - http.ResponseWriter 和 *http.Request

2. 处理器函数
   - func(w http.ResponseWriter, r *http.Request)
   - 返回 http.HandlerFunc 的函数
   - 路径参数解析

3. HTTP方法处理
   - r.Method 获取请求方法
   - switch 语句处理不同方法
   - 请求体读取 io.ReadAll(r.Body)

4. JSON处理
   - json.Marshal() 编码
   - json.Unmarshal() 解码
   - json.NewEncoder(w).Encode() 直接写入响应

5. 中间件模式
   - 函数包装函数的模式
   - 请求前后处理逻辑
   - 中间件链式调用

6. HTTP客户端
   - http.Get() 简单GET请求
   - http.Client 自定义客户端
   - http.NewRequest() 创建自定义请求

实践建议：
- 从简单的Hello World服务器开始
- 逐步添加JSON处理功能
- 实现基础的CRUD操作
- 学习中间件的使用
- 练习HTTP客户端调用
*/