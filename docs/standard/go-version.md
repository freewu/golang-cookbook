# version

> https://pkg.go.dev/go/version

```
版本操作相关方法
```

# Functions 
## func Compare 版本大小比较
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

## 参考资料
- [go/version](https://pkg.go.dev/go/version)