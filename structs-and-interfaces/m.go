package main

import (
	"fmt"
	"structs_and_interfaces/foo"
)

type MyString string

type User struct {
	foo.Foo

	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func (u *User) UpdateName(newName string) {
	u.Name = newName
}

func UpdateName(u *User, newName string) {
	u.Name = newName
}

type Animal interface {
	MakeSound() string
}

type Dog struct{}

func (*Dog) MakeSound() string {
	return "Woof! Woof!"
}

func (Dog) InterfaceMethod() {
	fmt.Println("Interface method called from Dog")
}

type Cat struct{}

func (c *Cat) MakeSound() string {
	return "Meow! Meow!"
}

// func whatDoesThisAnimalSay(a Animal) {
// 	fmt.Println(a.MakeSound())
// }

// func takeAnimal(a Animal) {
// 	switch t := a.(type) {
// 	case *Dog:
// 		t.MakeSound()
// 	case *Cat:
// 		t.MakeSound()
// 	}
// }

type Joao struct{}

func (Joao) String() string {
	return "This is João"
}

func main() {
	j := Joao{}
	fmt.Println(j)

	// user := User{
	// 	ID:   1,
	// 	Name: "John Doe",
	// 	Foo: foo.Foo{
	// 		Name: "Foo",
	// 	},
	// }
	// user.UpdateName("Jane Doe")
	// UpdateName(&user, "João")
	// fmt.Println(user)
	// user.Foo.Bar()

	// res, err := json.Marshal(user)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(res))

	// dog := Dog{}
	// whatDoesThisAnimalSay(dog)
	// bar.TakeFoo(dog)

	// var a Animal
	// var dog *Dog
	// fmt.Println(a == nil)
	// fmt.Println(dog == nil)
	// a = dog
	// fmt.Println(a.MakeSound())
	// fmt.Println(a == nil)
	// fmt.Println(a.MakeSound())
}
