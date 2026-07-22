package main

import "fmt"
import "path/filepath"
import "io/fs"

func main() {
    fmt.Println("// == Constants 常量 ====================================================================================================")
    fmt.Printf("Separator: %c\t %d\n", filepath.Separator, filepath.Separator) // Separator: \     92  目录分隔符
    fmt.Printf("ListSeparator: %c\t %d\n", filepath.ListSeparator, filepath.ListSeparator) // ListSeparator: ;         59  路径列表分隔符

    fmt.Println("// == Variables 变量 ====================================================================================================")
    fmt.Printf("ErrBadPattern: %v\n", filepath.ErrBadPattern) // ErrBadPattern: syntax error in pattern
    fmt.Printf("SkipAll: %v\n", filepath.SkipAll) // SkipAll: skip everything and stop the walk
    fmt.Printf("SkipDir: %v\n", filepath.SkipDir) // SkipDir: skip the directory and continue the walk

    fmt.Println("// == func Abs 返回文件绝对路径 ====================================================================================================")
    fmt.Println(filepath.Abs("/a/b")) // E:\a\b <nil>
    fmt.Println(filepath.Abs("/a/b/")) // E:\a\b <nil>
    fmt.Println(filepath.Abs("/a/b/c.txt")) // E:\a\b\c.txt <nil>
    fmt.Println(filepath.Abs("/")) // E:\ <nil>
    fmt.Println(filepath.Abs("")) // E:\work\github\golang-cookbook\code\standard <nil>

    fmt.Println("// == func Base 返回文件基础名 ====================================================================================================")
    fmt.Println(filepath.Base("/foo/bar/baz.js")) // baz.js
    fmt.Println(filepath.Base("/foo/bar/baz")) // baz
    fmt.Println(filepath.Base("/foo/bar/baz/")) // baz
    fmt.Println(filepath.Base("/foo/bar/baz/dev.txt")) // dev.txt
    fmt.Println(filepath.Base("../todo.txt")) // todo.txt
    fmt.Println(filepath.Base("..")) // ..
    fmt.Println(filepath.Base(".")) // .
    fmt.Println(filepath.Base("/")) // \
    fmt.Println(filepath.Base("")) // .

    fmt.Println("// == func Clean 返回清理路径 ====================================================================================================")
    fmt.Println(filepath.Clean("/foo/bar/baz.js")) // \foo\bar\baz.js
    fmt.Println(filepath.Clean("/foo/bar/baz")) // \foo\bar\baz
    fmt.Println(filepath.Clean("/foo/bar/baz/")) // \foo\bar\baz
    fmt.Println(filepath.Clean("/foo/bar/baz/dev.txt")) // \foo\bar\baz\dev.txt
    fmt.Println(filepath.Clean("../todo.txt")) // ..\todo.txt
    fmt.Println(filepath.Clean("..")) // ..
    fmt.Println(filepath.Clean(".")) // .
    fmt.Println(filepath.Clean("////")) // \\\\\
    fmt.Println(filepath.Clean("")) // .

    fmt.Println("// == func Dir 返回目录路径 ====================================================================================================")
    fmt.Println(filepath.Dir("/foo/bar/baz.js")) // \foo\bar
    fmt.Println(filepath.Dir("/foo/bar/baz")) // \foo\bar
    fmt.Println(filepath.Dir("/foo/bar/baz/")) // \foo\bar\baz
    fmt.Println(filepath.Dir("/foo/bar/baz/dev.txt")) // \foo\bar\baz
    fmt.Println(filepath.Dir("../todo.txt")) // ..
    fmt.Println(filepath.Dir("..")) // .
    fmt.Println(filepath.Dir(".")) // .
    fmt.Println(filepath.Dir("////")) // \\\\\
    fmt.Println(filepath.Dir("")) // .

    fmt.Println("// == func EvalSymlinks 返回符号链接后的路径 ====================================================================================================")
    fmt.Println(filepath.EvalSymlinks("/foo/bar/baz.js")) // GetFileAttributesEx /foo: The system cannot find the file specified.
    fmt.Println(filepath.EvalSymlinks("/foo/bar/baz")) // GetFileAttributesEx /foo: The system cannot find the file specified.
    fmt.Println(filepath.EvalSymlinks("/foo/bar/baz/")) // GetFileAttributesEx /foo: The system cannot find the file specified.
    fmt.Println(filepath.EvalSymlinks("/foo/bar/baz/dev.txt")) // GetFileAttributesEx /foo: The system cannot find the file specified.   
    fmt.Println(filepath.EvalSymlinks("../todo.txt")) // .. <nil>
    fmt.Println(filepath.EvalSymlinks("..")) // .. <nil>
    fmt.Println(filepath.EvalSymlinks(".")) // . <nil>
    fmt.Println(filepath.EvalSymlinks("////")) // \\\\ <nil>
    fmt.Println(filepath.EvalSymlinks("")) // . <nil>

    fmt.Println("// == func Ext 返回文件扩展名 ====================================================================================================")
    fmt.Printf("No dots: %q\n", filepath.Ext("index")) // No dots: "" 
    fmt.Printf("One dot: %q\n", filepath.Ext("index.js")) // One dot: ".js"
    fmt.Printf("Two dots: %q\n", filepath.Ext("main.test.js")) // Two dots: ".js"
    fmt.Println(filepath.Ext("////")) // ""
    fmt.Println(filepath.Ext("")) // ""
    fmt.Println(filepath.Ext("..")) // .
    fmt.Println(filepath.Ext(".")) // .
    fmt.Println(filepath.Ext("../..")) // .

    fmt.Println("// == func Glob 返回匹配模式的文件路径下的所有文件&目录 ====================================================================================================")
    fmt.Println(filepath.Glob("./*")) // [buildin-constants.go buildin-type.go buildin.go filepath.go path.go version.go] <nil>
    fmt.Println(filepath.Glob("../../docs/*.md")) // [..\..\docs\README.md ..\..\docs\_coverpage.md ..\..\docs\_glossary.md ..\..\docs\_sidebar.md] <nil>
    fmt.Println(filepath.Glob("/")) // [/] <nil>
    fmt.Println(filepath.Glob("e:/*")) // [e:\目录下的所有文件] <nil>

    fmt.Println("// == func IsAbs 返回路径是否为绝对路径 ====================================================================================================")
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

    fmt.Println("// == func IsLocal 返回路径是否为本地路径 ====================================================================================================") 
    fmt.Println(filepath.IsLocal("d:/baz.js")) // false  
    fmt.Println(filepath.IsLocal("/foo/bar/baz")) // false  win & *nix 系统下有区别
    fmt.Println(filepath.IsLocal("/foo/bar/baz/")) // false
    fmt.Println(filepath.IsLocal("/foo/bar/baz/dev.txt")) // false   
    fmt.Println(filepath.IsLocal("../todo.txt")) // false
    fmt.Println(filepath.IsLocal("..")) // false
    fmt.Println(filepath.IsLocal(".")) // true
    fmt.Println(filepath.IsLocal("////")) // false
    fmt.Println(filepath.IsLocal("")) // false
    fmt.Println(filepath.IsLocal("..\\..")) // false  

    fmt.Println("// == func Join 连接任意数量的路径元素 ====================================================================================================") 
    fmt.Println(filepath.Join("a", "b", "c")) // a\b\c
    fmt.Println(filepath.Join("a", "b/c")) // a\b\c
    fmt.Println(filepath.Join("a/b", "c")) // a\b\c
    fmt.Println(filepath.Join("a/b", "/c")) // a\b\c
    fmt.Println(filepath.Join("a/b", "../../../xyz")) // ..\xyz

    fmt.Println("// == func Localize 返回本地路径转换为Separator分隔的路径 ====================================================================================================") 
    fmt.Println(filepath.Localize("d:/baz.js")) //  invalid path
    fmt.Println(filepath.Localize("/foo/bar/baz")) //  invalid path
    fmt.Println(filepath.Localize("/foo/bar/baz/")) //  invalid path
    fmt.Println(filepath.Localize("/foo/bar/baz/dev.txt")) //  invalid path   
    fmt.Println(filepath.Localize("../todo.txt")) //  invalid path
    fmt.Println(filepath.Localize("..")) //  invalid path
    fmt.Println(filepath.Localize(".")) // . <nil>
    fmt.Println(filepath.Localize("////")) //  invalid path
    fmt.Println(filepath.Localize("")) //  invalid path
    fmt.Println(filepath.Localize("..\\..")) //  invalid path   

    fmt.Println("// == func Match 路径是否匹配模式 ====================================================================================================") 
    fmt.Println(filepath.Match("/home/catch/*", "/home/catch/foo")) // true <nil>s
    fmt.Println(filepath.Match("/home/catch/*", "/home/catch/foo/bar")) // true <nil>
    fmt.Println(filepath.Match("/home/?opher", "/home/gopher")) // true <nil>
    fmt.Println(filepath.Match("/home/\\*", "/home/*")) // false <nil>

    fmt.Println("// == func Rel 返回相对路径名 ====================================================================================================")
    paths := []string{
        "/a/b/c", // "/a/b/c": "b\\c" <nil>
        "/b/c", // "/b/c": "..\\b\\c" <nil>
        "./b/c", // "./b/c": "" Rel: can't make ./b/c relative to /a
    }
    for _, p := range paths {
        rel, err := filepath.Rel("/a", p)
        fmt.Printf("%q: %q %v\n", p, rel, err)
    }

    fmt.Println("// == func Split 返回目录和文件路径 ====================================================================================================")
    fmt.Println(filepath.Split("/home/arnie/amelia.jpg")) // /home/arnie, amelia.jpg
    fmt.Println(filepath.Split("/mnt/photos/")) // /mnt/photos/, ""
    fmt.Println(filepath.Split("rabbit.jpg")) //  "", rabbit.jpg
    fmt.Println(filepath.Split("/usr/local//go")) // /usr/local// go
    fmt.Println(filepath.Split("")) // ”“, ""
    fmt.Println(filepath.Split("..")) //  ”“, ..
    fmt.Println(filepath.Split("/")) // /, ""

    fmt.Println("// == func SplitList 获取路径列表 ====================================================================================================")
    fmt.Println("On Unix:", filepath.SplitList("/a/b/c:/usr/bin")) // On Unix: [/a/b/c:/usr/bin]
    fmt.Println("On Windows:", filepath.SplitList("/a/b/c;/usr/bin")) // On Windows: [/a/b/c /usr/bin]

    
    fmt.Println("// == func ToSlash 将路径 Separator 分隔的路径转换为 / 分隔的路径 ====================================================================================================")
    fmt.Println(filepath.ToSlash("d:\\aaa\\bbb\\ccc.txt")) // d:/aaa/bbb/ccc.txt
    fmt.Println(filepath.ToSlash("/home/arnie/amelia.jpg")) // /home/arnie/amelia.jpg
    fmt.Println(filepath.ToSlash("/mnt/photos/")) // /mnt/photos/
    fmt.Println(filepath.ToSlash("rabbit.jpg")) //  rabbit.jpg
    fmt.Println(filepath.ToSlash("/usr/local//go")) // /usr/local//go
    fmt.Println(filepath.ToSlash("")) // ""
    fmt.Println(filepath.ToSlash(".")) //.
    fmt.Println(filepath.ToSlash("..")) // ..
    fmt.Println(filepath.ToSlash("/")) // /

    fmt.Println("// == func VolumeName 获取盘符名称 ====================================================================================================")
    fmt.Println(filepath.VolumeName("C:\\foo\\bar")) // C:
    fmt.Println(filepath.VolumeName("\\host\\share\\foo")) // ""
    fmt.Println(filepath.VolumeName("/mnt/photos/")) // ""
    fmt.Println(filepath.VolumeName("rabbit.jpg")) //  ""
    fmt.Println(filepath.VolumeName("/usr/local//go")) // ""
    fmt.Println(filepath.VolumeName("")) // ""
    fmt.Println(filepath.VolumeName(".")) // ""
    fmt.Println(filepath.VolumeName("..")) // ""
    fmt.Println(filepath.VolumeName("/")) // ""

    fmt.Println("// == func Walk 遍历目录树 ====================================================================================================")
    path1 := "../../"
    err1 := filepath.Walk(path1, func(path string, info fs.FileInfo, err error) error {
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
    if err1 != nil {
        fmt.Printf("error walking the path %q: %v\n", path1, err1)
        return
    }

    fmt.Println("// == func WalkDir 遍历目录树 ====================================================================================================")
    path2 := "../../"
    err2 := filepath.WalkDir(path2, func(path string, info fs.DirEntry, err error) error {
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
    if err2 != nil {
        fmt.Printf("error walking the path %q: %v\n", path2, err2)
        return
    }
}
