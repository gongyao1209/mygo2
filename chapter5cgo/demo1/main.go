package main

import "C"
import "fmt"

func main()  {
	//println("Hello cgo")
	//fmt.Println("hello go")
	s := C.CString("Hello, World\n")

	fmt.Println(*s)
}