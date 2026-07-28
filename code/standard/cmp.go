package main

import "fmt"
import "cmp"
import "math"
import "strings"
import "slices"

func main() {
    fmt.Println("// == func Compare 比较 x 和 y 的大小 ====================================================================================================")
    fmt.Println(cmp.Compare(1, 2)) // -1
    fmt.Println(cmp.Compare(1, 1)) // 0
    fmt.Println(cmp.Compare(2, 1)) // 1
    fmt.Println(cmp.Compare("a", "aa")) // -1
    fmt.Println(cmp.Compare(1.5, 1.5)) // 0
    fmt.Println(cmp.Compare(math.NaN(), 1.0)) // 0

    fmt.Println("// == func Less 判断 x 是否小于 y ====================================================================================================")
    fmt.Println(cmp.Less(1, 2)) // true
    fmt.Println(cmp.Less(1, 1)) // false
    fmt.Println(cmp.Less(2, 1)) // false
    fmt.Println(cmp.Less("a", "aa")) // true
    fmt.Println(cmp.Less(1.5, 1.5)) // false
    fmt.Println(cmp.Less(math.NaN(), 1.0)) // true

    fmt.Println("// == func Or 第一个不等于零值的参数 ====================================================================================================")
    userInput1 := ""
    userInput2 := "some text"
    fmt.Println(cmp.Or(userInput1, "default")) // default
    fmt.Println(cmp.Or(userInput2, "default")) // some text
    fmt.Println(cmp.Or(userInput1, userInput2, "default")) // some text
    fmt.Println(cmp.Or(0,0,0,0,0,9,0)) // 9

    fmt.Println("// == func Compare for sort 示例 ====================================================================================================")
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
            strings.Compare(a.Customer, b.Customer), // 先按客户排序 a-z
            strings.Compare(a.Product, b.Product), // 再按产品排序 a-z
            //cmp.Compare(b.Price, a.Price), // 最后按价格排序 从高到低
            cmp.Compare(a.Price, b.Price), // 最后按价格排序 从低到高
        )
    })
    for _, order := range orders {
        fmt.Printf("%s %s %.2f\n", order.Product, order.Customer, order.Price)
    }
}