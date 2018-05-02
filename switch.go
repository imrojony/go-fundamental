package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//switch "docker" {
	//case "linux":
	//	fmt.Println("this is linux course ")
	//case "ubuntu":
	//	fmt.Println("this is ubuntu course")
	//case "docker":
	//	fmt.Println("this docker course")
	//	fallthrough
	//default:
	//	fmt.Println("there is no course")
	//}
	switch tmp:= randomNumber(); tmp {
	case 1,2,3,10:
		fmt.Println("number fo10und in first block")
		fmt.Println(tmp)
	case 5,7,9:
		fmt.Println("number found in second block")
		fmt.Println(tmp)
	default:
		fmt.Println("number not found")
		fmt.Println(tmp)

	}
}
func randomNumber()  int{
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)


}