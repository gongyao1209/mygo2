package main

import "fmt"

func main()  {

	//defer func() {
	//	defer fmt.Println("defer3")
	//	fmt.Println("defer1")
	//}()
	//
	//defer func() {
	//	fmt.Println("defer2")
	//}()
	//
	//fmt.Println("function")

	a := f
	a = f2
	fmt.Println(a())
}

func f() int {
	a := 1

	defer func(i int) { // defer 必须先注册才能执行
		fmt.Println("defer i: ", i)
	}(a)

	a++

	return a
}

func f2() int {
	a := 1
	defer fmt.Println("defer a: ", a)

	a++
	return a
}