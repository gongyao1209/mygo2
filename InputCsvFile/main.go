package main

import (
	"fmt"
	"mygo2/InputCsvFile/Services"
)

func main()  {
	fmt.Println(Services.ReadFile())
	fmt.Println(Services.ReadFile2())
}
