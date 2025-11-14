// Day 02 - 变量与结构体 CRUD 练习
// 日期: 2025年x月 (请根据实际日期修改)
// 学习内容: 变量与结构体的 创建(增)、读取(查)、更新(改)、置空(“删”)

package main

import "fmt"

// Person 用于结构体 CRUD 演示
type Person struct {
    Name string
    Age  int
    Job  string
}

func variableCRUD() {
    fmt.Println("\n[变量 CRUD]")

    // 增: 创建变量（声明 + 赋值）
    var name string = "Go 学习者"
    var age int = 25
    isStudent := false // 短变量声明
    fmt.Printf("创建 -> 姓名:%s 年龄:%d 学生:%v\n", name, age, isStudent)

    // 查: 读取变量
    fmt.Printf("读取 -> name=%q age=%d isStudent=%v\n", name, age, isStudent)

    // 改: 更新变量（重新赋值）
    name = "Gopher"
    age += 1
    isStudent = true
    fmt.Printf("更新 -> 姓名:%s 年龄:%d 学生:%v\n", name, age, isStudent)

    // “删”: 将变量置为零值/空值（Go 没有真正的删除变量语义）
    name = ""     // 字符串零值
    age = 0        // 数值零值
    isStudent = false
    fmt.Printf("置空 -> 姓名:%q 年龄:%d 学生:%v\n", name, age, isStudent)
}

func structCRUD() {
    fmt.Println("\n[结构体 CRUD]")

    // 增: 创建结构体实例
    p := Person{Name: "Alice", Age: 30, Job: "Engineer"}
    fmt.Printf("创建 -> %+v\n", p)

    // 查: 读取字段
    fmt.Printf("读取 -> Name=%s Age=%d Job=%s\n", p.Name, p.Age, p.Job)

    // 改: 更新字段
    p.Age++
    p.Job = "Senior Engineer"
    fmt.Printf("更新 -> %+v\n", p)

    // “删”: 将结构体置为零值（或指针置为 nil）
    p = Person{} // 所有字段置为零值
    fmt.Printf("置空 -> %+v\n", p)

    // 指针形式的“删”演示
    pp := &Person{Name: "Bob", Age: 28, Job: "Designer"}
    fmt.Printf("指针创建 -> %+v\n", *pp)
    pp = nil
    fmt.Printf("指针置空 -> %v\n", pp)
}

func numberTypes() {
    fmt.Println("\n[数值类型与转换]")
    var i8 int8 = 127
    var f32 float32 = 3.1415927
    var f64 float64 = 3.141592653589793
    fmt.Printf("int8=%d f32=%.7f f64=%.15f\n", i8, f32, f64)

    // 类型转换演示
    i32 := int32(i8)
    f := float64(i32)
    fmt.Printf("转换 -> int32=%d float64=%.0f\n", i32, f)
}

func main() {
    fmt.Println("=== Day 02: 变量与结构体 CRUD 练习 ===")

    variableCRUD()
    structCRUD()
    numberTypes()
}
