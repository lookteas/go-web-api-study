// Day 03 - 切片 CRUD 练习
// 日期: 2025年x月 (请根据实际日期修改)
// 学习内容: 切片的 创建(增)、读取(查)、更新(改)、删除(删)

package main

import "fmt"

// removeAt 删除指定索引元素（不保证稳定顺序）
func removeAt[T any](s []T, idx int) []T {
    if idx < 0 || idx >= len(s) {
        return s
    }
    return append(s[:idx], s[idx+1:]...)
}

// removeRange 删除区间 [start, end) 的元素
func removeRange[T any](s []T, start, end int) []T {
    if start < 0 {
        start = 0
    }
    if end > len(s) {
        end = len(s)
    }
    if start >= end {
        return s
    }
    return append(s[:start], s[end:]...)
}

func sliceCRUD() {
    fmt.Println("\n[切片 CRUD]")

    // 增: 创建切片
    s := []int{1, 2, 3}
    t := make([]int, 0, 8) // len=0, cap=8
    fmt.Printf("创建 -> s=%v len=%d cap=%d | t=%v len=%d cap=%d\n", s, len(s), cap(s), t, len(t), cap(t))

    // 增: 追加元素
    s = append(s, 4)
    s = append(s, 5, 6)
    fmt.Printf("追加 -> s=%v len=%d cap=%d\n", s, len(s), cap(s))

    // 查: 读取元素与遍历
    fmt.Printf("读取 -> s[2]=%d\n", s[2])
    fmt.Print("遍历 -> ")
    for i, v := range s {
        fmt.Printf("(%d:%d) ", i, v)
    }
    fmt.Println()

    // 改: 更新指定索引
    s[0] = 10
    s[1] *= 10
    fmt.Printf("更新 -> s=%v\n", s)

    // 切片拷贝（不影响原切片）
    cp := make([]int, len(s))
    copy(cp, s)
    cp[0] = -1
    fmt.Printf("拷贝 -> 原=%v | 拷贝=%v\n", s, cp)

    // 删: 删除单个元素（索引 2）
    s = removeAt(s, 2)
    fmt.Printf("删除索引2 -> s=%v\n", s)

    // 删: 删除区间 [1,3)
    s = removeRange(s, 1, 3)
    fmt.Printf("删除区间[1,3) -> s=%v\n", s)

    // 清空切片（保留底层数组容量）
    s = s[:0]
    fmt.Printf("清空 -> s=%v len=%d cap=%d\n", s, len(s), cap(s))

    // 彻底重置（释放引用，让 GC 可回收）
    s = nil
    fmt.Printf("置为nil -> s=%v len=%d cap=%d\n", s, len(s), cap(s))
}

func main() {
    fmt.Println("=== Day 03: 切片 CRUD 练习 ===")
    sliceCRUD()
}
