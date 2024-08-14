package bar

import "structs_and_interfaces/foo"

func TakeFoo(i foo.Interface) {
	i.InterfaceMethod()
}
