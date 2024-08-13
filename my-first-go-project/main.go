package main

import (
	"fmt"
	myPackage "myFirstGoProject/my-package"
	"strconv"
)

var globalAge int = 21
// midName := "Pedro" // Short-syntax only works inside functions, not at the package level

func main() {
	fmt.Println("hello, world!")
	fmt.Println(myPackage.Foo)
	fmt.Println(myPackage.Bar)
	// fmt.Println(myPackage.baz) // It doesn't work because it is a private variable

	myPackage.PrintMine()
	sayHi()
	
	fmt.Println(sum(1, 2))
	fmt.Println(swap(1, 2))
	res, rem := divide(10, 3)
	fmt.Println(res, rem)
	fmt.Println(higherOrderSum(1)(2))

	// Closure vs anonymous function vs higher order function
	// Closure - function that can access variables from outside its scope
	// Anonymous function - function without a name
	// Higher order function - function that can take a function as an argument or return a function
	// Variadic function - function that can take a variable number of arguments
	anonymousFunction := func() {
		fmt.Println("anonymous function")
	}
	anonymousFunction()

	fmt.Println(variaticSum(1, 2, 3, 4, 5))

	// Variables
	firstName := "João"
	var lastName string = "Melo"
	fmt.Println(firstName, lastName, globalAge)

	// Type and constant system
	// Basic types
	// bool - true or false
	// int int8 int16 int32 int64 - signed numbers (positive and negative)
	// uint uint8 uint16 uint32 uint64 uintptr (low-level programming) - unsigned numbers (only positive numbers)
	// byte - same as uint8
	// rune - same as int32
	// float32 float64
	// complex64 complex128
	// string

	// Converting an int to a string without converting to its rune representation
	fmt.Println(formatIntToString(10))

	// Constants
	const x = 11
	takeInt32(x)
	takeInt64(x)
	takeString("foo")
	takeFloat64(x)

	// Arrays
	arr := [10]int{5: 400, 9: 11}
	fmt.Println(arr)
}

func takeFloat64(x float64) {
	fmt.Println(x)
}

func takeInt32(x int32) {
	fmt.Println(x)
}

func takeInt64(x int64) {
	fmt.Println(x)
}

func takeString(x string) {
	fmt.Println(x)
}

func sayHi() {
	fmt.Println("hi!")
}

func sum(a, b int) int {
	return a + b
}

func swap(a, b int) (int, int) {
	return b, a
}

func divide(a, b int) (res int, rem int) {
	res = a / b
	rem = a % b
	return res, rem
}

// Higher order function
func higherOrderSum(a int) func(int) int {
	// Closure
	return func(b int) int {
		return a + b
	}
}

func variaticSum(nums ...int) int {
	var out int
	for _, num := range nums {
		out += num
	}
	return out
}

func formatIntToString(num int) string {
	return strconv.FormatInt(int64(num), 10)
}