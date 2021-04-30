package main

import "fmt"

func main()  {
	a := 1
	b := 2
	defer func() {
		test(a, test(a, b))
	}()

	c := make(map[string]string)
	c["aa"] = "bbb"
	fmt.Println(c)

	fmt.Println("aaa")
}

func test(a, b int) int {
	ret := a + b
	fmt.Println("test, a:", a, " ", "b:", b)
	return ret
}