package main

import (
	"fmt"
	"runtime"
)

func main()  {

	fmt.Println(runtime.GOROOT())
	fmt.Println(runtime.GOMAXPROCS(1))

	fmt.Println("Hello World")
}
