// Go高级特性 - 并发、反射、泛型等高级功能
// 学习目标：掌握Go语言的高级特性和最佳实践

package main

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"runtime"
	"sync"
	"time"
)

// 泛型示例结构体和接口
type Comparable interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
	~float32 | ~float64 | ~string
}

// Stack 泛型栈实现
type Stack[T any] struct {
	items []T
	mutex sync.RWMutex
}

// Cache 泛型缓存实现
type Cache[K comparable, V any] struct {
	data  map[K]V
	mutex sync.RWMutex
}

// Worker 工作者模式
type Worker struct {
	ID       int
	JobQueue chan Job
	quit     chan bool
}

// Job 任务接口
type Job interface {
	Execute() error
}

// SimpleJob 简单任务实现
type SimpleJob struct {
	ID   int
	Data string
}

// WorkerPool 工作者池
type WorkerPool struct {
	workers   []*Worker
	jobQueue  chan Job
	quit      chan bool
	wg        sync.WaitGroup
}

// User 用户结构体（用于反射示例）
type User struct {
	ID       int    `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required,min=2,max=50"`
	Email    string `json:"email" validate:"required,email"`
	Age      int    `json:"age" validate:"min=0,max=150"`
	IsActive bool   `json:"is_active"`
}

func main() {
	fmt.Println("=== Go高级特性学习 ===")
	
	// 1. 并发编程进阶
	fmt.Println("\n1. 并发编程进阶")
	demonstrateAdvancedConcurrency()
	
	// 2. 反射机制
	fmt.Println("\n2. 反射机制")
	demonstrateReflection()
	
	// 3. 泛型编程
	fmt.Println("\n3. 泛型编程")
	demonstrateGenerics()
	
	// 4. 上下文管理
	fmt.Println("\n4. 上下文管理")
	demonstrateContext()
	
	// 5. 内存管理
	fmt.Println("\n5. 内存管理")
	demonstrateMemoryManagement()
	
	// 6. 性能优化
	fmt.Println("\n6. 性能优化")
	demonstratePerformanceOptimization()
	
	// 7. 设计模式
	fmt.Println("\n7. 设计模式")
	demonstrateDesignPatterns()
	
	// 8. 测试和基准测试
	fmt.Println("\n8. 测试和基准测试")
	demonstrateTesting()
	
	fmt.Println("\n=== 高级特性学习完成 ===")
}

// 1. 并发编程进阶演示
func demonstrateAdvancedConcurrency() {
	fmt.Println("高级并发编程模式：")
	
	// 工作者池模式
	fmt.Println("\n1. 工作者池模式：")
	demonstrateWorkerPool()
	
	// 扇入扇出模式
	fmt.Println("\n2. 扇入扇出模式：")
	demonstrateFanInFanOut()
	
	// 管道模式
	fmt.Println("\n3. 管道模式：")
	demonstratePipeline()
	
	// 超时和取消
	fmt.Println("\n4. 超时和取消：")
	demonstrateTimeoutCancel()
}

// 工作者池演示
func demonstrateWorkerPool() {
	fmt.Println("工作者池实现：")
	
	// 创建工作者池
	pool := NewWorkerPool(3, 10)
	pool.Start()
	
	// 提交任务
	for i := 1; i <= 5; i++ {
		job := &SimpleJob{
			ID:   i,
			Data: fmt.Sprintf("任务数据 %d", i),
		}
		pool.Submit(job)
	}
	
	// 等待完成并停止
	time.Sleep(2 * time.Second)
	pool.Stop()
	
	fmt.Println("工作者池演示完成")
}

// 扇入扇出模式演示
func demonstrateFanInFanOut() {
	fmt.Println("扇入扇出模式：")
	
	// 扇出：一个输入分发到多个处理器
	input := make(chan int, 10)
	
	// 启动多个处理器（扇出）
	processor1 := fanOut(input, "处理器1")
	processor2 := fanOut(input, "处理器2")
	processor3 := fanOut(input, "处理器3")
	
	// 扇入：多个处理器的结果合并
	output := fanIn(processor1, processor2, processor3)
	
	// 发送数据
	go func() {
		for i := 1; i <= 9; i++ {
			input <- i
		}
		close(input)
	}()
	
	// 接收结果
	for result := range output {
		fmt.Printf("结果: %s\n", result)
	}
}

// 管道模式演示
func demonstratePipeline() {
	fmt.Println("管道模式：")
	
	// 创建管道：数字生成 -> 平方 -> 过滤偶数 -> 输出
	numbers := generateNumbers(1, 10)
	squares := squareNumbers(numbers)
	evens := filterEvenNumbers(squares)
	
	// 处理结果
	for result := range evens {
		fmt.Printf("管道结果: %d\n", result)
	}
}

// 超时和取消演示
func demonstrateTimeoutCancel() {
	fmt.Println("超时和取消机制：")
	
	// 超时示例
	fmt.Println("超时示例：")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	result := make(chan string, 1)
	go func() {
		// 模拟长时间运行的任务
		time.Sleep(3 * time.Second)
		result <- "任务完成"
	}()
	
	select {
	case res := <-result:
		fmt.Printf("收到结果: %s\n", res)
	case <-ctx.Done():
		fmt.Printf("任务超时: %v\n", ctx.Err())
	}
	
	// 取消示例
	fmt.Println("取消示例：")
	ctx2, cancel2 := context.WithCancel(context.Background())
	
	go func() {
		time.Sleep(1 * time.Second)
		cancel2() // 1秒后取消
	}()
	
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("任务正常完成")
	case <-ctx2.Done():
		fmt.Printf("任务被取消: %v\n", ctx2.Err())
	}
}

// 2. 反射机制演示
func demonstrateReflection() {
	fmt.Println("反射机制应用：")
	
	// 类型检查
	fmt.Println("\n1. 类型检查：")
	demonstrateTypeInspection()
	
	// 结构体标签
	fmt.Println("\n2. 结构体标签：")
	demonstrateStructTags()
	
	// 动态调用
	fmt.Println("\n3. 动态调用：")
	demonstrateDynamicCalls()
	
	// 反射性能考虑
	fmt.Println("\n4. 反射性能考虑：")
	demonstrateReflectionPerformance()
}

// 类型检查演示
func demonstrateTypeInspection() {
	user := User{
		ID:       1,
		Name:     "张三",
		Email:    "zhangsan@example.com",
		Age:      25,
		IsActive: true,
	}
	
	t := reflect.TypeOf(user)
	v := reflect.ValueOf(user)
	
	fmt.Printf("类型名称: %s\n", t.Name())
	fmt.Printf("类型种类: %s\n", t.Kind())
	fmt.Printf("字段数量: %d\n", t.NumField())
	
	// 遍历字段
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		
		fmt.Printf("字段 %s: 类型=%s, 值=%v\n", 
			field.Name, field.Type, value.Interface())
	}
}

// 结构体标签演示
func demonstrateStructTags() {
	t := reflect.TypeOf(User{})
	
	fmt.Println("结构体标签信息：")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		
		jsonTag := field.Tag.Get("json")
		validateTag := field.Tag.Get("validate")
		
		fmt.Printf("字段 %s: json=%s, validate=%s\n", 
			field.Name, jsonTag, validateTag)
	}
	
	// 标签验证示例
	fmt.Println("\n标签验证示例：")
	user := User{Name: "A", Age: -1} // 无效数据
	errors := validateStruct(user)
	
	if len(errors) > 0 {
		fmt.Println("验证错误：")
		for field, err := range errors {
			fmt.Printf("  %s: %s\n", field, err)
		}
	}
}

// 动态调用演示
func demonstrateDynamicCalls() {
	user := &User{
		ID:   1,
		Name: "李四",
		Age:  30,
	}
	
	v := reflect.ValueOf(user)
	
	// 动态设置字段值
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	
	nameField := v.FieldByName("Name")
	if nameField.IsValid() && nameField.CanSet() {
		nameField.SetString("王五")
		fmt.Printf("动态设置后的名称: %s\n", user.Name)
	}
	
	ageField := v.FieldByName("Age")
	if ageField.IsValid() && ageField.CanSet() {
		ageField.SetInt(35)
		fmt.Printf("动态设置后的年龄: %d\n", user.Age)
	}
}

// 反射性能考虑演示
func demonstrateReflectionPerformance() {
	fmt.Println("反射性能对比：")
	
	user := User{Name: "测试用户", Age: 25}
	
	// 直接访问
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		_ = user.Name
	}
	directTime := time.Since(start)
	
	// 反射访问
	v := reflect.ValueOf(user)
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		_ = v.FieldByName("Name").String()
	}
	reflectTime := time.Since(start)
	
	fmt.Printf("直接访问耗时: %v\n", directTime)
	fmt.Printf("反射访问耗时: %v\n", reflectTime)
	fmt.Printf("性能差异: %.2fx\n", float64(reflectTime)/float64(directTime))
}

// 3. 泛型编程演示
func demonstrateGenerics() {
	fmt.Println("泛型编程应用：")
	
	// 泛型函数
	fmt.Println("\n1. 泛型函数：")
	demonstrateGenericFunctions()
	
	// 泛型数据结构
	fmt.Println("\n2. 泛型数据结构：")
	demonstrateGenericDataStructures()
	
	// 类型约束
	fmt.Println("\n3. 类型约束：")
	demonstrateTypeConstraints()
}

// 泛型函数演示
func demonstrateGenericFunctions() {
	// 泛型比较函数
	fmt.Printf("Max(10, 20) = %d\n", Max(10, 20))
	fmt.Printf("Max(3.14, 2.71) = %.2f\n", Max(3.14, 2.71))
	fmt.Printf("Max(\"apple\", \"banana\") = %s\n", Max("apple", "banana"))
	
	// 泛型切片操作
	numbers := []int{1, 2, 3, 4, 5}
	doubled := Map(numbers, func(x int) int { return x * 2 })
	fmt.Printf("原数组: %v\n", numbers)
	fmt.Printf("翻倍后: %v\n", doubled)
	
	evens := Filter(numbers, func(x int) bool { return x%2 == 0 })
	fmt.Printf("偶数: %v\n", evens)
	
	sum := Reduce(numbers, 0, func(acc, x int) int { return acc + x })
	fmt.Printf("求和: %d\n", sum)
}

// 泛型数据结构演示
func demonstrateGenericDataStructures() {
	// 泛型栈
	fmt.Println("泛型栈操作：")
	stack := NewStack[int]()
	
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	
	fmt.Printf("栈大小: %d\n", stack.Size())
	
	for !stack.IsEmpty() {
		value, _ := stack.Pop()
		fmt.Printf("弹出: %d\n", value)
	}
	
	// 泛型缓存
	fmt.Println("\n泛型缓存操作：")
	cache := NewCache[string, User]()
	
	user1 := User{ID: 1, Name: "用户1", Age: 25}
	user2 := User{ID: 2, Name: "用户2", Age: 30}
	
	cache.Set("user1", user1)
	cache.Set("user2", user2)
	
	if user, exists := cache.Get("user1"); exists {
		fmt.Printf("缓存命中: %+v\n", user)
	}
	
	fmt.Printf("缓存大小: %d\n", cache.Size())
}

// 类型约束演示
func demonstrateTypeConstraints() {
	// 数值类型约束
	fmt.Printf("整数求和: %d\n", Sum([]int{1, 2, 3, 4, 5}))
	fmt.Printf("浮点数求和: %.2f\n", Sum([]float64{1.1, 2.2, 3.3}))
	
	// 自定义约束
	type MyInt int
	myNumbers := []MyInt{1, 2, 3}
	fmt.Printf("自定义类型求和: %d\n", Sum(myNumbers))
}

// 4. 上下文管理演示
func demonstrateContext() {
	fmt.Println("上下文管理应用：")
	
	// 超时控制
	fmt.Println("\n1. 超时控制：")
	demonstrateContextTimeout()
	
	// 取消传播
	fmt.Println("\n2. 取消传播：")
	demonstrateContextCancellation()
	
	// 值传递
	fmt.Println("\n3. 值传递：")
	demonstrateContextValues()
}

// 上下文超时演示
func demonstrateContextTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	
	result := make(chan string, 1)
	
	go func() {
		// 模拟耗时操作
		time.Sleep(500 * time.Millisecond)
		result <- "操作完成"
	}()
	
	select {
	case res := <-result:
		fmt.Printf("结果: %s\n", res)
	case <-ctx.Done():
		fmt.Printf("操作超时: %v\n", ctx.Err())
	}
}

// 上下文取消传播演示
func demonstrateContextCancellation() {
	ctx, cancel := context.WithCancel(context.Background())
	
	// 启动多个goroutine
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			select {
			case <-time.After(2 * time.Second):
				fmt.Printf("Goroutine %d 正常完成\n", id)
			case <-ctx.Done():
				fmt.Printf("Goroutine %d 被取消: %v\n", id, ctx.Err())
			}
		}(i)
	}
	
	// 1秒后取消所有操作
	time.Sleep(1 * time.Second)
	cancel()
	
	wg.Wait()
	fmt.Println("所有goroutine已完成")
}

// 上下文值传递演示
func demonstrateContextValues() {
	type contextKey string
	
	userIDKey := contextKey("userID")
	requestIDKey := contextKey("requestID")
	
	// 创建带值的上下文
	ctx := context.WithValue(context.Background(), userIDKey, "user123")
	ctx = context.WithValue(ctx, requestIDKey, "req456")
	
	// 模拟处理请求
	processRequest(ctx, userIDKey, requestIDKey)
}

func processRequest(ctx context.Context, userIDKey, requestIDKey interface{}) {
	userID := ctx.Value(userIDKey)
	requestID := ctx.Value(requestIDKey)
	
	fmt.Printf("处理请求 - 用户ID: %v, 请求ID: %v\n", userID, requestID)
}

// 5. 内存管理演示
func demonstrateMemoryManagement() {
	fmt.Println("内存管理和垃圾回收：")
	
	// 内存分配
	fmt.Println("\n1. 内存分配模式：")
	demonstrateMemoryAllocation()
	
	// 垃圾回收
	fmt.Println("\n2. 垃圾回收控制：")
	demonstrateGarbageCollection()
	
	// 内存泄漏预防
	fmt.Println("\n3. 内存泄漏预防：")
	demonstrateMemoryLeakPrevention()
}

// 内存分配演示
func demonstrateMemoryAllocation() {
	var m1, m2 runtime.MemStats
	
	// 获取初始内存状态
	runtime.ReadMemStats(&m1)
	
	// 分配大量内存
	data := make([][]byte, 1000)
	for i := range data {
		data[i] = make([]byte, 1024) // 1KB per slice
	}
	
	// 获取分配后内存状态
	runtime.ReadMemStats(&m2)
	
	fmt.Printf("分配前堆内存: %d KB\n", m1.HeapAlloc/1024)
	fmt.Printf("分配后堆内存: %d KB\n", m2.HeapAlloc/1024)
	fmt.Printf("内存增长: %d KB\n", (m2.HeapAlloc-m1.HeapAlloc)/1024)
	
	// 清理引用
	data = nil
}

// 垃圾回收演示
func demonstrateGarbageCollection() {
	var m1, m2 runtime.MemStats
	
	// 分配内存
	for i := 0; i < 1000; i++ {
		_ = make([]byte, 1024*1024) // 1MB
	}
	
	runtime.ReadMemStats(&m1)
	fmt.Printf("GC前堆内存: %d MB\n", m1.HeapAlloc/1024/1024)
	
	// 强制垃圾回收
	runtime.GC()
	
	runtime.ReadMemStats(&m2)
	fmt.Printf("GC后堆内存: %d MB\n", m2.HeapAlloc/1024/1024)
	fmt.Printf("GC次数: %d\n", m2.NumGC-m1.NumGC)
}

// 内存泄漏预防演示
func demonstrateMemoryLeakPrevention() {
	fmt.Println("常见内存泄漏场景和预防：")
	
	leakScenarios := []string{
		"1. Goroutine泄漏 - 确保goroutine能够正常退出",
		"2. 定时器泄漏 - 使用完毕后调用Stop()",
		"3. 循环引用 - 避免强引用循环",
		"4. 大切片引用 - 复制需要的部分而不是引用整个切片",
		"5. 全局变量 - 谨慎使用全局变量存储大对象",
	}
	
	for _, scenario := range leakScenarios {
		fmt.Println(scenario)
	}
	
	// 正确的定时器使用
	fmt.Println("\n正确的定时器使用：")
	timer := time.NewTimer(1 * time.Second)
	defer timer.Stop() // 确保定时器被停止
	
	select {
	case <-timer.C:
		fmt.Println("定时器触发")
	case <-time.After(500 * time.Millisecond):
		fmt.Println("提前退出")
	}
}

// 6. 性能优化演示
func demonstratePerformanceOptimization() {
	fmt.Println("性能优化技巧：")
	
	// 字符串拼接优化
	fmt.Println("\n1. 字符串拼接优化：")
	demonstrateStringOptimization()
	
	// 切片预分配
	fmt.Println("\n2. 切片预分配：")
	demonstrateSliceOptimization()
	
	// 并发优化
	fmt.Println("\n3. 并发优化：")
	demonstrateConcurrencyOptimization()
}

// 字符串拼接优化演示
func demonstrateStringOptimization() {
	const iterations = 10000
	
	// 使用 + 操作符（低效）
	start := time.Now()
	result1 := ""
	for i := 0; i < iterations; i++ {
		result1 += "a"
	}
	time1 := time.Since(start)
	
	// 使用 strings.Builder（高效）
	start = time.Now()
	var builder strings.Builder
	builder.Grow(iterations) // 预分配容量
	for i := 0; i < iterations; i++ {
		builder.WriteString("a")
	}
	result2 := builder.String()
	time2 := time.Since(start)
	
	fmt.Printf("+ 操作符耗时: %v (长度: %d)\n", time1, len(result1))
	fmt.Printf("Builder耗时: %v (长度: %d)\n", time2, len(result2))
	fmt.Printf("性能提升: %.2fx\n", float64(time1)/float64(time2))
}

// 切片预分配演示
func demonstrateSliceOptimization() {
	const size = 100000
	
	// 不预分配（低效）
	start := time.Now()
	var slice1 []int
	for i := 0; i < size; i++ {
		slice1 = append(slice1, i)
	}
	time1 := time.Since(start)
	
	// 预分配（高效）
	start = time.Now()
	slice2 := make([]int, 0, size)
	for i := 0; i < size; i++ {
		slice2 = append(slice2, i)
	}
	time2 := time.Since(start)
	
	fmt.Printf("不预分配耗时: %v\n", time1)
	fmt.Printf("预分配耗时: %v\n", time2)
	fmt.Printf("性能提升: %.2fx\n", float64(time1)/float64(time2))
}

// 并发优化演示
func demonstrateConcurrencyOptimization() {
	const workload = 1000000
	
	// 串行处理
	start := time.Now()
	sum1 := 0
	for i := 1; i <= workload; i++ {
		sum1 += i * i
	}
	time1 := time.Since(start)
	
	// 并行处理
	start = time.Now()
	numWorkers := runtime.NumCPU()
	chunkSize := workload / numWorkers
	
	var wg sync.WaitGroup
	var mu sync.Mutex
	sum2 := 0
	
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			localSum := 0
			for i := start; i <= end; i++ {
				localSum += i * i
			}
			mu.Lock()
			sum2 += localSum
			mu.Unlock()
		}(w*chunkSize+1, (w+1)*chunkSize)
	}
	
	wg.Wait()
	time2 := time.Since(start)
	
	fmt.Printf("串行处理耗时: %v (结果: %d)\n", time1, sum1)
	fmt.Printf("并行处理耗时: %v (结果: %d)\n", time2, sum2)
	fmt.Printf("性能提升: %.2fx\n", float64(time1)/float64(time2))
}

// 7. 设计模式演示
func demonstrateDesignPatterns() {
	fmt.Println("常用设计模式：")
	
	patterns := []string{
		"1. 单例模式 - 确保类只有一个实例",
		"2. 工厂模式 - 创建对象的接口",
		"3. 观察者模式 - 对象间的一对多依赖",
		"4. 策略模式 - 算法族的封装",
		"5. 装饰器模式 - 动态添加对象功能",
		"6. 适配器模式 - 接口转换",
	}
	
	for _, pattern := range patterns {
		fmt.Println(pattern)
	}
	
	// 单例模式示例
	fmt.Println("\n单例模式示例：")
	instance1 := GetSingletonInstance()
	instance2 := GetSingletonInstance()
	
	fmt.Printf("实例1地址: %p\n", instance1)
	fmt.Printf("实例2地址: %p\n", instance2)
	fmt.Printf("是否为同一实例: %t\n", instance1 == instance2)
}

// 8. 测试和基准测试演示
func demonstrateTesting() {
	fmt.Println("测试最佳实践：")
	
	testingPractices := []string{
		"1. 单元测试 - 测试单个函数或方法",
		"2. 集成测试 - 测试组件间的交互",
		"3. 基准测试 - 性能测试和优化",
		"4. 模糊测试 - 自动生成测试输入",
		"5. 表驱动测试 - 使用测试表格",
		"6. 测试覆盖率 - 确保代码覆盖",
	}
	
	for _, practice := range testingPractices {
		fmt.Println(practice)
	}
	
	// 基准测试示例
	fmt.Println("\n基准测试示例代码：")
	benchmarkExample := `
func BenchmarkStringConcat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        result := ""
        for j := 0; j < 100; j++ {
            result += "a"
        }
    }
}

func BenchmarkStringBuilder(b *testing.B) {
    for i := 0; i < b.N; i++ {
        var builder strings.Builder
        for j := 0; j < 100; j++ {
            builder.WriteString("a")
        }
        _ = builder.String()
    }
}
`
	fmt.Println(benchmarkExample)
}

// 工具函数和类型实现

// 泛型函数实现
func Max[T Comparable](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Map[T, R any](slice []T, fn func(T) R) []R {
	result := make([]R, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

func Filter[T any](slice []T, fn func(T) bool) []T {
	var result []T
	for _, v := range slice {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func Reduce[T, R any](slice []T, initial R, fn func(R, T) R) R {
	result := initial
	for _, v := range slice {
		result = fn(result, v)
	}
	return result
}

func Sum[T Comparable](numbers []T) T {
	var sum T
	for _, num := range numbers {
		sum += num
	}
	return sum
}

// Stack 实现
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		items: make([]T, 0),
	}
}

func (s *Stack[T]) Push(item T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	
	index := len(s.items) - 1
	item := s.items[index]
	s.items = s.items[:index]
	return item, true
}

func (s *Stack[T]) Size() int {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return len(s.items)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Size() == 0
}

// Cache 实现
func NewCache[K comparable, V any]() *Cache[K, V] {
	return &Cache[K, V]{
		data: make(map[K]V),
	}
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.data[key] = value
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	value, exists := c.data[key]
	return value, exists
}

func (c *Cache[K, V]) Size() int {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return len(c.data)
}

// 工作者池实现
func NewWorkerPool(numWorkers, queueSize int) *WorkerPool {
	return &WorkerPool{
		workers:  make([]*Worker, numWorkers),
		jobQueue: make(chan Job, queueSize),
		quit:     make(chan bool),
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < len(wp.workers); i++ {
		worker := &Worker{
			ID:       i + 1,
			JobQueue: wp.jobQueue,
			quit:     make(chan bool),
		}
		wp.workers[i] = worker
		wp.wg.Add(1)
		go worker.Start(&wp.wg)
	}
}

func (wp *WorkerPool) Submit(job Job) {
	wp.jobQueue <- job
}

func (wp *WorkerPool) Stop() {
	close(wp.jobQueue)
	
	for _, worker := range wp.workers {
		worker.Stop()
	}
	
	wp.wg.Wait()
}

func (w *Worker) Start(wg *sync.WaitGroup) {
	defer wg.Done()
	
	for {
		select {
		case job, ok := <-w.JobQueue:
			if !ok {
				fmt.Printf("Worker %d 停止\n", w.ID)
				return
			}
			
			fmt.Printf("Worker %d 开始执行任务\n", w.ID)
			if err := job.Execute(); err != nil {
				fmt.Printf("Worker %d 任务执行失败: %v\n", w.ID, err)
			} else {
				fmt.Printf("Worker %d 任务执行完成\n", w.ID)
			}
			
		case <-w.quit:
			fmt.Printf("Worker %d 收到停止信号\n", w.ID)
			return
		}
	}
}

func (w *Worker) Stop() {
	close(w.quit)
}

func (j *SimpleJob) Execute() error {
	// 模拟任务执行
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("执行任务 %d: %s\n", j.ID, j.Data)
	return nil
}

// 扇入扇出函数
func fanOut(input <-chan int, name string) <-chan string {
	output := make(chan string)
	go func() {
		defer close(output)
		for num := range input {
			result := fmt.Sprintf("%s处理了%d", name, num)
			output <- result
		}
	}()
	return output
}

func fanIn(channels ...<-chan string) <-chan string {
	output := make(chan string)
	var wg sync.WaitGroup
	
	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan string) {
			defer wg.Done()
			for result := range c {
				output <- result
			}
		}(ch)
	}
	
	go func() {
		wg.Wait()
		close(output)
	}()
	
	return output
}

// 管道函数
func generateNumbers(start, end int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for i := start; i <= end; i++ {
			output <- i
		}
	}()
	return output
}

func squareNumbers(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for num := range input {
			output <- num * num
		}
	}()
	return output
}

func filterEvenNumbers(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for num := range input {
			if num%2 == 0 {
				output <- num
			}
		}
	}()
	return output
}

// 单例模式实现
type Singleton struct {
	data string
}

var (
	singletonInstance *Singleton
	singletonOnce     sync.Once
)

func GetSingletonInstance() *Singleton {
	singletonOnce.Do(func() {
		singletonInstance = &Singleton{
			data: "单例数据",
		}
	})
	return singletonInstance
}

// 验证函数
func validateStruct(user User) map[string]string {
	errors := make(map[string]string)
	
	if user.Name == "" {
		errors["Name"] = "名称不能为空"
	} else if len(user.Name) < 2 {
		errors["Name"] = "名称长度不能少于2个字符"
	}
	
	if user.Age < 0 || user.Age > 150 {
		errors["Age"] = "年龄必须在0-150之间"
	}
	
	return errors
}

/*
学习要点总结：

1. 并发编程进阶
   - 工作者池模式
   - 扇入扇出模式
   - 管道模式
   - 超时和取消机制

2. 反射机制
   - 类型检查和字段遍历
   - 结构体标签处理
   - 动态方法调用
   - 性能考虑

3. 泛型编程
   - 泛型函数定义
   - 类型约束使用
   - 泛型数据结构
   - 实际应用场景

4. 上下文管理
   - 超时控制
   - 取消传播
   - 值传递
   - 最佳实践

5. 内存管理
   - 内存分配模式
   - 垃圾回收机制
   - 内存泄漏预防
   - 性能监控

6. 性能优化
   - 字符串操作优化
   - 切片预分配
   - 并发优化
   - 基准测试

7. 设计模式
   - 常用设计模式
   - Go语言实现
   - 适用场景
   - 最佳实践

8. 测试策略
   - 单元测试
   - 基准测试
   - 测试覆盖率
   - 测试驱动开发

实践建议：
- 从简单的并发模式开始
- 谨慎使用反射，注意性能
- 合理使用泛型提高代码复用
- 正确使用context管理生命周期
- 关注内存使用和性能优化
- 学习和应用设计模式
- 编写全面的测试用例
*/