package main

import (
	"fmt"
	"time"
)

func main()  {
	var c = make(chan int, 1)

	go func() {
		//time.Sleep(time.Second * 3)
		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()

	for  {
		time.Sleep(1 * time.Second)

		select {
		case v := <- c:
			fmt.Println(v)
			if v == 0 {
				goto Gongyao
			}
		default:
			goto Gongyao
		}
	}

	Gongyao:
		fmt.Println("gongyao")
}

func rangeArr()  {
	a := [5]int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("rangeArr a. ", a)

	for k, v := range a { //a is a copy
		if k == 0 {
			a[1] = 12
			a[2] = 13
		}

		r[k] = v
	}

	fmt.Println("rangeArr a. ", a)
	fmt.Println("rangeArr r. ", r)
}

func rangeSlice()  {
	a := []int{1, 2, 3, 4, 5}
	var r [5]int

	fmt.Println("rangeSlice a. ", a)

	for k, v := range a { // 这里也是副本，但是 由于slice的结构不同
		if k == 0 {
			a[1] = 12
			a[2] = 13
		}

		r[k] = v
	}

	fmt.Println("rangeSlice a. ", a)
	fmt.Println("rangeSlice r. ", r)
}

func rangeMap()  {
	m := map[string]interface{} {
		"name":"gongyao",
		"age":26,
		"home":"shanxiheshun",
		"town":"taiyuan",
	}

	counter := 0
	for k, v := range m{ //map 是无序的 map也是传的指针

		fmt.Println(k, " ", v)

		if counter == 0 {
			delete(m, "age") //因为map是无序的，所以删除了 可能输出也可能不输出
		}
		counter ++
	}

}