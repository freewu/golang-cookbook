# cmp
```
系统默认封装的比较操作相关
```

## Functions 方法
### func Compare 比较 x 和 y 的大小
```
func Compare[T Ordered](x, y T) int

 返回值：
 - -1: x < y
 - 0: x == y
 - 1: x > y
```
示例代码:
```go
package main

import "fmt"
import "cmp"
import "math"

func main() {
    fmt.Println(cmp.Compare(1, 2)) // -1
    fmt.Println(cmp.Compare(1, 1)) // 0
    fmt.Println(cmp.Compare(2, 1)) // 1
    fmt.Println(cmp.Compare("a", "aa")) // -1
    fmt.Println(cmp.Compare(1.5, 1.5)) // 0
    fmt.Println(cmp.Compare(math.NaN(), 1.0)) // 0
}
```

### func Less 判断 x 是否小于 y
```
func Less[T Ordered](x, y T) bool

返回值：
Less reports whether x is less than y. 
For floating-point types, a NaN is considered less than any non-NaN, and -0.0 is not less than (is equal to) 0.0
```
示例代码:
```go
package main

import "fmt"
import "cmp"
import "math"

func main() {
    fmt.Println(cmp.Less(1, 2)) // true
    fmt.Println(cmp.Less(1, 1)) // false
    fmt.Println(cmp.Less(2, 1)) // false
    fmt.Println(cmp.Less("a", "aa")) // true
    fmt.Println(cmp.Less(1.5, 1.5)) // false
    fmt.Println(cmp.Less(math.NaN(), 1.0)) // true
}
```

### func Or 第一个不等于零值的参数
```
func Or[T comparable](vals ...T) T

返回值：
返回第一个不等于零值的参数。如果所有参数都为零，则返回零值
```
示例代码:
```go
package main

import "fmt"
import "cmp"

func main() {
    userInput1 := ""
    userInput2 := "some text"
    fmt.Println(cmp.Or(userInput1, "default")) // default
    fmt.Println(cmp.Or(userInput2, "default")) // some text
    fmt.Println(cmp.Or(userInput1, userInput2, "default")) // some text
    fmt.Println(cmp.Or(0,0,0,0,0,9,0)) // 9
}
```

## Types 类型
### type Ordered 
```
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}
```
示例代码:
```go
package main

import "fmt"
import "cmp"
import "strings"
import "slices"

func main() {
	type Order struct {
		Product  string
		Customer string
		Price    float64
	}
	orders := []Order{
		{"foo", "alice", 1.00},
		{"bar", "bob", 3.00},
		{"baz", "carol", 4.00},
		{"foo", "alice", 2.00},
		{"bar", "carol", 1.00},
		{"foo", "bob", 4.00},
	}
	// Sort by customer first, product second, and last by higher price
	slices.SortFunc(orders, func(a, b Order) int {
		return cmp.Or(
			strings.Compare(a.Customer, b.Customer),
			strings.Compare(a.Product, b.Product),
			cmp.Compare(b.Price, a.Price),
		)
	})
	for _, order := range orders {
		fmt.Printf("%s %s %.2f\n", order.Product, order.Customer, order.Price)
	}
}
```

## 参考资料
```
https://pkg.go.dev/cmp
```