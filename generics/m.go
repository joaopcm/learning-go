package main

import (
	"fmt"
	"slices"
)

// type MyType string
// func (MyType) Foo() {}

// type MyConstraint interface {
// 	~int | ~string | ~[]int
// }
// func foo[T MyConstraint](arg T) {
// 	fmt.Println(arg)
// }

// type MyStruct[T any] struct {
// 	Foo T
// }
// func (MyStruct[T]) foo(t T) {}

// type Foo struct {
// 	Field string
// }

// func (Foo) Do() {}

// type Bar struct {
// 	Field string
// }

// func (Bar) Do() {}

// type MyInterface interface {
// 	Bar | Foo
// 	Do()
// }

// func foo[T MyInterface](arg T) {
// 	arg.Do()
// }

func Contains[T comparable](s []T, cmp T) bool {
	for _, str := range s {
		if str == cmp {
			return true
		}
	}
	return false
}

func main() {
	// var my MyType = ""
	// foo(my)
	// foo(123)
	// foo("Hello, World!")

	// var ms MyStruct[string] = MyStruct[string]{}

	fmt.Println(Contains([]string{"a", "b", "c"}, "a"))
	fmt.Println(Contains([]string{"a", "b", "c"}, "d"))
	// OR you can use the built-in function `slices.Contains`
	fmt.Println(slices.Contains([]string{"a", "b", "c"}, "a"))
	fmt.Println(slices.Contains([]string{"a", "b", "c"}, "d"))
}
