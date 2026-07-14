# builtin
```
系统默认封装的方法,无需引入包
```

# Constants 常量
```
const (
	true  = 0 == 0 // Untyped bool.
	false = 0 != 0 // Untyped bool.
)
true and false are the two untyped boolean values.

const iota = 0 // Untyped int.

iota is a predeclared identifier representing the untyped integer ordinal number of the current const specification in a (usually parenthesized) const declaration. It is zero-indexed.
```

# Variables 变量
```
var nil Type // Type must be a pointer, channel, func, interface, map, or slice type
```

# Functions 方法
## func append 追加切片元素
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
