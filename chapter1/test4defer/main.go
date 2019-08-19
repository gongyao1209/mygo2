package main

import "fmt"

func main()  {
	c := test()

	fmt.Println(c)
}

func test() (close bool) {
	defer func() {
		close = false
	}()

	close = true
	return
}