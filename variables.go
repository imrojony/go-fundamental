package main

import (
	"fmt"
	"reflect"
)

var (
	name, course , module = "rojony","go course", 3.2
)
func main()  {
	fmt.Println("name is set to ", name , " and is type of ", reflect.TypeOf(name))
	fmt.Println("module is set to ",module, " and is type of ",reflect.TypeOf(module))
}