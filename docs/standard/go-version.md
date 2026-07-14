# version

> https://pkg.go.dev/go/version

```
版本操作相关方法
```

## Functions 方法
### func Compare 版本大小比较
```
func Compare(x, y string) int

 返回值：
 - -1: x < y
 - 0: x == y
 - 1: x > y
```
示例代码:
```go
package main

import "fmt"
import "go/version"

func main() {
    fmt.Println(version.Compare("go1.21", "go1.20")) // 1
    fmt.Println(version.Compare("go1.20", "go1.20")) // 0
    fmt.Println(version.Compare("go1.20", "go1.21")) // -1
}
```

### func IsValid 验证版本是否有效
```
func IsValid(v string) bool

 返回值：
 - true: v 是一个有效的版本字符串
 - false: v 不是一个有效的版本字符串
```
示例代码:
```go
package main

import "fmt"
import "go/version"

func main() {
    fmt.Println(version.IsValid("go1.20")) // true
    fmt.Println(version.IsValid("go1.20.0")) // true
    fmt.Println(version.IsValid("go1.20.0.1")) // false
	fmt.Println(version.IsValid("go1.21rc1")) // true	
    fmt.Println(version.IsValid("go1.21xxyyzz1")) // true	
}
```

### func Lang 返回版本 x 的大版本
```
func Lang(x string) string

 返回值：
 - 返回 版本 x 的大版本
 - 如果 x 不是一个有效的版本字符串, 则返回空字符串.

```
示例代码:
```go
package main

import "fmt"
import "go/version"

func main() {
    fmt.Println(version.Lang("go1.21rc2")) // go1.21
    fmt.Println(version.Lang("go1.21.2")) // "go1.21"
    fmt.Println(version.Lang("go1.21")) // "go1.21"
    fmt.Println(version.Lang("go1")) // "go1"
    fmt.Println(version.Lang("bad")) // ""
    fmt.Println(version.Lang("1.21")) // ""
    fmt.Println(version.Lang("go1.20.0.1")) // ""
}
```

## 参考资料
- [go/version](https://pkg.go.dev/go/version)