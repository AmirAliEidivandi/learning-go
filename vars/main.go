package main

import (
	"fmt"
	"strings"
)

func main() {
	// int age = 10; // c++ way
	// int
	var num1 int = 100 // min value = -2147483648 - max value = 2147483647
	_ = num1
	// int8
	var num2 int8 = 127 // mnv = -128 - mxv = 127
	_ = num2
	// int 16
	var num3 int16 = -32768 // mnv = -32768 - mxv = 32767
	_ = num3
	// int 32
	var num4 int32 = 2147483647 // mnv = -2147483648 - mxv = 2147483647 === int
	_ = num4
	// int 64
	var num5 int64 = 9223372036854775807 // mnv = -9223372036854775808 - mxv = 9223372036854775807
	_ = num5

	// uint 8
	var num6 uint8 = 255 // 0 to 255
	_ = num6
	// uint 16
	var num7 uint16 = 65535 // 0 to 65535
	_ = num7
	// uint 32
	var num8 uint32 = 4294967295 // 0 to 4294967295
	_ = num8
	// uint 64
	var num9 uint64 = 18446744073709551615 // 0 to 18446744073709551615
	_ = num9

	// float types
	// the predefined architecture-independent float tyes are -

	// float32
	var floatNum float32 = 756548731902874.5235424234325831724981753495847249731 // IEEE-754 32-bit floating-point numbers
	_ = floatNum
	// float64
	var floatNum2 float64 = 756548731902874.5235424234325831724981753495847249731 // IEEE-754 64-bit floating-point numbers
	_ = floatNum2
	// complex64
	// complex numbers with float32 real and imaginary parts
	var complexNum complex64 = 2432352.53253456651143
	_ = complexNum
	// complex128
	// complex numbers with float64 real and imaginary parts
	var complexNum2 complex128 = 2432352.53253456651143
	_ = complexNum2

	// other numeric types
	// there is also a set of numberic types with implementation-specific sizes -
	// byte same as uint8
	var byteNum byte = 255
	_ = byteNum
	// rune same as int32
	var runeNum rune = 2147483647
	_ = runeNum

	// uintptr
	// an unsigned integer to store the uninterpreted bits of pointer value
	var uintptrNum uintptr = 3425325424327084
	_ = uintptrNum

	// type
	type UserId int
	type Direction byte
	type Speed float64
	type Velocity Speed

	_ = UserId(432445)
	_ = Speed(54.5)
	_ = Direction(255)
	_ = Velocity(432.254)

	// strings
	var str = "Amount is $32\n hi"
	fmt.Println(str)

	// variables
	var a, b, c = 1, 5, "hello world"
	fmt.Println(a, b, c)
	var x int = 5
	var strr string = "hello"
	fmt.Println(x, strr)

	var (
		y int = 534
		n string = "sample"
		g = 432432
	)
	fmt.Println(y, n, g)

	example := 54
	h, r := 43, "helloworld"
	fmt.Println(strings.ToUpper(r), h, example)

	// scope
	// kl := 4
	// var kl = 6
	// fmt.Println(kl)
	// result === kl redeclared in this block error
}
