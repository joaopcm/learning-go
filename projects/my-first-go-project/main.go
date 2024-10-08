package main

import (
	"errors"
	"fmt"
	"math"
	myPackage "myFirstGoProject/my-package"
	"os"
	"strconv"
	"time"
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

	// Loops
	loopArr := [10]int{1,2,3,4,5,6,7,8,9,10}
	for idx, value := range loopArr {
		fmt.Println(idx, value)
	}

	// Since Go 1.22
	// Iterate over an int
	for range 10 { // or "for i := range 10 {" if you want to use the index
		fmt.Println("hi!")
	}

	// IF conditions
	if x := math.Sqrt(4); x < 1 {
		fmt.Println(x)
	} else if x > 0 {
		fmt.Println("x is greater than 0")
	} else {
		fmt.Println("else statement!")
	}

	if err := doError(); err != nil {
		fmt.Println(err)
	}

	// Switches
	fmt.Println(isWeekend(time.Now()))

	switch x := math.Sqrt(4); x {
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("default")
	}

	do(1)
	do("foo")
	do(true)

	// Defer statement
	doDefer()
}

var files = []string{"foo.txt", "bar.txt", "baz.txt"}
func doDefer() {
	for _, f := range files {
		func () {
			file, err := os.Open(f)
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()
		}()
	}
}

func do(x any) {
	switch t := x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
		takeString(t)
	default:
		fmt.Println("default")
	}
}

func isWeekend(x time.Time) bool {
	switch x.Weekday() {
		case time.Sunday, time.Saturday:
			return true
		default:
			return false
	}
}

func doError() error {
	return errors.New("error")
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