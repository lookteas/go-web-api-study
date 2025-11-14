// Day 02 - 变量和类型练习
// 日期: 2025年x月 (请根据实际日期修改)
// 学习内容: 变量声明、基本类型、类型转换

package main

import "fmt"

func main() {

	day := 3
	year := 2025
	excName := "变量和类型练习"
	version := "v1.0.0"
	fmt.Printf("=== Day 0%d: %s ===\n", day, excName)
	//短变量声明（只能在函数内使用）
	fmt.Printf("今天是第 %d 天练习，%d年, 版本号为:%s\n", day, year, version)

	// 浮点类型
	var float32Val float32 = 3.141592653589793
	var float64Val float64 = 3.141592653589793

	fmt.Printf("float32浮点类型： %f\n", float32Val)
	fmt.Printf("float64浮点类型： %f\n", float64Val)

	//布尔类型
	var boolVar bool = false
	fmt.Printf("布尔类型：boolVar：%t\n", boolVar)

	//数组
	var numbers [5]int = [5]int{1, 2, 3, 5, 7}
	fmt.Printf("数组变量： numbers : %v\n", numbers)

	var Strings [4]string = [4]string{"one", "two", "three", "four"}
	fmt.Printf("字符串数组 Strings: %v\n", Strings)

	//切片
	var Stringv []string = []string{"one", "two", "three", "four"}
	fmt.Printf("切片字符类型 Stringv: %v\n", Stringv)

	var intdocker []int = []int{1, 2, 3}
	fmt.Printf("切片数字类型 intdocker: %v\n", intdocker)

	var boolDocker []bool = []bool{true, false, true}
	fmt.Printf("切片布尔类型 boolDocker: %v\n", boolDocker)

	//映射
	var mapDocker map[string]int = map[string]int{"one": 1, "two": 222, "three": 33}
	fmt.Printf("映射类型 mapDocker: %v\n", mapDocker)

	var mapIntDocker map[int]int = map[int]int{111: 1, 222: 2, 3333: 3}
	fmt.Printf("映射整数类型 mapIntDocker: %v\n", mapIntDocker)

	fmt.Println(Stringv[2], intdocker[2], mapDocker["two"], mapIntDocker[222])

}
