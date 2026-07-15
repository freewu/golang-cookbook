# builtin
```
系统默认封装的方法,无需引入包
```

## Constants 常量
```
const (
	true  = 0 == 0 // Untyped bool.
	false = 0 != 0 // Untyped bool.
)
true and false are the two untyped boolean values.

const iota = 0 // Untyped int.

iota is a predeclared identifier representing the untyped integer ordinal number of the current const specification in a (usually parenthesized) const declaration. It is zero-indexed.
```
示例代码:
```go
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
```

## Variables 变量
```
var nil Type // Type must be a pointer, channel, func, interface, map, or slice type
```

## Functions 方法
### func append 追加切片元素
```
func append(slice []Type, elems ...Type) []Type

将 elems 追加到切片 slice 中, 返回新的切片.
```
示例代码:
```go
package main

import "fmt"

func main() {
    // 追加切片元素
    arr := []int{1, 2, 3}
    fmt.Println(arr) // [1 2 3]
    fmt.Println(append(arr, 4, 5)) // [1 2 3 4 5]
    arr1 := []string{"a", "b", "c"}
    fmt.Println(arr1) // [a b c]
    fmt.Println(append(arr1, "bluefrog", "freewu")) // [a b c bluefrog freewu]
}
```

### func cap 获取切片&数组容量
```
func cap(v Type) int

返回切片 v 的容量, 即 v 可以存储的最大元素数量, 超过当前元素数量.

The cap built-in function returns the capacity of v, according to its type:
    Array: the number of elements in v (same as len(v)).
    Pointer to array: the number of elements in *v (same as len(v)).
    Slice: the maximum length the slice can reach when resliced; if v is nil, cap(v) is zero.
    Channel: the channel buffer capacity, in units of elements; if v is nil, cap(v) is zero.
```
示例代码:
```go
package main

import "fmt"

func main() {
    // 获取切片 & 数组容量
    // 数组
    var arr11 [3]int // 方式1：固定长度声明 [0,0,0] 
    arr12 := [3]int{1,2,3}  // 方式2：初始化赋值 
    arr13 := [...]int{10,20,30} // 方式3：自动推导长度
    fmt.Println(cap(arr11)) // 3
    fmt.Println(cap(arr12)) // 3
    fmt.Println(cap(arr13)) // 3
    // 切片
    var slice11 []int // 零值切片：nil len=0 cap=0
    slice12 := []int{1,2,3} // 直接初始化 [1 2 3] len=3 cap=3
    slice13 := make([]int, 2) // make 创建 len=2, cap=2，元素默认0
    slice14 := make([]int, 2, 5)   // make 创建 len=2，cap=5，预留扩容空间
    fmt.Println(slice11, cap(slice11)) // nil 0
    fmt.Println(slice12, cap(slice12)) // [1 2 3] 3
    fmt.Println(slice13, cap(slice13)) // [0 0] 2
    fmt.Println(slice14, cap(slice14)) // [0 0] 5
}
```

### func clear 清空切片或 map
```
func clear[T ~[]Type | ~map[Type]Type1](t T)

清空切片 t 或映射 t, 即将 t 的长度和容量设置为 0.
```
示例代码:
```go
package main

import "fmt"

func main() {
    // 清空切片或 map
    slice21 := []int{1, 2, 3}
    fmt.Println(slice21) // [1 2 3]
    clear(slice21)
    fmt.Println(slice21) // []
    mp21 := make(map[string]int)
    mp21["a"] = 1
    mp21["b"] = 2
    fmt.Println(mp21) // map[a:1 b:2]
    clear(mp21)
    fmt.Println(mp21) // map[]
}
```

### func close 关闭 channel
```
func close(c chan<- Type)
```
示例代码:
```go 
package main

import "fmt"

func producer(ch chan<- int) { // chan<- int 只写通道，匹配close参数
    for i := 1; i <= 3; i++ {
        ch <- i
    }
    // 发送完毕，关闭通道
    close(ch)
    fmt.Println("生产者：通道已关闭")
}

func main() {
    ch := make(chan int, 3)
    go producer(ch)
    // range 读取，通道关闭自动跳出循环
    for val := range ch {
        fmt.Println("收到：", val)
    }
    fmt.Println("读取完成，通道已关闭")
}
```

### func complex 生成一个 ComplexType 变量
```
func complex(r, i FloatType) ComplexType
```
示例代码:
```go
import "fmt"

func main() {
    complex1 := complex(1, 2) // 1+2i
    fmt.Println(complex1) // (1+2i)
}
```

### func copy 复制切片
```
func copy(dst, src []Type) int
返回复制在数量
```
示例代码:
```go
import "fmt"

func main() {
    slice31 := []int{1, 2, 3}
    fmt.Println(slice31) // [1 2 3]
    slice32 := make([]int, 3)
    fmt.Println(copy(slice32, slice31)) // 3 复制切片 slice31 到 slice32
    fmt.Println(slice32) // [1 2 3]
    slice33 := make([]int, 5)
    fmt.Println(copy(slice33, slice31)) // 3 复制切片 slice31 到 slice33
    fmt.Println(slice33) // [1 2 3 0 0]
}
```

### func delete 删除 map 元素
```
func delete(m map[Type]Type1, key Type)
```
示例代码:
```go
import "fmt"

func main() {
    mp41 := map[string]int{"a": 1, "b": 2}
    fmt.Println(mp41) // map[a:1 b:2]
    delete(mp41, "a")
    fmt.Println(mp41) // map[b:2]
    delete(mp41, "c") // 删除一个不存在的
    fmt.Println(mp41) // map[b:2]
}
```

### func imag 获取复数虚数部分的值
```
func imag(c ComplexType) FloatType
```
示例代码:
```go
import "fmt"

func main() {
    complex51 := complex(1, 2) // 1+2i
    fmt.Println(imag(complex51)) // 2
    complex52 := complex(1024, -2) // 1024-2i
    fmt.Println(imag(complex52)) // -2
}
```

### func len 获取切片、数组、字符串、通道的长度
```
func len(v Type) int

Array: the number of elements in v.
Pointer to array: the number of elements in *v (even if v is nil).
Slice, or map: the number of elements in v; if v is nil, len(v) is zero.
String: the number of bytes in v.
Channel: the number of elements queued (unread) in the channel buffer; if v is nil, len(v) is zero.
```
示例代码:
```go
import "fmt"

func main() {
    arr61 := [5]int{1,2,3,4,5}
    fmt.Println(len(arr61)) // 数组 5
    slice61 := []int{1,2,3,4,5,6,7,8}
    fmt.Println(len(slice61)) // 切片 8
    fmt.Println(len("bluefrog")) // 字符串 8
    ch61 := make(chan int, 5)
    fmt.Println(len(ch61)) // channel 0
    ch61 <- 12
    fmt.Println(len(ch61)) // channel 1
    mp61 := map[string]int{"a": 1, "b": 2}
    fmt.Println(len(mp61)) // map 2
}
```

### func make 创建 切片、Map、通道
```
func make(t Type, size ...IntegerType) Type

Slice: 
    The size specifies the length. 
    The capacity of the slice is equal to its length. 
    A second integer argument may be provided to specify a different capacity; it must be no smaller than the length. 
    For example, make([]int, 0, 10) allocates an underlying array of size 10 and returns a slice of length 0 and capacity 10 that is backed by this underlying array.
Map: 
    An empty map is allocated with enough space to hold the specified number of elements. 
    The size may be omitted, in which case a small starting size is allocated.
Channel: 
    The channel's buffer is initialized with the specified buffer capacity. 
    If zero, or the size is omitted, the channel is unbuffered.
```
示例代码:
```go
import "fmt"

func main() {
    slice71 := make([]int, 1, 5)
    fmt.Println(slice71) // [0]
    mp71 := make(map[string]int, 5)
    fmt.Println(mp71) // map[]
    ch71 := make(chan int, 5)
    fmt.Println(ch71) // 0xc000024140
}
```

### func max 取最大入参
```
func max[T cmp.Ordered](x T, y ...T) T
```
示例代码:
```go
import "fmt"

func main() {
    fmt.Println(max(1, 2, 3)) // 3
    fmt.Println(max(1, 2, 3, 4, 5)) // 5
    fmt.Println(max(1, 2, 3, 4, 5, 6)) // 6
}
```

### func min 取最小入参
```
func min[T cmp.Ordered](x T, y ...T) T
```
示例代码:
```go
import "fmt"

func main() {
    fmt.Println(min(1, 2, 3)) // 1
    fmt.Println(min(1, 2, 3, 4, 5)) // 1
    fmt.Println(min(1, 2, 3, 4, 5, 6)) // 1
}
```

### func new 创建结构体实例
```
func new(TypeOrExpr) *Type
```
示例代码:
```go
import "fmt"

type Student struct {
    Name string
    Age int
}

func main() {
    s := new(Student)
    fmt.Println(s) // &{}
    fmt.Println(s.Name) // ""
    fmt.Println(s.Age) // 0
    s.Name = "张三"
    s.Age = 18
    fmt.Println(s) // &{Name:张三 Age:18}
}
```

### func panic 引发 panic 终止程序
```
func panic(v any)
```
示例代码:
```go
import "fmt"

func main() {
    panic("wrong wrong wrong!!!")
    fmt.Println("bluefrog") // bluefrog 永远不会输出
}
```

### func print 打印输出
```
func print(args ...Type)
```
示例代码:
```go
import "fmt"

func main() {
    print("bluefrog ") // bluefrog 
    print("go go go ")
    // 会输出在一行  bluefrog go go go
    // 等价于 fmt.Print
    fmt.Println()
    fmt.Print("bluefrog ")
    fmt.Print("go go go ")
}
```

### func println 打印输出(结束换行)
```
func println(args ...Type)
```
示例代码:
```go
import "fmt"

func main() {
    println("bluefrog ") // bluefrog 
    println("go go go ") // go go go
    // 等价于 fmt.Println
    fmt.Println()
    fmt.Println("bluefrog ")
    fmt.Println("go go go ")
    // 等价于 fmt.Println
    fmt.Println()
    print("bluefrog\r\n") // bluefrog 
    print("go go go \r\n") // go go go
}
```

### func real 获取复数的实部
```
func real(c ComplexType) FloatType
```
示例代码:
```go
import "fmt"

func main() {
    complex81 := complex(1, 2) // 1+2i
    fmt.Println(real(complex51)) // 1
    complex82 := complex(1024, -2) // 1024-2i
    fmt.Println(real(complex52)) // 1024
}
```

### func recover 恢复 panic
```
func recover() any
```
示例代码:
```go
import "fmt"

func handleRequest() {
    // 请求兜底捕获
    defer func() {
        if err := recover(); err != nil {
            // 日志记录堆栈
            fmt.Printf("请求崩溃，err=%v\n", err)
        }
    }()
    // 业务逻辑
    var m map[string]int
    m["key"] = 1 // 空map赋值，触发panic
}

func main() {
    go handleRequest() // 请求崩溃，err=assignment to entry in nil map
    fmt.Println("1111") 
    //select {} // 阻塞不退出
}
```

## Types 类型
```
type ComplexType
type FloatType
type IntegerType
type Type
type Type1
type TypeOrExpr
type any
type bool
type byte
type comparable
type complex64
type complex128
type error
type float32
type float64
type int
type int8
type int16
type int32
type int64
type rune
type string
type uint
type uint8
type uint16
type uint32
type uint64
type uintptr
```
示例代码:
```go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("==== Go 内置类型内存占用大小（64位系统） ====\n")

	// 1. 布尔
	var vBool bool
	fmt.Printf("bool        变量大小: %d 字节\n", unsafe.Sizeof(vBool)) // bool        变量大小: 1 字节

	// 2. 无符号整数
	var vByte byte // alias uint8
	var vUint8 uint8
	var vUint16 uint16
	var vUint32 uint32
	var vUint64 uint64
	var vUint uint
	var vUintptr uintptr
	fmt.Printf("byte(uint8) 变量大小: %d 字节\n", unsafe.Sizeof(vByte)) // byte(uint8) 变量大小: 1 字节
	fmt.Printf("uint8       变量大小: %d 字节\n", unsafe.Sizeof(vUint8)) // uint8       变量大小: 1 字节
	fmt.Printf("uint16      变量大小: %d 字节\n", unsafe.Sizeof(vUint16)) // uint16      变量大小: 2 字节
	fmt.Printf("uint32      变量大小: %d 字节\n", unsafe.Sizeof(vUint32)) // uint32      变量大小: 4 字节
	fmt.Printf("uint64      变量大小: %d 字节\n", unsafe.Sizeof(vUint64)) // uint64      变量大小: 8 字节
	fmt.Printf("uint        变量大小: %d 字节\n", unsafe.Sizeof(vUint)) // uint        变量大小: 8 字节
	fmt.Printf("uintptr     变量大小: %d 字节\n", unsafe.Sizeof(vUintptr)) // uintptr     变量大小: 8 字节

	// 3. 有符号整数
	var vInt8 int8
	var vInt16 int16
	var vInt32 int32
	var vInt64 int64
	var vInt int
	var vRune rune // alias int32
	// int8        变量大小: 1 字节
	fmt.Printf("int8        变量大小: %d 字节\n", unsafe.Sizeof(vInt8)) // int8        变量大小: 1 字节
	fmt.Printf("int16       变量大小: %d 字节\n", unsafe.Sizeof(vInt16)) // int16       变量大小: 2 字节
	fmt.Printf("int32(rune) 变量大小: %d 字节\n", unsafe.Sizeof(vInt32)) // int32(rune) 变量大小: 4 字节
	fmt.Printf("int64       变量大小: %d 字节\n", unsafe.Sizeof(vInt64)) // int64       变量大小: 8 字节
	fmt.Printf("int         变量大小: %d 字节\n", unsafe.Sizeof(vInt)) // int         变量大小: 8 字节
	fmt.Printf("rune        变量大小: %d 字节\n", unsafe.Sizeof(vRune)) // rune        变量大小	: 4 字节

	// 4. 浮点 FloatType
	var vFloat32 float32
	var vFloat64 float64
	fmt.Printf("float32     变量大小: %d 字节\n", unsafe.Sizeof(vFloat32)) // float32     变量大小: 4 字节
	fmt.Printf("float64     变量大小: %d 字节\n", unsafe.Sizeof(vFloat64)) // float64     变量大小: 8 字节

	// 5. 复数 ComplexType
	var vComplex64 complex64
	var vComplex128 complex128
	fmt.Printf("complex64   变量大小: %d 字节\n", unsafe.Sizeof(vComplex64)) // complex64   变量大小: 8 字节
	fmt.Printf("complex128  变量大小: %d 字节\n", unsafe.Sizeof(vComplex128)) // complex128  变量大小: 16 字节

	// 6. string 字符串（底层结构体，存指针+len）
	var vStr string
	fmt.Printf("string      变量大小: %d 字节\n", unsafe.Sizeof(vStr)) // string      变量大小: 16 字节

	// 7. any = interface{} 空接口（2个指针：类型+数据）
	var vAny any
	fmt.Printf("any(interface{}) 变量大小: %d 字节\n", unsafe.Sizeof(vAny)) // any(interface{}) 变量大小: 16 字节

	// 8. error 接口（接口类型和any结构一致）
	var vErr error
	fmt.Printf("error       变量大小: %d 字节\n", unsafe.Sizeof(vErr)) // error       变量大小: 16 字节

	// 9. comparable 注意：不能定义变量，是泛型约束，无法Sizeof
	fmt.Println("\n注：comparable 仅泛型约束，不能声明变量，无内存大小")
	fmt.Println("注：Type/Type1/IntegerType 等是编译器语法标识，非真实类型")
}
```

## 参考资料
- [builtin](https://pkg.go.dev/builtin)