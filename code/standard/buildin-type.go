package main

import (
    "fmt"
    "unsafe"
)

func main() {
    fmt.Println("==== Go 内置类型内存占用大小（64位系统） ====\n")

    // 1. 布尔
    var vBool bool
    fmt.Printf("bool        变量大小: %d 字节\n", unsafe.Sizeof(vBool)) // bool        变量大小: 1 字节

    // 2. 无符号整数
    var vByte byte // alias uint8
    var vUint8 uint8
    var vUint16 uint16
    var vUint32 uint32
    var vUint64 uint64
    var vUint uint
    var vUintptr uintptr
    fmt.Printf("byte(uint8) 变量大小: %d 字节\n", unsafe.Sizeof(vByte)) // byte(uint8) 变量大小: 1 字节
    fmt.Printf("uint8       变量大小: %d 字节\n", unsafe.Sizeof(vUint8)) // uint8       变量大小: 1 字节
    fmt.Printf("uint16      变量大小: %d 字节\n", unsafe.Sizeof(vUint16)) // uint16      变量大小: 2 字节
    fmt.Printf("uint32      变量大小: %d 字节\n", unsafe.Sizeof(vUint32)) // uint32      变量大小: 4 字节
    fmt.Printf("uint64      变量大小: %d 字节\n", unsafe.Sizeof(vUint64)) // uint64      变量大小: 8 字节
    fmt.Printf("uint        变量大小: %d 字节\n", unsafe.Sizeof(vUint)) // uint        变量大小: 8 字节
    fmt.Printf("uintptr     变量大小: %d 字节\n", unsafe.Sizeof(vUintptr)) // uintptr     变量大小: 8 字节

    // 3. 有符号整数
    var vInt8 int8
    var vInt16 int16
    var vInt32 int32
    var vInt64 int64
    var vInt int
    var vRune rune // alias int32
    // int8        变量大小: 1 字节
    fmt.Printf("int8        变量大小: %d 字节\n", unsafe.Sizeof(vInt8)) // int8        变量大小: 1 字节
    fmt.Printf("int16       变量大小: %d 字节\n", unsafe.Sizeof(vInt16)) // int16       变量大小: 2 字节
    fmt.Printf("int32(rune) 变量大小: %d 字节\n", unsafe.Sizeof(vInt32)) // int32(rune) 变量大小: 4 字节
    fmt.Printf("int64       变量大小: %d 字节\n", unsafe.Sizeof(vInt64)) // int64       变量大小: 8 字节
    fmt.Printf("int         变量大小: %d 字节\n", unsafe.Sizeof(vInt)) // int         变量大小: 8 字节
    fmt.Printf("rune        变量大小: %d 字节\n", unsafe.Sizeof(vRune)) // rune        变量大小	: 4 字节

    // 4. 浮点 FloatType
    var vFloat32 float32
    var vFloat64 float64
    fmt.Printf("float32     变量大小: %d 字节\n", unsafe.Sizeof(vFloat32)) // float32     变量大小: 4 字节
    fmt.Printf("float64     变量大小: %d 字节\n", unsafe.Sizeof(vFloat64)) // float64     变量大小: 8 字节

    // 5. 复数 ComplexType
    var vComplex64 complex64
    var vComplex128 complex128
    fmt.Printf("complex64   变量大小: %d 字节\n", unsafe.Sizeof(vComplex64)) // complex64   变量大小: 8 字节
    fmt.Printf("complex128  变量大小: %d 字节\n", unsafe.Sizeof(vComplex128)) // complex128  变量大小: 16 字节

    // 6. string 字符串（底层结构体，存指针+len）
    var vStr string
    fmt.Printf("string      变量大小: %d 字节\n", unsafe.Sizeof(vStr)) // string      变量大小: 16 字节

    // 7. any = interface{} 空接口（2个指针：类型+数据）
    var vAny any
    fmt.Printf("any(interface{}) 变量大小: %d 字节\n", unsafe.Sizeof(vAny)) // any(interface{}) 变量大小: 16 字节

    // 8. error 接口（接口类型和any结构一致）
    var vErr error
    fmt.Printf("error       变量大小: %d 字节\n", unsafe.Sizeof(vErr)) // error       变量大小: 16 字节

    // 9. comparable 注意：不能定义变量，是泛型约束，无法Sizeof
    fmt.Println("\n注：comparable 仅泛型约束，不能声明变量，无内存大小")
    fmt.Println("注：Type/Type1/IntegerType 等是编译器语法标识，非真实类型")
}