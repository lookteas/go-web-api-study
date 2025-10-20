package main

import (
	"fmt"
	"reflect"
)

// 演示Go语言的变量声明和基本类型
func main() {
	fmt.Println("=== Go 语言基础：变量和类型 ===")
	
	// 1. 变量声明的几种方式
	fmt.Println("\n1. 变量声明方式：")
	
	// 方式1：var 关键字声明
	var name string = "Go语言"
	var age int = 15
	fmt.Printf("var声明: name=%s, age=%d\n", name, age)
	
	// 方式2：类型推断
	var language = "Golang"
	var year = 2009
	fmt.Printf("类型推断: language=%s, year=%d\n", language, year)
	
	// 方式3：短变量声明（只能在函数内使用）
	version := "1.24"
	isPopular := true
	fmt.Printf("短声明: version=%s, isPopular=%t\n", version, isPopular)
	
	// 2. 基本数据类型
	fmt.Println("\n2. 基本数据类型：")
	
	// 整数类型
	var int8Val int8 = 127
	var int16Val int16 = 32767
	var int32Val int32 = 2147483647
	var int64Val int64 = 9223372036854775807
	var uintVal uint = 42
	
	fmt.Printf("整数类型: int8=%d, int16=%d, int32=%d, int64=%d, uint=%d\n", 
		int8Val, int16Val, int32Val, int64Val, uintVal)
	
	// 浮点类型
	var float32Val float32 = 3.14
	var float64Val float64 = 3.141592653589793
	fmt.Printf("浮点类型: float32=%.2f, float64=%.15f\n", float32Val, float64Val)
	
	// 布尔类型
	var boolVal bool = true
	fmt.Printf("布尔类型: %t\n", boolVal)
	
	// 字符串类型
	var stringVal string = "Hello, 世界!"
	fmt.Printf("字符串类型: %s\n", stringVal)
	
	// 3. 复合类型
	fmt.Println("\n3. 复合类型：")
	
	// 数组
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("数组: %v\n", numbers)
	
	// 切片
	var slice []string = []string{"Go", "Python", "Java"}
	fmt.Printf("切片: %v\n", slice)
	
	// 映射
	var languages map[string]int = map[string]int{
		"Go":     2009,
		"Python": 1991,
		"Java":   1995,
	}
	fmt.Printf("映射: %v\n", languages)
	
	// 4. 零值
	fmt.Println("\n4. 零值（变量的默认值）：")
	var zeroInt int
	var zeroFloat float64
	var zeroBool bool
	var zeroString string
	var zeroSlice []int
	var zeroMap map[string]int
	
	fmt.Printf("int零值: %d\n", zeroInt)
	fmt.Printf("float64零值: %f\n", zeroFloat)
	fmt.Printf("bool零值: %t\n", zeroBool)
	fmt.Printf("string零值: '%s'\n", zeroString)
	fmt.Printf("slice零值: %v (nil: %t)\n", zeroSlice, zeroSlice == nil)
	fmt.Printf("map零值: %v (nil: %t)\n", zeroMap, zeroMap == nil)
	
	// 5. 类型转换
	fmt.Println("\n5. 类型转换：")
	var intNum int = 42
	var floatNum float64 = float64(intNum)
	var stringNum string = fmt.Sprintf("%d", intNum)
	
	fmt.Printf("原始int: %d\n", intNum)
	fmt.Printf("转换为float64: %f\n", floatNum)
	fmt.Printf("转换为string: %s\n", stringNum)
	
	// 6. 常量
	fmt.Println("\n6. 常量：")
	const PI = 3.14159
	const GREETING = "Hello, Go!"
	
	fmt.Printf("常量PI: %f\n", PI)
	fmt.Printf("常量GREETING: %s\n", GREETING)
	
	// 7. 使用reflect包查看类型信息
	fmt.Println("\n7. 类型信息：")
	fmt.Printf("name的类型: %s\n", reflect.TypeOf(name))
	fmt.Printf("age的类型: %s\n", reflect.TypeOf(age))
	fmt.Printf("slice的类型: %s\n", reflect.TypeOf(slice))
	fmt.Printf("languages的类型: %s\n", reflect.TypeOf(languages))
}