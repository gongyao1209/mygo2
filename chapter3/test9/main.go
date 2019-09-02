package main

import (
	"fmt"
)

func main()  {

	s := GetArr()
	//s := 1
	fun2(s)

	//time.Sleep(1 * time.Second)
	//fmt.Printf("s type is %T\n", s)
}

func GetArr() []int {
	s := make([]int, 0, 1000)
	for i := 0; i < 100; i++ {
		s = append(s, i)
	}
	return s
}

func fun1(arr []int) {
	for a := range arr {
		fmt.Println(a + 1)
	}
}

func fun2(arr []int) {
	myCh := make(chan int, 3)

	for i := 0; i < 4; i++ {
		go func(ch chan int) {
			for true {
				a :=<- ch
				fmt.Println(a + 1)
			}
		}(myCh)
	}

	for a := range arr {
		myCh <- a
	}
}