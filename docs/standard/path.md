# path

> https://pkg.go.dev/path

```
目录路径相关
```

## Variables 变量
```
var ErrBadPattern = errors.New("syntax error in pattern") // 系统封装好的路径格式错误类型
```
示例代码:
```go
package main

import "fmt"
import "path"

func main() {
    // Variables 变量
    fmt.Println(path.ErrBadPattern) // syntax error in pattern
    fmt.Println(path.ErrBadPattern.Error()) // syntax error in pattern
    // 等同于 errors.New("syntax error in pattern")
    fmt.Println(errors.New("syntax error in pattern")) // syntax error in pattern
}
```

## Functions 方法
### func Base 返回路径的文件名
```
func Base(path string) string

返回值：
Base returns the last element of path. Trailing slashes are removed before extracting the last element. 
If the path is empty, Base returns ".". 
If the path consists entirely of slashes, Base returns "/".
```
示例代码:
```go
package main

import "fmt"
import "path"

func main() {
    fmt.Println(path.Base("/a/b")) // b
    fmt.Println(path.Base("/a/b/")) // b
    fmt.Println(path.Base("/a/b/c.txt")) // c.txt
    fmt.Println(path.Base("/")) // / 为 "/" 返回 /
    fmt.Println(path.Base("")) // . 为空返回 .
}
```

### func Clean 清理路径中的冗余斜杠
```
func Clean(path string) string

返回值：
Clean returns the cleaned path, with redundant slashes removed.
```
示例代码:
```go
package main

import "fmt"
import "path"

func main() {
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
}
```

### func Dir 返回路径的目录名
```
func Dir(path string) string

返回值：
Dir returns the directory part of path, with trailing slash if it is a directory.
If the path is empty, Dir returns ".".
```
示例代码:
```go
package main

import "fmt"
import "path"

func main() {
    fmt.Println(path.Dir("/a/b/c")) // /a/b
    fmt.Println(path.Dir("a/b/c")) // a/b
    fmt.Println(path.Dir("/a/")) // /a
    fmt.Println(path.Dir("a/")) // a
    fmt.Println(path.Dir("/")) // /
    fmt.Println(path.Dir("")) // .
}
```

### func Ext 返回路径的文件扩展名
```
func Ext(path string) string

返回值：
Ext returns the file extension of path. If the path has no extension, Ext returns the empty string.
```
示例代码:
```go
package main

import "fmt"
import "path"

func main() {
    fmt.Println(path.Ext("/a/b/c.txt")) // .txt
    fmt.Println(path.Ext("/a/b/c/bar.css")) // .css
    fmt.Println(path.Ext("/")) // ""
    fmt.Println(path.Ext("")) // ""
}
```

### func IsAbs 判断路径是否为绝对路径
```
func IsAbs(path string) bool

返回值：
IsAbs reports whether the path is absolute.
```
示例代码:
```go
package main

import "fmt"
import "path"

func main() {
    fmt.Println(path.IsAbs("/dev/null")) // true
    fmt.Println(path.IsAbs("/a/b/b/c")) // true
    fmt.Println(path.IsAbs("a/b/c")) // false
    fmt.Println(path.IsAbs("/")) // true
    fmt.Println(path.IsAbs("")) // false
    fmt.Println(path.IsAbs("../a/b/b/c")) // false
}
```

### func Join 合并路径
```
func Join(path1, path2 string) string

返回值：
Join returns the path1 and path2 joined together, with a single slash separating them.
```
示例代码:
```go
package main

import "fmt"
import "path"

func main() {
    fmt.Println(path.Join("a", "b", "c"))
    fmt.Println(path.Join("a", "b/c"))
    fmt.Println(path.Join("a/b", "c"))
    fmt.Println(path.Join("a/b", "../../../xyz"))
    fmt.Println(path.Join("", ""))
    fmt.Println(path.Join("a", ""))
    fmt.Println(path.Join("", "a"))
}
```

### func Match 匹配路径
```
func Match(pattern, name string) (matched bool, err error)

Match reports whether name matches the shell pattern. The pattern syntax is:

pattern:
	{ term }
term:
	'*'         matches any sequence of non-/ characters
	'?'         matches any single non-/ character
	'[' [ '^' ] { character-range } ']'
	            character class (must be non-empty)
	c           matches character c (c != '*', '?', '\\', '[')
	'\\' c      matches character c

character-range:
	c           matches character c (c != '\\', '-', ']')
	'\\' c      matches character c
	lo '-' hi   matches character c for lo <= c <= hi
```
示例代码:
```go
package main

import "fmt"
import "path"

func main() {
    fmt.Println(path.Match("abc", "abc")) // true <nil>
    fmt.Println(path.Match("a*", "abc")) // true <nil>
    fmt.Println(path.Match("a*/b", "a/c/b")) // false <nil>
    fmt.Println(path.Match("[a-z]", "a/c/b")) // false syntax error in pattern 正则表达式错误
}
```

### func Split 分割路径
```
func Split(path string) (dir, file string)
```
示例代码:
```go
package main

import "fmt"
import "path"

func main() {
    fmt.Println(path.Match("abc", "abc")) // true <nil>
    fmt.Println(path.Match("a*", "abc")) // true <nil>
    fmt.Println(path.Match("a*/b", "a/c/b")) // false <nil>
    fmt.Println(path.Match("[a-z]", "a/c/b")) // false syntax error in pattern 正则表达式错误
}
```

## 参考资料
```
https://pkg.go.dev/path
```