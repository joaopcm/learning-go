package main

import (
	"encoding/json"
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

func main() {
	user := User{
		ID:   1,
		Name: "John Doe",
		Foo: foo.Foo{
			Name: "Foo",
		},
	}
	user.UpdateName("Jane Doe")
	UpdateName(&user, "João")
	fmt.Println(user)
	user.Foo.Bar()

	res, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}
