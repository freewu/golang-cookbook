package main

import "fmt"
import "go/version"

func main() {
    fmt.Println(version.Compare("go1.21", "go1.20")) // 1
    fmt.Println(version.Compare("go1.20", "go1.20")) // 0
    fmt.Println(version.Compare("go1.20", "go1.21")) // -1
}