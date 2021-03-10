package main

import "fmt"

func main() {
    // bool
    var a bool = true
    var b bool = false
    fmt.Println("a:", a)
    fmt.Println("b:", b)
    c := a && b
    fmt.Println("c:", c)
    d := a || b
    fmt.Println("d:", d)

    // Signed integers
    // int8: size : 8 bits, range: -128 to 127
    var a_int8 int8 = 10
    fmt.Println("int8: ", a_int8)
    // int16: size: 16 bits, range: -32768 to 32767
    var a_int16 int16 = 32767
    fmt.Println("int16:", a_int16)
    // int32: size: 32 bits, range: -2147483648  to 2147483647
    var a_int32 int32 = 2147483647
    fmt.Println("int32", a_int32)
    // int64: size: 64 bits, range: -9223372036854775808  to 9223372036854775807
    var a_int64 int64 = 9223372036854775807
    fmt.Println("int64", a_int64)

    // Unsigned intergers
    // uint8: size: 8 bits, range: 0 to 255
    var a_uint8 uint8 = 255
    fmt.Println("unit8", a_uint8)
    // uint16: size: 16 bits, range: 0 to 65535
    var a_uint16 uint16 = 65535
    fmt.Println("uint16", a_uint16)
    // unit32: size: 32 bits, ranges: 0 to 4294967295
    var a_uint32 uint32 = 4294967295
    fmt.Println("unit32", a_uint32)
    // unit641: size: 64 bits, ranges: 0 to 18446744073709551615
    var a_uint64 uint64 = 18446744073709551615
    fmt.Println("unit64", a_uint64)

    // unit: represent 32 or 64 bits unsigned integers
    // depending on the underlying platform

    // Floating point types
    x, y := 3.14, 5.6
    fmt.Println("x:", x, "y", y)
    diff := y - x
    fmt.Println("diff:", diff)
    sum := x + y
    fmt.Println("sum:", sum)

    // Complex types
    // complex64: complex numbers which have float32 real and imaginary parts
    // complex128: complex numbers which have float64 real and imaginary parts
    c1 := 6 + 7i
    fmt.Println("c:", c1)
    c2 := complex(5, 7)
    fmt.Println("sum:", c1+c2)

    // Other numeric types
    // byte is an alias of uint8
    // rune is an alias of int32

    // String types
    first_name := "Kai"
    last_name := "Nguyen"
    name := first_name + last_name
    fmt.Println("My name is:", name)

    // Type conversion
    i := 55
    j := 67.8
    // sum := i + j -> int + float64 not allowed
    // (mismatched types int and float64)
    sum1 := float64(i) + j
    fmt.Println("sum:", sum1)
    sum2 := i + int(j)
    fmt.Println("sum:", sum2)

}
