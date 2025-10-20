// API开发 - RESTful API设计与实现
// 学习目标：掌握RESTful API设计原则和Go语言实现

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Book 图书结构体
type Book struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	ISBN        string    `json:"isbn"`
	Price       float64   `json:"price"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// CreateBookRequest 创建图书请求
type CreateBookRequest struct {
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	ISBN        string    `json:"isbn"`
	Price       float64   `json:"price"`
	PublishedAt time.Time `json:"published_at"`
}

// UpdateBookRequest 更新图书请求
type UpdateBookRequest struct {
	Title       *string    `json:"title,omitempty"`
	Author      *string    `json:"author,omitempty"`
	ISBN        *string    `json:"isbn,omitempty"`
	Price       *float64   `json:"price,omitempty"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
}

// APIResponse 统一API响应格式
type APIResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// PaginationResponse 分页响应
type PaginationResponse struct {
	Items      interface{} `json:"items"`
	Total      int         `json:"total"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
	TotalPages int         `json:"total_pages"`
}

// BookService 图书服务（模拟数据层）
type BookService struct {
	books  []Book
	nextID int
}

// NewBookService 创建图书服务
func NewBookService() *BookService {
	return &BookService{
		books:  make([]Book, 0),
		nextID: 1,
	}
}

func main() {
	fmt.Println("=== Go API开发学习 ===")
	
	// 1. RESTful API设计原则
	fmt.Println("\n1. RESTful API设计原则")
	demonstrateRESTfulPrinciples()
	
	// 2. HTTP状态码使用
	fmt.Println("\n2. HTTP状态码使用")
	demonstrateStatusCodes()
	
	// 3. 请求和响应格式
	fmt.Println("\n3. 请求和响应格式")
	demonstrateRequestResponse()
	
	// 4. CRUD操作实现
	fmt.Println("\n4. CRUD操作实现")
	demonstrateCRUDOperations()
	
	// 5. 错误处理
	fmt.Println("\n5. 错误处理")
	demonstrateErrorHandling()
	
	// 6. 数据验证
	fmt.Println("\n6. 数据验证")
	demonstrateValidation()
	
	// 7. 分页和过滤
	fmt.Println("\n7. 分页和过滤")
	demonstratePaginationFiltering()
	
	// 8. API版本控制
	fmt.Println("\n8. API版本控制")
	demonstrateVersioning()
	
	fmt.Println("\n=== API开发学习完成 ===")
}

// 1. RESTful API设计原则演示
func demonstrateRESTfulPrinciples() {
	fmt.Println("RESTful API设计原则：")
	
	// 资源命名规则
	fmt.Println("\n资源命名规则：")
	apiEndpoints := map[string]string{
		"GET /api/v1/books":         "获取所有图书",
		"GET /api/v1/books/123":     "获取ID为123的图书",
		"POST /api/v1/books":        "创建新图书",
		"PUT /api/v1/books/123":     "更新ID为123的图书",
		"PATCH /api/v1/books/123":   "部分更新ID为123的图书",
		"DELETE /api/v1/books/123":  "删除ID为123的图书",
	}
	
	for endpoint, description := range apiEndpoints {
		fmt.Printf("%-30s %s\n", endpoint, description)
	}
	
	// REST原则
	fmt.Println("\nREST设计原则：")
	principles := []string{
		"1. 统一接口 - 使用标准HTTP方法",
		"2. 无状态 - 每个请求包含所有必要信息",
		"3. 可缓存 - 响应应该明确是否可缓存",
		"4. 客户端-服务器 - 关注点分离",
		"5. 分层系统 - 支持中间层（代理、网关等）",
		"6. 按需代码 - 可选的代码下载",
	}
	
	for _, principle := range principles {
		fmt.Println(principle)
	}
}

// 2. HTTP状态码使用演示
func demonstrateStatusCodes() {
	fmt.Println("常用HTTP状态码：")
	
	statusCodes := map[int]string{
		200: "OK - 请求成功",
		201: "Created - 资源创建成功",
		204: "No Content - 请求成功但无返回内容",
		400: "Bad Request - 请求参数错误",
		401: "Unauthorized - 未授权",
		403: "Forbidden - 禁止访问",
		404: "Not Found - 资源不存在",
		409: "Conflict - 资源冲突",
		422: "Unprocessable Entity - 数据验证失败",
		500: "Internal Server Error - 服务器内部错误",
	}
	
	for code, description := range statusCodes {
		fmt.Printf("%d: %s\n", code, description)
	}
	
	// 状态码使用场景
	fmt.Println("\n状态码使用场景：")
	scenarios := map[string]int{
		"获取资源成功":     200,
		"创建资源成功":     201,
		"删除资源成功":     204,
		"请求参数格式错误":   400,
		"资源不存在":      404,
		"数据验证失败":     422,
		"服务器处理异常":    500,
	}
	
	for scenario, code := range scenarios {
		fmt.Printf("%-15s -> %d\n", scenario, code)
	}
}

// 3. 请求和响应格式演示
func demonstrateRequestResponse() {
	fmt.Println("API请求和响应格式：")
	
	// 创建请求示例
	createReq := CreateBookRequest{
		Title:       "Go语言编程",
		Author:      "张三",
		ISBN:        "978-7-111-12345-6",
		Price:       89.90,
		PublishedAt: time.Now(),
	}
	
	reqJSON, _ := json.MarshalIndent(createReq, "", "  ")
	fmt.Printf("\n创建请求格式：\n%s\n", reqJSON)
	
	// 成功响应示例
	book := Book{
		ID:          1,
		Title:       createReq.Title,
		Author:      createReq.Author,
		ISBN:        createReq.ISBN,
		Price:       createReq.Price,
		PublishedAt: createReq.PublishedAt,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	
	successResp := APIResponse{
		Success: true,
		Code:    201,
		Message: "图书创建成功",
		Data:    book,
	}
	
	respJSON, _ := json.MarshalIndent(successResp, "", "  ")
	fmt.Printf("\n成功响应格式：\n%s\n", respJSON)
	
	// 错误响应示例
	errorResp := APIResponse{
		Success: false,
		Code:    400,
		Message: "请求参数错误",
		Error:   "title字段不能为空",
	}
	
	errorJSON, _ := json.MarshalIndent(errorResp, "", "  ")
	fmt.Printf("\n错误响应格式：\n%s\n", errorJSON)
}

// 4. CRUD操作实现演示
func demonstrateCRUDOperations() {
	fmt.Println("CRUD操作实现：")
	
	service := NewBookService()
	
	// Create - 创建
	fmt.Println("\n1. Create (创建):")
	book1 := service.CreateBook(CreateBookRequest{
		Title:       "Go语言实战",
		Author:      "李四",
		ISBN:        "978-7-111-11111-1",
		Price:       99.00,
		PublishedAt: time.Now(),
	})
	fmt.Printf("创建图书: %+v\n", book1)
	
	// Read - 读取
	fmt.Println("\n2. Read (读取):")
	foundBook := service.GetBookByID(1)
	if foundBook != nil {
		fmt.Printf("找到图书: %+v\n", *foundBook)
	}
	
	allBooks := service.GetAllBooks()
	fmt.Printf("所有图书数量: %d\n", len(allBooks))
	
	// Update - 更新
	fmt.Println("\n3. Update (更新):")
	newTitle := "Go语言高级编程"
	newPrice := 129.00
	updateReq := UpdateBookRequest{
		Title: &newTitle,
		Price: &newPrice,
	}
	updatedBook := service.UpdateBook(1, updateReq)
	if updatedBook != nil {
		fmt.Printf("更新后图书: %+v\n", *updatedBook)
	}
	
	// Delete - 删除
	fmt.Println("\n4. Delete (删除):")
	deleted := service.DeleteBook(1)
	fmt.Printf("删除结果: %t\n", deleted)
	fmt.Printf("删除后图书数量: %d\n", len(service.GetAllBooks()))
}

// 5. 错误处理演示
func demonstrateErrorHandling() {
	fmt.Println("API错误处理策略：")
	
	// 错误类型分类
	errorTypes := map[string]string{
		"ValidationError":   "数据验证错误 - 422",
		"NotFoundError":     "资源不存在错误 - 404",
		"ConflictError":     "资源冲突错误 - 409",
		"AuthenticationError": "认证错误 - 401",
		"AuthorizationError":  "授权错误 - 403",
		"InternalError":     "内部服务错误 - 500",
	}
	
	for errorType, description := range errorTypes {
		fmt.Printf("%-20s: %s\n", errorType, description)
	}
	
	// 错误响应格式
	fmt.Println("\n标准错误响应格式：")
	errorResponse := map[string]interface{}{
		"success": false,
		"code":    422,
		"message": "数据验证失败",
		"error":   "title字段长度不能超过100个字符",
		"details": map[string]string{
			"field": "title",
			"rule":  "max_length",
			"value": "100",
		},
	}
	
	errorJSON, _ := json.MarshalIndent(errorResponse, "", "  ")
	fmt.Printf("%s\n", errorJSON)
}

// 6. 数据验证演示
func demonstrateValidation() {
	fmt.Println("数据验证实现：")
	
	// 验证规则示例
	fmt.Println("\n验证规则：")
	validationRules := map[string]string{
		"title":        "必填，长度1-100字符",
		"author":       "必填，长度1-50字符",
		"isbn":         "必填，符合ISBN格式",
		"price":        "必填，大于0的数字",
		"published_at": "必填，有效的日期格式",
	}
	
	for field, rule := range validationRules {
		fmt.Printf("%-15s: %s\n", field, rule)
	}
	
	// 验证函数示例
	fmt.Println("\n验证函数实现：")
	req := CreateBookRequest{
		Title:  "",  // 无效：空标题
		Author: "作者名字",
		ISBN:   "invalid-isbn",  // 无效：ISBN格式
		Price:  -10.0,  // 无效：负价格
	}
	
	errors := validateCreateBookRequest(req)
	if len(errors) > 0 {
		fmt.Println("验证错误：")
		for field, error := range errors {
			fmt.Printf("  %s: %s\n", field, error)
		}
	}
}

// 7. 分页和过滤演示
func demonstratePaginationFiltering() {
	fmt.Println("分页和过滤实现：")
	
	// 分页参数
	fmt.Println("\n分页参数：")
	paginationParams := map[string]string{
		"page":      "页码，从1开始",
		"page_size": "每页大小，默认10，最大100",
		"sort":      "排序字段，如：title,created_at",
		"order":     "排序方向，asc或desc",
	}
	
	for param, description := range paginationParams {
		fmt.Printf("%-10s: %s\n", param, description)
	}
	
	// 过滤参数
	fmt.Println("\n过滤参数：")
	filterParams := map[string]string{
		"title":     "标题模糊搜索",
		"author":    "作者精确匹配",
		"min_price": "最低价格",
		"max_price": "最高价格",
		"year":      "出版年份",
	}
	
	for param, description := range filterParams {
		fmt.Printf("%-10s: %s\n", param, description)
	}
	
	// 分页响应示例
	fmt.Println("\n分页响应格式：")
	paginationResp := PaginationResponse{
		Items:      []Book{},  // 实际数据
		Total:      150,
		Page:       2,
		PageSize:   10,
		TotalPages: 15,
	}
	
	pageJSON, _ := json.MarshalIndent(paginationResp, "", "  ")
	fmt.Printf("%s\n", pageJSON)
}

// 8. API版本控制演示
func demonstrateVersioning() {
	fmt.Println("API版本控制策略：")
	
	// 版本控制方法
	fmt.Println("\n版本控制方法：")
	versioningMethods := map[string]string{
		"URL路径版本":    "/api/v1/books, /api/v2/books",
		"请求头版本":     "Accept: application/vnd.api+json;version=1",
		"查询参数版本":    "/api/books?version=1",
		"子域名版本":     "v1.api.example.com/books",
	}
	
	for method, example := range versioningMethods {
		fmt.Printf("%-12s: %s\n", method, example)
	}
	
	// 版本兼容性
	fmt.Println("\n版本兼容性原则：")
	compatibilityRules := []string{
		"1. 向后兼容 - 新版本支持旧版本功能",
		"2. 渐进式弃用 - 提前通知API变更",
		"3. 文档维护 - 每个版本独立文档",
		"4. 测试覆盖 - 多版本并行测试",
	}
	
	for _, rule := range compatibilityRules {
		fmt.Println(rule)
	}
}

// BookService 方法实现

func (s *BookService) CreateBook(req CreateBookRequest) Book {
	book := Book{
		ID:          s.nextID,
		Title:       req.Title,
		Author:      req.Author,
		ISBN:        req.ISBN,
		Price:       req.Price,
		PublishedAt: req.PublishedAt,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	
	s.books = append(s.books, book)
	s.nextID++
	
	return book
}

func (s *BookService) GetBookByID(id int) *Book {
	for _, book := range s.books {
		if book.ID == id {
			return &book
		}
	}
	return nil
}

func (s *BookService) GetAllBooks() []Book {
	return s.books
}

func (s *BookService) UpdateBook(id int, req UpdateBookRequest) *Book {
	for i, book := range s.books {
		if book.ID == id {
			if req.Title != nil {
				s.books[i].Title = *req.Title
			}
			if req.Author != nil {
				s.books[i].Author = *req.Author
			}
			if req.ISBN != nil {
				s.books[i].ISBN = *req.ISBN
			}
			if req.Price != nil {
				s.books[i].Price = *req.Price
			}
			if req.PublishedAt != nil {
				s.books[i].PublishedAt = *req.PublishedAt
			}
			s.books[i].UpdatedAt = time.Now()
			
			return &s.books[i]
		}
	}
	return nil
}

func (s *BookService) DeleteBook(id int) bool {
	for i, book := range s.books {
		if book.ID == id {
			s.books = append(s.books[:i], s.books[i+1:]...)
			return true
		}
	}
	return false
}

// 工具函数

// validateCreateBookRequest 验证创建图书请求
func validateCreateBookRequest(req CreateBookRequest) map[string]string {
	errors := make(map[string]string)
	
	if req.Title == "" {
		errors["title"] = "标题不能为空"
	} else if len(req.Title) > 100 {
		errors["title"] = "标题长度不能超过100个字符"
	}
	
	if req.Author == "" {
		errors["author"] = "作者不能为空"
	} else if len(req.Author) > 50 {
		errors["author"] = "作者名长度不能超过50个字符"
	}
	
	if req.ISBN == "" {
		errors["isbn"] = "ISBN不能为空"
	} else if !isValidISBN(req.ISBN) {
		errors["isbn"] = "ISBN格式不正确"
	}
	
	if req.Price <= 0 {
		errors["price"] = "价格必须大于0"
	}
	
	return errors
}

// isValidISBN 验证ISBN格式（简化版）
func isValidISBN(isbn string) bool {
	// 简化的ISBN验证，实际应该更严格
	return len(isbn) >= 10 && strings.Contains(isbn, "-")
}

// WriteAPIResponse 写入API响应
func WriteAPIResponse(w http.ResponseWriter, code int, success bool, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	
	response := APIResponse{
		Success: success,
		Code:    code,
		Message: message,
		Data:    data,
	}
	
	json.NewEncoder(w).Encode(response)
}

// WriteAPIError 写入API错误响应
func WriteAPIError(w http.ResponseWriter, code int, message string, err string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	
	response := APIResponse{
		Success: false,
		Code:    code,
		Message: message,
		Error:   err,
	}
	
	json.NewEncoder(w).Encode(response)
}

// ParsePaginationParams 解析分页参数
func ParsePaginationParams(r *http.Request) (page, pageSize int) {
	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		page = 1
	} else {
		page, _ = strconv.Atoi(pageStr)
		if page < 1 {
			page = 1
		}
	}
	
	pageSizeStr := r.URL.Query().Get("page_size")
	if pageSizeStr == "" {
		pageSize = 10
	} else {
		pageSize, _ = strconv.Atoi(pageSizeStr)
		if pageSize < 1 {
			pageSize = 10
		} else if pageSize > 100 {
			pageSize = 100
		}
	}
	
	return page, pageSize
}

// ExtractIDFromPath 从路径中提取ID
func ExtractIDFromPath(path string) (int, error) {
	parts := strings.Split(path, "/")
	if len(parts) == 0 {
		return 0, fmt.Errorf("无效的路径")
	}
	
	idStr := parts[len(parts)-1]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("无效的ID格式")
	}
	
	return id, nil
}

/*
学习要点总结：

1. RESTful API设计
   - 使用标准HTTP方法 (GET, POST, PUT, DELETE)
   - 资源导向的URL设计
   - 统一的响应格式

2. HTTP状态码
   - 2xx: 成功响应
   - 4xx: 客户端错误
   - 5xx: 服务器错误

3. 请求响应格式
   - JSON作为主要数据格式
   - 统一的响应结构
   - 错误信息标准化

4. CRUD操作
   - Create: POST /resources
   - Read: GET /resources, GET /resources/:id
   - Update: PUT /resources/:id, PATCH /resources/:id
   - Delete: DELETE /resources/:id

5. 错误处理
   - 分类错误类型
   - 提供详细错误信息
   - 使用合适的状态码

6. 数据验证
   - 输入参数验证
   - 业务规则验证
   - 返回具体错误信息

7. 分页和过滤
   - 查询参数支持
   - 分页元数据返回
   - 排序和过滤功能

8. 版本控制
   - URL路径版本控制
   - 向后兼容性
   - 版本迁移策略

实践建议：
- 从简单的CRUD API开始
- 逐步添加验证和错误处理
- 实现分页和搜索功能
- 学习API文档编写
- 进行API测试和调试
*/