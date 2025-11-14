package main

import "fmt"

// 演示切片与映射的完整 CRUD
func main() {
	fmt.Println("=== Go 语言基础：切片与映射 CRUD ===")
	sliceCRUD()
	mapCRUD()
}

// ---------- 切片 CRUD ----------
func sliceCRUD() {
	fmt.Println("\n--- 切片 CRUD ---")

	// 1) Create：创建与追加
	nums := []int{1, 2, 3}
	fmt.Println("Create:", nums)
	nums = append(nums, 4, 5)
	fmt.Println("Append:", nums)

	// 2) Read：读取与遍历
	fmt.Println("Read index=2:", nums[2])
	fmt.Print("Range: ")
	for i, v := range nums {
		fmt.Printf("[%d]=%d ", i, v)
	}
	fmt.Println()

	// 3) Update：修改元素
	nums[1] = 20
	fmt.Println("Update index=1:", nums)

	// 4) Delete：删除元素
	// 删除单个元素（索引 2）
	index := 2
	nums = append(nums[:index], nums[index+1:]...)
	fmt.Println("Delete single (index=2):", nums)

	// 删除区间 [1,3)
	start, end := 1, 3
	nums = append(nums[:start], nums[end:]...)
	fmt.Println("Delete range [1,3):", nums)

	// 清空：保留容量 vs 释放引用
	nums = nums[:0] // 仅清空元素，底层数组仍在
	fmt.Println("Clear (len=0):", nums, "cap=", cap(nums))
	nums = nil // 彻底置空
	fmt.Println("Nil:", nums, nums == nil)
}

// ---------- 映射 CRUD ----------
func mapCRUD() {
	fmt.Println("\n--- 映射 CRUD ---")

	// 1) Create：创建与写入
	userAge := map[string]int{"Alice": 30, "Bob": 25}
	fmt.Println("Create:", userAge)
	userAge["Charlie"] = 28 // 新增键值
	fmt.Println("Add:", userAge)

	// 2) Read：读取与“ok”习语
	age, ok := userAge["Alice"]
	if ok {
		fmt.Println("Read Alice age:", age)
	} else {
		fmt.Println("Alice not found")
	}
	_, exist := userAge["Dave"]
	fmt.Println("Dave exist?", exist)

	// 遍历
	fmt.Print("Range: ")
	for k, v := range userAge {
		fmt.Printf("%s:%d ", k, v)
	}
	fmt.Println()

	// 3) Update：修改值
	userAge["Alice"] = 31
	fmt.Println("Update Alice:", userAge)

	// 4) Delete：删除键
	delete(userAge, "Bob")
	fmt.Println("Delete Bob:", userAge)

	// 清空：两种常见做法
	// 方法 A：重新赋值（原 map 被 GC）
	userAge = make(map[string]int)
	fmt.Println("Clear (make new):", userAge)

	// 方法 B：逐个删除（复用原 map）
	userAge["X"] = 1
	userAge["Y"] = 2
	for k := range userAge {
		delete(userAge, k)
	}
	fmt.Println("Clear (delete all):", userAge, "len=", len(userAge))
}