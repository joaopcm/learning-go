package main

import (
	"errors"
	"fmt"
)

func main() {
	// err := foo()
	// if err != nil && errors.Is(err, AnyError) {
	// 	fmt.Println(err)
	// 	return
	// }

	err := c()
	fmt.Println(err)
	fmt.Println(errors.Is(err, AnyError))
	fmt.Println(errors.Is(err, AnyError2))
}

var AnyError = errors.New("any error")
var AnyError2 = errors.New("any error 2")

func a() error { return AnyError }
func b() error { return AnyError2 }
func c() error {
	var errorResult error
	if err := a(); err != nil {
		errorResult = errors.Join(errorResult, err)
	}

	if err := b(); err != nil {
		errorResult = errors.Join(errorResult, err)
	}

	return errorResult
}

func foo() error {
	err := bar()
	if err != nil {
		return fmt.Errorf("foo error: %w", err)
	}
	return nil
}

func bar() error {
	return AnyError
}
