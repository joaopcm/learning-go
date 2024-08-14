package main

import (
	"errors"
	"fmt"
)

type User struct {
	foo string
}

func (u User) Foo() {
	fmt.Println(u.foo)
}

func NewUser(wantError bool) (*User, error) {
	if wantError {
		return nil, errors.New("error")
	}
	return &User{}, nil
}

func main() {
	user, err := NewUser(true)
	if err != nil {
		fmt.Println(err)
		return
	}
	user.Foo()

	// a := 10
	// b := 0
	// res, err := divide(a, b)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(res)
}

// func divide(a, b int) (int, error) {
// 	if b == 0 {
// 		return 0, errors.New("cannot divide by zero")
// 	}

// 	return a / b, nil
// }
