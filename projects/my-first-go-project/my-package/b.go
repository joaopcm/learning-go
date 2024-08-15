package myPackage

import (
	"fmt"
	"myFirstGoProject/my-package/internal/foo"
)

var Bar string = "hello, bar"

func PrintMine() {
	fmt.Println(foo.Mine)
}