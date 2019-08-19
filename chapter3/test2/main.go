package main

import (
	"fmt"
	"time"
)

func main()  {
	ch := make(chan int)

	go func(c chan int) {
		defer close(c)

		for i := 0; i < 10; i++ {
			ch <- i
			time.Sleep(1 * time.Second)
		}
	}(ch)

	for v := range ch {
		fmt.Println(v)
	}
	//OUT:
	//for true {
	//	select {
	//	case v :=<- ch:
	//		fmt.Println(v)
	//	default:
	//		break OUT
	//		//time.Sleep(1 * time.Second)
	//	}
	//}

}


