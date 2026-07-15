package main

import "fmt"

// iota 是 Go 预定义标识符，仅在 const 常量块内生效
// 普通自增枚举
const (
    Status1 = iota // 0
    Status2        // 1
    Status3        // 2
)

// 同一行多个常量，iota 相同
const (
    A, B = iota, iota // iota=0 → A=0, B=0
    C, D              // iota=1 → C=1, D=1
)

// 自定义偏移量
const (
    Offset1 = iota + 10 // 0+10=10
    Offset2             // 1+10=11
    Offset3             // 2+10=12
)

// 位运算（权限掩码，经典用法）
const (
    Read = 1 << iota  // 1 << 0 = 0b001 = 1
    Write             // 1 << 1 = 0b010 = 2
    Execute           // 1 << 2 = 0b100 = 4
)

// 跳过值（使用 _ 占位）
const (
    _ = iota // 0
    Apple    // 1
    Banana   // 2
    Orange   // 3
)

// 复杂表达式复用
const (
    Double = iota * 2 // 0
    TwoX              // 2
    FourX             // 4
    SixX              // 6
)

// 手动赋值会打断自动递增
const (
    A1 = iota // 0
    B1 = 100  // 手动赋值，不再继承iota
    C1        // 复用 B1 的值 = 100，不是2
)

func main() {
    // iota
    // 普通自增枚举
    fmt.Println(Status1, Status2, Status3) // 0 1 2
    // 同一行多个常量，iota 相同
    fmt.Println(A, B, C, D) // 0 0 1 1
    // 自定义偏移量
    fmt.Println(Offset1, Offset2, Offset3) // 10 11 12
    // 位运算（权限掩码，经典用法）
    fmt.Println(Read, Write, Execute) // 1 2 4
    // 跳过值（使用 _ 占位）
    fmt.Println(Apple, Banana, Orange) // 1 2 3
    // 复杂表达式复用
    fmt.Println(Double, TwoX, FourX, SixX) // 0 2 4 6
    // 手动赋值会打断自动递增
    fmt.Println(A1, B1, C1) // 0 100 100

    // 打印常量 true 和 false
    fmt.Println(true) // true
    fmt.Println(false) // false
}