package main

import "fmt"
import "go/version"

func main() {
	// 版本大小比较
    fmt.Println(version.Compare("go1.21", "go1.20")) // 1
    fmt.Println(version.Compare("go1.20", "go1.20")) // 0
    fmt.Println(version.Compare("go1.20", "go1.21")) // -1

	// 验证版本是否有效	
	fmt.Println(version.IsValid("go1.20")) // true
    fmt.Println(version.IsValid("go1.20.0")) // true
    fmt.Println(version.IsValid("go1.20.0.1")) // false
	fmt.Println(version.IsValid("go1.21rc1")) // true	
	fmt.Println(version.IsValid("go1.21xxyyzz1")) // true	

	// 返回版本 x 的大版本
	fmt.Println(version.Lang("go1.21rc2")) // go1.21
    fmt.Println(version.Lang("go1.21.2")) // "go1.21"
    fmt.Println(version.Lang("go1.21")) // "go1.21"
	fmt.Println(version.Lang("bad")) // ""
    fmt.Println(version.Lang("1.21")) // ""
    fmt.Println(version.Lang("go1.20.0.1")) // ""
    fmt.Println(version.Lang("go1")) // "go1"

}