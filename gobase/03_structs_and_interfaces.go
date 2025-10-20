package main

import (
	"fmt"
	"math"
)

// 演示Go语言的结构体和接口
func main() {
	fmt.Println("=== Go 语言基础：结构体和接口 ===")
	
	// 1. 基本结构体
	fmt.Println("\n1. 基本结构体：")
	person1 := Person{
		Name: "张三",
		Age:  25,
		City: "北京",
	}
	fmt.Printf("人员信息: %+v\n", person1)
	
	// 2. 结构体方法
	fmt.Println("\n2. 结构体方法：")
	person1.Introduce()
	person1.HaveBirthday()
	person1.Introduce()
	
	// 3. 结构体嵌入（组合）
	fmt.Println("\n3. 结构体嵌入：")
	emp := Employee{
		Person: Person{
			Name: "李四",
			Age:  30,
			City: "上海",
		},
		JobTitle: "软件工程师",
		Salary:   15000,
	}
	emp.Introduce() // 调用嵌入的Person的方法
	emp.Work()
	
	// 4. 接口使用
	fmt.Println("\n4. 接口使用：")
	
	// 创建不同的形状
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}
	triangle := Triangle{Base: 3, Height: 4}
	
	// 使用接口
	shapes := []Shape{circle, rectangle, triangle}
	
	for i, shape := range shapes {
		fmt.Printf("形状 %d:\n", i+1)
		fmt.Printf("  面积: %.2f\n", shape.Area())
		fmt.Printf("  周长: %.2f\n", shape.Perimeter())
		
		// 类型断言
		if c, ok := shape.(Circle); ok {
			fmt.Printf("  这是一个圆，半径: %.2f\n", c.Radius)
		}
		fmt.Println()
	}
	
	// 5. 空接口
	fmt.Println("5. 空接口：")
	var anything interface{}
	
	anything = 42
	fmt.Printf("存储整数: %v (类型: %T)\n", anything, anything)
	
	anything = "Hello"
	fmt.Printf("存储字符串: %v (类型: %T)\n", anything, anything)
	
	anything = person1
	fmt.Printf("存储结构体: %v (类型: %T)\n", anything, anything)
	
	// 6. 接口组合
	fmt.Println("\n6. 接口组合：")
	var rw ReadWriter = &File{Name: "test.txt"}
	rw.Read()
	rw.Write("Hello, Go!")
	
	// 7. 方法集
	fmt.Println("\n7. 方法集（值接收者 vs 指针接收者）：")
	
	// 值接收者
	var counter1 Counter = ValueCounter{count: 0}
	counter1.Increment()
	fmt.Printf("值接收者计数: %d\n", counter1.GetCount())
	
	// 指针接收者
	var counter2 Counter = &PointerCounter{count: 0}
	counter2.Increment()
	fmt.Printf("指针接收者计数: %d\n", counter2.GetCount())
}

// 1. 基本结构体定义
type Person struct {
	Name string
	Age  int
	City string
}

// 结构体方法 - 值接收者
func (p Person) Introduce() {
	fmt.Printf("大家好，我是%s，今年%d岁，来自%s\n", p.Name, p.Age, p.City)
}

// 结构体方法 - 指针接收者（可以修改结构体）
func (p *Person) HaveBirthday() {
	p.Age++
	fmt.Printf("%s 过生日了，现在%d岁\n", p.Name, p.Age)
}

// 2. 结构体嵌入示例
type Employee struct {
	Person   // 嵌入Person结构体
	JobTitle string
	Salary   float64
}

// Employee的方法
func (e Employee) Work() {
	fmt.Printf("%s 正在工作，职位是%s\n", e.Name, e.JobTitle)
}

// 3. 接口定义
type Shape interface {
	Area() float64
	Perimeter() float64
}

// 实现Shape接口的结构体

// 圆形
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 矩形
type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 三角形
type Triangle struct {
	Base, Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

func (t Triangle) Perimeter() float64 {
	// 假设是等腰三角形
	side := math.Sqrt((t.Base/2)*(t.Base/2) + t.Height*t.Height)
	return t.Base + 2*side
}

// 4. 接口组合
type Reader interface {
	Read() string
}

type Writer interface {
	Write(data string)
}

// 组合接口
type ReadWriter interface {
	Reader
	Writer
}

// 实现ReadWriter接口
type File struct {
	Name string
}

func (f *File) Read() string {
	fmt.Printf("从文件 %s 读取数据\n", f.Name)
	return "file content"
}

func (f *File) Write(data string) {
	fmt.Printf("向文件 %s 写入数据: %s\n", f.Name, data)
}

// 5. 方法集示例
type Counter interface {
	Increment()
	GetCount() int
}

// 值接收者实现
type ValueCounter struct {
	count int
}

func (vc ValueCounter) Increment() {
	vc.count++ // 这不会修改原始值
}

func (vc ValueCounter) GetCount() int {
	return vc.count
}

// 指针接收者实现
type PointerCounter struct {
	count int
}

func (pc *PointerCounter) Increment() {
	pc.count++ // 这会修改原始值
}

func (pc *PointerCounter) GetCount() int {
	return pc.count
}