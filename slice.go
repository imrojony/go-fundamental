package main

import "fmt"

func main() {
	//myCourse := make([] string, 5, 10)
	myCourse := [] string{"romony","saleha","sunnat","romony","sellim"}
	fmt.Printf("length is : %d \n capacity is :%d ", len(myCourse),cap(myCourse))

	myNumbers := []int{1,45,678}
	fmt.Printf("the length is :%d\n the capacity is :%d",len(myNumbers),cap(myNumbers))
	}
