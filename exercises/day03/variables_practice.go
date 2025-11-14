// Day 02 - 变量和类型练习
// 日期: 2025年x月 (请根据实际日期修改)
// 学习内容: 变量声明、基本类型、类型转换

package main

import "fmt"

func main() {
	fmt.Println("=== Day 02: 变量和类型练习 ===")

	// TODO: 在这里完成今天的练习
	// 1. 声明不同类型的变量
	var name string = "go 语言爱好者"
	var age int = 25
	var isStudnet bool = false
	fmt.Printf("姓名：%s, 年龄：%d, 是否学生：%v\n", name, age, isStudnet)

	//短变量声明（只能在函数内使用）
	day := 2
	year := 2025
	version := "v1.0.0"
	fmt.Printf("今天是第 %d 天练习，%d年, 版本号为:%s\n", day, year, version)

	//基本数据类型
	var int8Val int8 = 127

	// 浮点类型
	var float32Val float32 = 3.141592653589793
	var float64Val float64 = 3.141592653589793

	fmt.Printf("整数类型： %d\n", int8Val)
	fmt.Printf("float32浮点类型： %f\n", float32Val)
	fmt.Printf("float64浮点类型： %f\n", float64Val)

	// 2. 练习类型转换
	// 3. 使用常量
	// 4. 数组和切片操作

	fmt.Println("今天的练习: 变量和类型")
	fmt.Println("请完成上面的TODO项目")
}
