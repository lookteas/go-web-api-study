package main

import (
	"fmt"
	"math"
)

// 演示Go语言的函数定义和使用
func main() {
	fmt.Println("=== Go 语言基础：函数 ===")
	
	// 1. 基本函数调用
	fmt.Println("\n1. 基本函数：")
	greet("Go语言学习者")
	
	// 2. 带返回值的函数
	fmt.Println("\n2. 带返回值的函数：")
	sum := add(10, 20)
	fmt.Printf("10 + 20 = %d\n", sum)
	
	// 3. 多返回值函数
	fmt.Println("\n3. 多返回值函数：")
	quotient, remainder := divide(17, 5)
	fmt.Printf("17 ÷ 5 = %d 余 %d\n", quotient, remainder)
	
	// 4. 命名返回值
	fmt.Println("\n4. 命名返回值：")
	area, perimeter := rectangleStats(5, 3)
	fmt.Printf("矩形(5x3) 面积=%d, 周长=%d\n", area, perimeter)
	
	// 5. 可变参数函数
	fmt.Println("\n5. 可变参数函数：")
	total1 := sumAll(1, 2, 3, 4, 5)
	total2 := sumAll(10, 20)
	fmt.Printf("1+2+3+4+5 = %d\n", total1)
	fmt.Printf("10+20 = %d\n", total2)
	
	// 6. 匿名函数和闭包
	fmt.Println("\n6. 匿名函数和闭包：")
	
	// 匿名函数
	square := func(x int) int {
		return x * x
	}
	fmt.Printf("5的平方 = %d\n", square(5))
	
	// 闭包
	counter := createCounter()
	fmt.Printf("计数器: %d\n", counter())
	fmt.Printf("计数器: %d\n", counter())
	fmt.Printf("计数器: %d\n", counter())
	
	// 7. 高阶函数
	fmt.Println("\n7. 高阶函数：")
	numbers := []int{1, 2, 3, 4, 5}
	
	// 使用函数作为参数
	doubled := mapFunc(numbers, func(x int) int { return x * 2 })
	fmt.Printf("原数组: %v\n", numbers)
	fmt.Printf("翻倍后: %v\n", doubled)
	
	// 8. 递归函数
	fmt.Println("\n8. 递归函数：")
	fmt.Printf("5的阶乘 = %d\n", factorial(5))
	fmt.Printf("斐波那契数列第10项 = %d\n", fibonacci(10))
	
	// 9. defer 语句
	fmt.Println("\n9. defer 语句：")
	deferExample()
}

// 1. 基本函数 - 无返回值
func greet(name string) {
	fmt.Printf("你好, %s!\n", name)
}

// 2. 带返回值的函数
func add(a, b int) int {
	return a + b
}

// 3. 多返回值函数
func divide(a, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

// 4. 命名返回值
func rectangleStats(length, width int) (area, perimeter int) {
	area = length * width
	perimeter = 2 * (length + width)
	return // 自动返回命名的返回值
}

// 5. 可变参数函数
func sumAll(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 6. 闭包示例
func createCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// 7. 高阶函数 - 接受函数作为参数
func mapFunc(slice []int, fn func(int) int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = fn(v)
	}
	return result
}

// 8. 递归函数示例
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 9. defer 示例
func deferExample() {
	fmt.Println("函数开始执行")
	
	defer fmt.Println("这是第一个defer") // 最后执行
	defer fmt.Println("这是第二个defer") // 倒数第二个执行
	defer fmt.Println("这是第三个defer") // 倒数第三个执行
	
	fmt.Println("函数主体执行")
	
	// defer 语句按照后进先出(LIFO)的顺序执行
}

// 10. 函数作为类型
type MathOperation func(int, int) int

func calculate(a, b int, op MathOperation) int {
	return op(a, b)
}

// 可以这样使用：
// multiply := func(x, y int) int { return x * y }
// result := calculate(5, 3, multiply)