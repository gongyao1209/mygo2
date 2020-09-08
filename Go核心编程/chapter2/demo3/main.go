package main

import "fmt"

func main()  {
	a := 1
	fmt.Println(&a, a)

	f := fu(a)
	g := fu(2)

	fmt.Println(f(1))
	fmt.Println(f(2))

	fmt.Println(g(1))
	fmt.Println(g(2))
}

func fu(a int) func(i int) int {
	return func(i int) int {
		fmt.Println(&a, a)
		return a + i
	}
}