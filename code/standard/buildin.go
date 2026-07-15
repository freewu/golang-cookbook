package main

import "fmt"

func main() {
    // == append 追加切片元素 ========================================================================================================
    arr := []int{1, 2, 3}
    fmt.Println(arr) // [1 2 3]
    fmt.Println(append(arr, 4, 5)) // [1 2 3 4 5]
    arr1 := []string{"a", "b", "c"}
    fmt.Println(arr1) // [a b c]
    fmt.Println(append(arr1, "bluefrog", "freewu")) // [a b c bluefrog freewu]

    // == cap 获取切片 & 数组容量 ====================================================================================================
    // 数组
    var arr11 [3]int // 方式1：固定长度声明 [0,0,0] 
    arr12 := [3]int{1,2,3}  // 方式2：初始化赋值 
    arr13 := [...]int{10,20,30} // 方式3：自动推导长度
    fmt.Println(cap(arr11)) // 3
    fmt.Println(cap(arr12)) // 3
    fmt.Println(cap(arr13)) // 3
    // 切片
    var slice11 []int // 零值切片：nil len=0 cap=0
    slice12 := []int{1,2,3} // 直接初始化 [1 2 3] len=3 cap=3
    slice13 := make([]int, 2) // make 创建 len=2, cap=2，元素默认0
    slice14 := make([]int, 2, 5)   // make 创建 len=2，cap=5，预留扩容空间
    fmt.Println(slice11, cap(slice11)) // nil 0
    fmt.Println(slice12, cap(slice12)) // [1 2 3] 3
    fmt.Println(slice13, cap(slice13)) // [0 0] 2
    fmt.Println(slice14, cap(slice14)) // [0 0] 5

    // == close 清空切片或 map ======================================================================================================
    slice21 := []int{1, 2, 3}
    fmt.Println(slice21) // [1 2 3]
    clear(slice21)
    fmt.Println(slice21) // []
    mp21 := make(map[string]int)
    mp21["a"] = 1
    mp21["b"] = 2
    fmt.Println(mp21) // map[a:1 b:2]
    clear(mp21)
    fmt.Println(mp21) // map[]

    // == close 关闭 channel ======================================================================================================
    ch := make(chan int, 3)
    go producer(ch)
    // range 读取，通道关闭自动跳出循环
    for val := range ch {
        fmt.Println("收到：", val)
    }
    fmt.Println("读取完成，通道已关闭")

    // == complex 生成一个 ComplexType 变量 ========================================================================================
    complex1 := complex(1, 2) // 1+2i
    fmt.Println(complex1) // (1+2i)
    complex2 := complex(1024, -2) // 1024-2i
    fmt.Println(complex2) // (1024-2i)

    // == copy 复制切片 ===========================================================================================================
    slice31 := []int{1, 2, 3}
    fmt.Println(slice31) // [1 2 3]
    slice32 := make([]int, 3)
    fmt.Println(copy(slice32, slice31)) // 3 复制切片 slice31 到 slice32
    fmt.Println(slice32) // [1 2 3]
    slice33 := make([]int, 5)
    fmt.Println(copy(slice33, slice31)) // 3 复制切片 slice31 到 slice33
    fmt.Println(slice33) // [1 2 3 0 0]

    // == delete 删除 map 元素 ====================================================================================================
    mp41 := map[string]int{"a": 1, "b": 2}
    fmt.Println(mp41) // map[a:1 b:2]
    delete(mp41, "a")
    fmt.Println(mp41) // map[b:2]
    delete(mp41, "c") // 删除一个不存在的
    fmt.Println(mp41) // map[b:2]

    // == imag 获取虚数部分的值 ====================================================================================================
    complex51 := complex(1, 2) // 1+2i
    fmt.Println(imag(complex51)) // 2
    complex52 := complex(1024, -2) // 1024-2i
    fmt.Println(imag(complex52)) // -2

    // == len 获取切片、数组、字符串、通道的长度 ====================================================================================================
    arr61 := [5]int{1,2,3,4,5}
    fmt.Println(len(arr61)) // 数组 5
    slice61 := []int{1,2,3,4,5,6,7,8}
    fmt.Println(len(slice61)) // 切片 8
    fmt.Println(len("bluefrog")) // 字符串 8
    ch61 := make(chan int, 5)
    fmt.Println(len(ch61)) // channel 0
    ch61 <- 12
    fmt.Println(len(ch61)) // channel 1
    mp61 := map[string]int{"a": 1, "b": 2}
    fmt.Println(len(mp61)) // map 2

    // == func make 创建 切片、Map、通道 ====================================================================================================
    slice71 := make([]int, 1, 5)
    fmt.Println(slice71) // [0]
    mp71 := make(map[string]int, 5)
    fmt.Println(mp71) // map[]
    ch71 := make(chan int, 5)
    fmt.Println(ch71) // 0xc000024140

    // func max 取最大入参 ====================================================================================================
    fmt.Println(max(1, 2, 3)) // 3
    fmt.Println(max(1, 2, 3, 4, 5)) // 5
    fmt.Println(max(1, 2, 3, 4, 5, 6)) // 6

    // func min 取最小入参 ====================================================================================================
    fmt.Println(min(1, 2, 3)) // 1
    fmt.Println(min(1, 2, 3, 4, 5)) // 1
    fmt.Println(min(1, 2, 3, 4, 5, 6)) // 1
    
    // func new 创建结构体实例 ====================================================================================================
    s := new(Student)
    fmt.Println(s) // &{}
    fmt.Println(s.Name) // ""
    fmt.Println(s.Age) // 0
    s.Name = "张三"
    s.Age = 18
    fmt.Println(s) // &{Name:张三 Age:18}

    // func panic 引发 panic 终止程序 ====================================================================================================
    // panic("wrong wrong wrong!!!") // panic: wrong wrong wrong!!! 程序终止
    fmt.Println("bluefrog") // bluefrog 永远不会输出

    // == func print 打印输出 ====================================================================================================
    print("bluefrog ") // bluefrog 
    print("go go go ")
    // 会输出在一行  bluefrog go go go
    // 等价于 fmt.Print
    fmt.Println()
    fmt.Print("bluefrog ")
    fmt.Print("go go go ")

    // == func println 打印输出(结束换行) ====================================================================================================
    println("bluefrog ") // bluefrog 
    println("go go go ") // go go go
    // 等价于 fmt.Println
    fmt.Println()
    fmt.Println("bluefrog ")
    fmt.Println("go go go ")
    // 等价于 fmt.Println
    fmt.Println()
    print("bluefrog\r\n") // bluefrog 
    print("go go go \r\n") // go go go

    // == func real 获取复数的实部 ====================================================================================================
    complex81 := complex(1, 2) // 1+2i
    fmt.Println(real(complex81)) // 1
    complex82 := complex(1024, -2) // 1024-2i
    fmt.Println(real(complex82)) // 1024

    // == func recover 恢复 panic ====================================================================================================
    go handleRequest() // 请求崩溃，err=assignment to entry in nil map
    fmt.Println("1111") 
    //select {} // 阻塞不退出
}

func producer(ch chan<- int) { // chan<- int 只写通道，匹配close参数
    for i := 1; i <= 3; i++ {
        ch <- i
    }
    // 发送完毕，关闭通道
    close(ch)
    fmt.Println("生产者：通道已关闭")
}

type Student struct {
    Name string
    Age int
}

func handleRequest() {
    // 请求兜底捕获
    // Web / 后台服务中，每个请求 goroutine 必须加 recover，防止单个请求崩溃导致整个服务下线
    defer func() { // recover 只能在 defer 匿名函数内部 调用才生效
        // 多个 defer 按注册逆序执行，第一个匹配的 recover 捕获 panic，上层不再处理
        if err := recover(); err != nil {
            // 日志记录堆栈
            fmt.Printf("请求崩溃，err=%v\n", err) // 捕获后程序恢复正常流程，不会退出
        }
    }()
    // 业务逻辑
    var m map[string]int
    m["key"] = 1 // 空map赋值，触发panic
}