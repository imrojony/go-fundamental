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



	fmt.Println("name is set to ", name , " and is type of ", reflect.TypeOf(name))
	fmt.Println("module is set to ",module, " and is type of ",reflect.TypeOf(module))
	changename(&name)
	fmt.Println(name)

	changecousre(&course)
	fmt.Println(course)
}
func changecousre(course *string)  string{
	*course ="php course"
	fmt.Println(*course)
	return *course
}
func changename(name *string)  string{
	*name="romony"
	fmt.Println(*name)
	return *name

}