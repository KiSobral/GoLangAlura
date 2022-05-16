package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Hugo"
	var version float32 = 1.1
	var age int = 22
	fmt.Println("Hello mr.", name, "your age is", age)
	fmt.Println("This program is under version", version)
	fmt.Println("The type of var name is", reflect.TypeOf(name))
}
