package main

import (
	"fmt"
	"runtime"
)

func main()  {
	fmt.Println("hello word", runtime.GOOS)
}