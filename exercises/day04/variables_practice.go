// Day 04 - Map CRUD 练习
// 日期: 2025年x月 (请根据实际日期修改)
// 学习内容: 映射(Map)的 创建(增)、读取(查)、更新(改)、删除(删)

package main

import "fmt"

func mapCRUD() {
    fmt.Println("\n[Map CRUD]")

    // 增: 创建 map
    m := map[string]int{"one": 1, "two": 2}
    n := make(map[int]string)
    n[2009] = "Go"
    n[1991] = "Python"
    fmt.Printf("创建 -> m=%v | n=%v\n", m, n)

    // 增: 新增键值对
    m["three"] = 3
    m["four"] = 4
    fmt.Printf("新增 -> m=%v\n", m)

    // 查: 读取（含 ok 习语）
    val := m["two"]
    fmt.Printf("读取 -> m[\"two\"]= %d\n", val)
    if v, ok := m["five"]; ok {
        fmt.Printf("读取存在 -> five=%d\n", v)
    } else {
        fmt.Println("读取不存在 -> five 不在 map 中")
    }

    // 改: 更新值
    m["two"] += 100
    fmt.Printf("更新 -> m=%v\n", m)

    // 删: 删除键
    delete(m, "three")
    fmt.Printf("删除键 three -> m=%v\n", m)

    // 遍历
    fmt.Print("遍历 -> ")
    for k, v := range m {
        fmt.Printf("(%s:%d) ", k, v)
    }
    fmt.Println()

    // 清空 map（两种方式）
    // 方式一：重新赋值为新的空 map
    m = map[string]int{}
    fmt.Printf("清空(重建) -> m=%v\n", m)

    // 方式二：逐个删除键（保留引用）
    tmp := map[string]int{"A": 1, "B": 2, "C": 3}
    for k := range tmp {
        delete(tmp, k)
    }
    fmt.Printf("清空(遍历删除) -> tmp=%v\n", tmp)
}

func main() {
    fmt.Println("=== Day 04: Map CRUD 练习 ===")
    mapCRUD()
}
