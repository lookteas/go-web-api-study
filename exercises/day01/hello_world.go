// Day 01 - Hello World 练习
// 日期: 2024年1月 (请根据实际日期修改)
// 学习内容: Go语言基础语法入门

package main

import "fmt"

func main() {
	// 基础输出
	fmt.Println("=== Day 01: Hello World 练习 ===")
	fmt.Println("Hello, World!")
	fmt.Println("你好，Go语言！")
	
	// 变量练习
	name := "Go学习者"
	day := 1
	fmt.Printf("我是%s，今天是第%d天学习Go语言\n", name, day)
	
	// 简单计算
	a, b := 10, 20
	sum := a + b
	fmt.Printf("%d + %d = %d\n", a, b, sum)
	
	fmt.Println("今天的学习完成！✅")
}