package foo

import "fmt"

type Foo struct {
	Name string
}

func (Foo) Bar() {
	fmt.Println("Bar")
}

type Interface interface {
	InterfaceMethod()
}
