package main

import "fmt"

func main() {
	x := create()
	fmt.Println(*x)
	take(x)
	fmt.Println(*x)
}

func create() *int {
	x := 10
	return &x
}

func take(x *int) {
	*x = 100
}
