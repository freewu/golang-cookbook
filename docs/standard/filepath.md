# filepath

> https://pkg.go.dev/path/filepath

```
filepath 包提供文件路径操作的函数
```

## Constants 常量
```
const (
	Separator     = os.PathSeparator
	ListSeparator = os.PathListSeparator
)
```
示例代码:
```go
package main

import "fmt"
import "path/filepath"

func main() {
    fmt.Printf("Separator: %c\t %d\n", filepath.Separator, filepath.Separator) // Separator: \     92  目录分隔符
    fmt.Printf("ListSeparator: %c\t %d\n", filepath.ListSeparator, filepath.ListSeparator) // ListSeparator: ;         59  路径列表分隔符
}
```

## Variables 变量
```
var ErrBadPattern = errors.New("syntax error in pattern") // 路径格式错误
var SkipAll error = fs.SkipAll // WalkFunc 跳过所有文件
var SkipDir error = fs.SkipDir // WalkFunc 跳过目录  
```
示例代码:
```go
package main

import "fmt"
import "path/filepath"

func main() {
    fmt.Printf("ErrBadPattern: %v\n", filepath.ErrBadPattern) // ErrBadPattern: syntax error in pattern
    fmt.Printf("SkipAll: %v\n", filepath.SkipAll) // SkipAll: skip everything and stop the walk
    fmt.Printf("SkipDir: %v\n", filepath.SkipDir) // SkipDir: skip the directory and continue the walk
}
```

## Functions 方法
### func Abs 返回文件绝对路径
```
func Abs(path string) (string, error)
```
示例代码:
```go
package main

import "fmt"
import "path/filepath"

func main() {
    fmt.Println(filepath.Abs("/a/b")) // E:\a\b <nil>
    fmt.Println(filepath.Abs("/a/b/")) // E:\a\b <nil>
    fmt.Println(filepath.Abs("/a/b/c.txt")) // E:\a\b\c.txt <nil>
    fmt.Println(filepath.Abs("/")) // E:\ <nil>
    fmt.Println(filepath.Abs("")) // E:\work\github\golang-cookbook\code\standard <nil>
}
```

### func Base 返回文件基础路径名
```
func Base(path string) string
```
示例代码:
```go
package main

import "fmt"
import "path/filepath"

func main() {
    fmt.Println(filepath.Base("/foo/bar/baz.js")) // baz.js
    fmt.Println(filepath.Base("/foo/bar/baz")) // baz
    fmt.Println(filepath.Base("/foo/bar/baz/")) // baz
    fmt.Println(filepath.Base("/foo/bar/baz/dev.txt")) // dev.txt
    fmt.Println(filepath.Base("../todo.txt")) // todo.txt
    fmt.Println(filepath.Base("..")) // ..
    fmt.Println(filepath.Base(".")) // .
    fmt.Println(filepath.Base("/")) // \
    fmt.Println(filepath.Base("")) // .
}
```

### func Clean 返回清理路径
```
func Clean(path string) string

1 Replace multiple Separator elements with a single one.
2 Eliminate each . path name element (the current directory).
3 Eliminate each inner .. path name element (the parent directory) along with the non-.. element that precedes it.
4 Eliminate .. elements that begin a rooted path: that is, replace "/.." by "/" at the beginning of a path, assuming Separator is '/'.
```
示例代码:
```go
package main

import "fmt"
import "path/filepath"

func main() {
    fmt.Println(filepath.Clean("/foo/bar/baz.js")) // \foo\bar\baz.js
    fmt.Println(filepath.Clean("/foo/bar/baz")) // \foo\bar\baz
    fmt.Println(filepath.Clean("/foo/bar/baz/")) // \foo\bar\baz
    fmt.Println(filepath.Clean("/foo/bar/baz/dev.txt")) // \foo\bar\baz\dev.txt
    fmt.Println(filepath.Clean("../todo.txt")) // ..\todo.txt
    fmt.Println(filepath.Clean("..")) // ..
    fmt.Println(filepath.Clean(".")) // .
    fmt.Println(filepath.Clean("////")) // \\\\\
    fmt.Println(filepath.Clean("")) // .
}
```

### func Dir 返回目录路径
```
func Dir(path string) string
```
示例代码:
```go
package main

import "fmt"
import "path/filepath"

func main() {
    fmt.Println(filepath.Dir("/foo/bar/baz.js")) // \foo\bar
    fmt.Println(filepath.Dir("/foo/bar/baz")) // \foo\bar
    fmt.Println(filepath.Dir("/foo/bar/baz/")) // \foo\bar\baz
    fmt.Println(filepath.Dir("/foo/bar/baz/dev.txt")) // \foo\bar\baz
    fmt.Println(filepath.Dir("../todo.txt")) // ..
    fmt.Println(filepath.Dir("..")) // .
    fmt.Println(filepath.Dir(".")) // .
    fmt.Println(filepath.Dir("////")) // \\\\\
    fmt.Println(filepath.Dir("")) // .
}
```

### func EvalSymlinks 返回符号链接的实际文件的路径
```
func EvalSymlinks(path string) (string, error)
```
示例代码:
```go
package main

import "fmt"
import "path/filepath"

func main() {
    fmt.Println(filepath.EvalSymlinks("/foo/bar/baz.js")) // GetFileAttributesEx /foo: The system cannot find the file specified.
    fmt.Println(filepath.EvalSymlinks("/foo/bar/baz")) // GetFileAttributesEx /foo: The system cannot find the file specified.
    fmt.Println(filepath.EvalSymlinks("/foo/bar/baz/")) // GetFileAttributesEx /foo: The system cannot find the file specified.
    fmt.Println(filepath.EvalSymlinks("/foo/bar/baz/dev.txt")) // GetFileAttributesEx /foo: The system cannot find the file specified.   
    fmt.Println(filepath.EvalSymlinks("../todo.txt")) // .. <nil>
    fmt.Println(filepath.EvalSymlinks("..")) // .. <nil>
    fmt.Println(filepath.EvalSymlinks(".")) // . <nil>
    fmt.Println(filepath.EvalSymlinks("////")) // \\\\ <nil>
    fmt.Println(filepath.EvalSymlinks("")) // . <nil>
}
```

### func Ext 返回文件扩展名
```
func Ext(path string) string
```
示例代码:
```go
package main

import "fmt"
import "path/filepath"

func main() {
    fmt.Printf("No dots: %q\n", filepath.Ext("index")) // No dots: "" 
    fmt.Printf("One dot: %q\n", filepath.Ext("index.js")) // One dot: ".js"
    fmt.Printf("Two dots: %q\n", filepath.Ext("main.test.js")) // Two dots: ".js"
    fmt.Println(filepath.Ext("////")) // ""
    fmt.Println(filepath.Ext("")) // ""
    fmt.Println(filepath.Ext("..")) // .
    fmt.Println(filepath.Ext(".")) // .
    fmt.Println(filepath.Ext("../..")) // .
}
```

### func FromSlash 返回路径转换为Separator分隔的路径
```
func FromSlash(path string) string

将 / 转换为 filepath.Separator
```
示例代码:
```go
package main

import "fmt"
import "path/filepath"

func main() {
    fmt.Println(filepath.FromSlash("/foo/bar/baz.js")) // \foo\bar\baz.js
    fmt.Println(filepath.FromSlash("/foo/bar/baz")) // \foo\bar\baz
    fmt.Println(filepath.FromSlash("/foo/bar/baz/")) // \foo\bar\baz
    fmt.Println(filepath.FromSlash("/foo/bar/baz/dev.txt")) // \foo\bar\baz\dev.txt   
    fmt.Println(filepath.FromSlash("../todo.txt")) // ..\todo.txt
    fmt.Println(filepath.FromSlash("..")) // .. 
    fmt.Println(filepath.FromSlash(".")) // .
    fmt.Println(filepath.FromSlash("////")) // \\\\
    fmt.Println(filepath.FromSlash("")) // ""
    fmt.Println(filepath.FromSlash("..\\..")) // ..\..
}
```

### func Glob 返回匹配模式的文件路径下的所有文件&目录
```
func Glob(pattern string) (matches []string, err error)
```
示例代码:
```go
package main

import "fmt"
import "path/filepath"

func main() {
    fmt.Println(filepath.Glob("./*")) // [buildin-constants.go buildin-type.go buildin.go filepath.go path.go version.go] <nil>
    fmt.Println(filepath.Glob("../../docs/*.md")) // [..\..\docs\README.md ..\..\docs\_coverpage.md ..\..\docs\_glossary.md ..\..\docs\_sidebar.md] <nil>
    fmt.Println(filepath.Glob("/")) // [/] <nil>
    fmt.Println(filepath.Glob("e:/*")) // [e:\目录下的所有文件] <nil>
}
```

### func IsAbs 返回路径是否为绝对路径
```
func IsAbs(path string) bool
```
示例代码:
```go
package main

import "fmt"
import "path/filepath"

func main() {
    fmt.Println(filepath.IsAbs("d:/baz.js")) // true  
    fmt.Println(filepath.IsAbs("/foo/bar/baz")) // false  win & *nix 系统下有区别
    fmt.Println(filepath.IsAbs("/foo/bar/baz/")) // false
    fmt.Println(filepath.IsAbs("/foo/bar/baz/dev.txt")) // false   
    fmt.Println(filepath.IsAbs("../todo.txt")) // false
    fmt.Println(filepath.IsAbs("..")) // false
    fmt.Println(filepath.IsAbs(".")) // false
    fmt.Println(filepath.IsAbs("////")) // true
    fmt.Println(filepath.IsAbs("")) // false
    fmt.Println(filepath.IsAbs("..\\..")) // false  
}
```

### func IsLocal 返回路径是否为本地路径
```
func IsLocal(path string) bool

is within the subtree rooted at the directory in which path is evaluated
is not an absolute path
is not empty
on Windows, is not a reserved name such as "NUL"
```
示例代码:
```go
package main

import "fmt"
import "path/filepath"

func main() {
    fmt.Println(filepath.IsLocal("d:/baz.js")) // true  
    fmt.Println(filepath.IsLocal("/foo/bar/baz")) // false  win & *nix 系统下有区别
    fmt.Println(filepath.IsLocal("/foo/bar/baz/")) // false
    fmt.Println(filepath.IsLocal("/foo/bar/baz/dev.txt")) // false   
    fmt.Println(filepath.IsLocal("../todo.txt")) // false
    fmt.Println(filepath.IsLocal("..")) // false
    fmt.Println(filepath.IsLocal(".")) // false
    fmt.Println(filepath.IsLocal("////")) // true
    fmt.Println(filepath.IsLocal("")) // false
    fmt.Println(filepath.IsLocal("..\\..")) // false  
}
```

### func Join 连接任意数量的路径元素
```
func Join(elem ...string) string
```
示例代码:
```go
package main

import "fmt"
import "path/filepath"

func main() {
    fmt.Println(filepath.Join("a", "b", "c")) // a\b\c
    fmt.Println(filepath.Join("a", "b/c")) // a\b\c
    fmt.Println(filepath.Join("a/b", "c")) // a\b\c
    fmt.Println(filepath.Join("a/b", "/c")) // a\b\c
    fmt.Println(filepath.Join("a/b", "../../../xyz")) // ..\xyz
}
```

### func Localize 返回本地路径转换为Separator分隔的路径
```
func Localize(path string) (string, error)

Localize converts a slash-separated path into an operating system path. The input path must be a valid path as reported by io/fs.ValidPath.
```
示例代码:
```go
package main

import "fmt"
import "path/filepath"

func main() {
    fmt.Println(filepath.Localize("d:/baz.js")) // d:\baz.js
    fmt.Println(filepath.Localize("/foo/bar/baz")) // false  win & *nix 系统下有区别
    fmt.Println(filepath.Localize("/foo/bar/baz/")) // false
    fmt.Println(filepath.Localize("/foo/bar/baz/dev.txt")) // false   
    fmt.Println(filepath.Localize("../todo.txt")) // false
    fmt.Println(filepath.Localize("..")) // false
    fmt.Println(filepath.Localize(".")) // false
    fmt.Println(filepath.Localize("////")) // true
    fmt.Println(filepath.Localize("")) // false
    fmt.Println(filepath.Localize("..\\..")) // false   
}
```

### func Match 路径是否匹配模式
```
func Match(pattern, name string) (matched bool, err error)

pattern:
    { term }
term:
    '*'         matches any sequence of non-Separator characters
    '?'         matches any single non-Separator character
    '[' [ '^' ] { character-range } ']'
                character class (must be non-empty)
    c           matches character c (c != '*', '?', '\\', '[')
    '\\' c      matches character c (except on Windows)

character-range:
    c           matches character c (c != '\\', '-', ']')
    '\\' c      matches character c (except on Windows)
    lo '-' hi   matches character c for lo <= c <= hi
```
示例代码:
```go
package main

import "fmt"
import "path/filepath"

func main() {
    fmt.Println(filepath.Match("/home/catch/*", "/home/catch/foo")) // true <nil>s
    fmt.Println(filepath.Match("/home/catch/*", "/home/catch/foo/bar")) // true <nil>
    fmt.Println(filepath.Match("/home/?opher", "/home/gopher")) // true <nil>
    fmt.Println(filepath.Match("/home/\\*", "/home/*")) // false <nil>
}
```


### func Rel 返回相对路径名
```
func Rel(basePath, targPath string) (string, error)
```
示例代码:
```go
package main

import "fmt"
import "path/filepath"

func main() {
    paths := []string{
        "/a/b/c", // "/a/b/c": "b\\c" <nil>
        "/b/c", // "/b/c": "..\\b\\c" <nil>
        "./b/c", // "./b/c": "" Rel: can't make ./b/c relative to /a
    }
    for _, p := range paths {
        rel, err := filepath.Rel("/a", p)
        fmt.Printf("%q: %q %v\n", p, rel, err)
    }
}
```

### func Split 返回目录和文件路径
```
func Split(path string) (dir, file string)
```
示例代码:
```go
package main

import "fmt"
import "path/filepath"

func main() {
    fmt.Println(filepath.Split("/home/arnie/amelia.jpg")) // /home/arnie, amelia.jpg
    fmt.Println(filepath.Split("/mnt/photos/")) // /mnt/photos/, ""
    fmt.Println(filepath.Split("rabbit.jpg")) //  "", rabbit.jpg
    fmt.Println(filepath.Split("/usr/local//go")) // /usr/local// go
    fmt.Println(filepath.Split("")) // ”“, ""
    fmt.Println(filepath.Split(".")) // ”“, .
    fmt.Println(filepath.Split("..")) //  ”“, ..
    fmt.Println(filepath.Split("/")) // /, ""
}
```

### func SplitList 获取路径列表
```
func SplitList(path string) []string
```
示例代码:
```go
package main

import "fmt"
import "path/filepath"

func main() {
    fmt.Println("On Unix:", filepath.SplitList("/a/b/c:/usr/bin")) // On Unix: [/a/b/c:/usr/bin]
    fmt.Println("On Windows:", filepath.SplitList("/a/b/c;/usr/bin")) // On Windows: [/a/b/c /usr/bin]
}
```

### func ToSlash 将路径 Separator 分隔的路径转换为 / 分隔的路径
```
func ToSlash(path string) string

replacing each separator character in path with a slash ('/') character. Multiple separators are replaced by multiple slashes.
```
示例代码:
```go
package main

import "fmt"
import "path/filepath"

func main() {
    fmt.Println(filepath.ToSlash("d:\\aaa\\bbb\\ccc.txt")) // d:/aaa/bbb/ccc.txt
    fmt.Println(filepath.ToSlash("/home/arnie/amelia.jpg")) // /home/arnie/amelia.jpg
    fmt.Println(filepath.ToSlash("/mnt/photos/")) // /mnt/photos/
    fmt.Println(filepath.ToSlash("rabbit.jpg")) //  rabbit.jpg
    fmt.Println(filepath.ToSlash("/usr/local//go")) // /usr/local//go
    fmt.Println(filepath.ToSlash("")) // ""
    fmt.Println(filepath.ToSlash(".")) //.
    fmt.Println(filepath.ToSlash("..")) // ..
    fmt.Println(filepath.ToSlash("/")) // /
}
```


### func VolumeName 获取盘符名称
```
func VolumeName(path string) string

VolumeName returns leading volume name. Given "C:\foo\bar" it returns "C:" on Windows. Given "\\host\share\foo" it returns "\\host\share". On other platforms it returns "".
```
示例代码:
```go
package main

import "fmt"
import "path/filepath"

func main() {
    fmt.Println(filepath.VolumeName("C:\\foo\\bar")) // C:
    fmt.Println(filepath.VolumeName("\\host\\share\\foo")) // ""
    fmt.Println(filepath.VolumeName("/mnt/photos/")) // ""
    fmt.Println(filepath.VolumeName("rabbit.jpg")) //  ""
    fmt.Println(filepath.VolumeName("/usr/local//go")) // ""
    fmt.Println(filepath.VolumeName("")) // ""
    fmt.Println(filepath.VolumeName(".")) // ""
    fmt.Println(filepath.VolumeName("..")) // ""
    fmt.Println(filepath.VolumeName("/")) // ""
}
```

### func Walk 遍历目录树
```
func Walk(root string, fn WalkFunc) error
```
示例代码:
```go
package main

import "fmt"
import "path/filepath"
import "io/fs"

func main() {
    path1 := "../"
    err := filepath.Walk(path1, func(path string, info fs.FileInfo, err error) error {
        if err != nil {
            fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
            return err
        }
        if info.IsDir() && info.Name() == ".git" {
            fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
            return filepath.SkipDir
        }
        fmt.Printf("visited file or dir: %q\n", path)
        return nil
    })
    if err != nil {
        fmt.Printf("error walking the path %q: %v\n", path1, err)
        return
    }
}
```

### func WalkDir 遍历目录树，不递归子目录
```
func WalkDir(root string, fn WalkFunc) error
```
示例代码:
```go
package main

import "fmt"
import "path/filepath"
import "io/fs"

func main() {
    path1 := "../"
    err := filepath.Walk(path1, func(path string, info fs.FileInfo, err error) error {
        if err != nil {
            fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
            return err
        }
        if info.IsDir() && info.Name() == ".git" {
            fmt.Printf("skipping a dir without errors: %+v \n", info.Name())
            return filepath.SkipDir
        }
        fmt.Printf("visited file or dir: %q\n", path)
        return nil
    })
    if err != nil {
        fmt.Printf("error walking the path %q: %v\n", path1, err)
        return
    }
}
```



## Types 类型
### type WalkFunc
```
type WalkFunc func(path string, info fs.FileInfo, err error) error
```
示例代码:
```go
<示例代码>
```
查看 func Walk 的示例代码
```