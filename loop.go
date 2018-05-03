//package main

//import (
//	"fmt"
//)
//
//func main() {
//
//	for i:= 10; i>=0; i-- {
//		if i== 0{
//			fmt.Println("boom")
//
//		}
//		fmt.Println(i)
//	}
//}
//package main
//
//import "fmt"
//
//func main() {
//	sum := 0
//	for i := 0; i < 10; i++ {
//		sum += i
//	}
//	fmt.Println(sum)
//}
package main

import "fmt"

func main() {
	sum := 1
	for ; sum < 6; {
		sum += sum
	}
	fmt.Println(sum)
}
