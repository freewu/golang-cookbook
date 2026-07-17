package main

import "fmt"
import "path"
import "errors"

func main() {
    // Variables 变量
    fmt.Println(path.ErrBadPattern) // syntax error in pattern
    fmt.Println(path.ErrBadPattern.Error()) // syntax error in pattern
    // 等同于 errors.New("syntax error in pattern")
    fmt.Println(errors.New("syntax error in pattern")) // syntax error in pattern

    // == func Base 返回路径的文件名 ====================================================================================================
    fmt.Println(path.Base("/a/b")) // b
    fmt.Println(path.Base("/a/b/")) // b
    fmt.Println(path.Base("/a/b/c.txt")) // c.txt
    fmt.Println(path.Base("/")) // / 为 "/" 返回 /
    fmt.Println(path.Base("")) // . 为空返回 .

    // == func Clean 清理路径中的冗余斜杠 ====================================================================================================
    paths := []string{
        "a/c",               // Clean("a/c") = "a/c"
        "a//c",              // Clean("a//c") = "a/c"
        "a/c/.",             // Clean("a/c/.") = "a/c"
        "a/c/b/..",          // Clean("a/c/b/..") = "a/c"
        "../a/c",            // Clean("../a/c") = "../a/c"
        "../a/b/../././/c",  // Clean("../a/b/../././/c") = "../a/c"
        "/../a/c",           // Clean("/../a/c") = "/a/c"
        "/../a/b/../././/c", // Clean("/../a/b/../././/c") = "/a/c"
        "",                  // Clean("") = "."  
    }
    for _, p := range paths {
        fmt.Printf("Clean(%q) = %q\n", p, path.Clean(p))
    }

    // == func Dir 返回路径的目录名 ====================================================================================================
    fmt.Println(path.Dir("/a/b/c")) // /a/b
    fmt.Println(path.Dir("a/b/c")) // a/b
    fmt.Println(path.Dir("/a/")) // /a
    fmt.Println(path.Dir("a/")) // a
    fmt.Println(path.Dir("/")) // /
    fmt.Println(path.Dir("")) // .

    // == func Ext 返回路径的文件扩展名 ====================================================================================================
    fmt.Println(path.Ext("/a/b/c.txt")) // .txt
    fmt.Println(path.Ext("/a/b/c/bar.css")) // .css
    fmt.Println(path.Ext("/")) // ""
    
    // == func IsAbs 判断路径是否为绝对路径 ====================================================================================================
    fmt.Println(path.IsAbs("/dev/null")) // true
    fmt.Println(path.IsAbs("/a/b/b/c")) // true
    fmt.Println(path.IsAbs("a/b/c")) // false
    fmt.Println(path.IsAbs("/")) // true
    fmt.Println(path.IsAbs("")) // false
    fmt.Println(path.IsAbs("../a/b/b/c")) // false

    // == func Join 合并路径 ====================================================================================================   
    fmt.Println(path.Join("a", "b", "c")) // a/b/c
	fmt.Println(path.Join("a", "b/c")) // a/b/c
	fmt.Println(path.Join("a/b", "c")) // a/b/c
	fmt.Println(path.Join("a/b", "../../../xyz")) // xyz
	fmt.Println(path.Join("", "")) // ""
	fmt.Println(path.Join("a", "")) // a
	fmt.Println(path.Join("", "a")) // a

    // == func Match 匹配路径 ====================================================================================================
    fmt.Println(path.Match("abc", "abc")) // true <nil>
    fmt.Println(path.Match("a*", "abc")) // true <nil>
    fmt.Println(path.Match("a*/b", "a/c/b")) // false <nil>
    fmt.Println(path.Match("[a-z]", "a/c/b")) // false syntax error in pattern 正则表达式错误

    // == func Split 分割路径 ====================================================================================================
    split := func(s string) {
		dir, file := path.Split(s)
		fmt.Printf("path.Split(%q) = dir: %q, file: %q\n", s, dir, file)
	}
	split("static/myfile.css") // path.Split("static/myfile.css") = dir: "static/", file: "myfile.css"
	split("myfile.css") // path.Split("myfile.css") = dir: "./", file: "myfile.css"
	split("") // path.Split("") = dir: "", file: ""
    split("../../../") // path.Split("../../../") = dir: "../../../", file: ""
    split("/../../ab/c") // path.Split("/../../ab/c") = dir: "/../../ab/", file: "c"
}
