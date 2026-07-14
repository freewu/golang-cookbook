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
    fmt.Println(version.Compare("go1.21", "go1.20")) // 1
    fmt.Println(version.Compare("go1.20", "go1.20")) // 0
    fmt.Println(version.Compare("go1.20", "go1.21")) // -1
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

import "fmt"

func main() {
    fmt.Println(append([]int{1, 2, 3}, 4, 5)) // [1 2 3 4 5]
}
```

## 参考资料
- [builtin](https://pkg.go.dev/builtin)