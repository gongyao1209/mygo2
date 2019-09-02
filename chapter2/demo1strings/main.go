package main

import (
	"fmt"
	"strings"
)

func main()  {

	str := "gongyao"

	fmt.Println(str)

	fmt.Println(strings.Index(str, "n"))

	fmt.Println(strings.Split(str, "n"))
}