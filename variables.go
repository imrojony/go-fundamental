package main

import (
	"fmt"
	"reflect"
)

//var (
//	name ="rojony"
//	course ="go course"
//	module =  3.2
//)
func main()  {
		name :="rojony"
		course :="go course"
		module :=  3.2

	fmt.Println(course)
	fmt.Println("name is set to ", name , " and is type of ", reflect.TypeOf(name))
	fmt.Println("module is set to ",module, " and is type of ",reflect.TypeOf(module))
}