package main

import "fmt"

func main()  {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()

	fmt.Println(doinput(sum, 1, 2))

	fmt.Println(doinput(func(i int, i2 int) int {
		return i / i2
	}, 1, 0))

	panic("dagong a ")
}

func sum(a, b int) int {
	return a + b
}

func doinput(f func(int, int) int, a, b int) int {
	return f(a, b)
}