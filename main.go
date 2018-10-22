package main

import (
	"context"
	"fmt"
	"mygo2/db"
	"mygo2/package1"
	"runtime"
	"sync"
	"time"
)

func main()  {

	db.Test()
	return
	i := package1.Defer1019_0101()
	fmt.Printf("Defer1019_01, addr: %p, i = %d \n", i, *i)

	//fmt.Println("-------------")
	//
	//j := package1.Defer1019_04()
	//fmt.Printf("Defer1019_04, addr: %p, j = %d \n", &j, j)

	//fmt.Println(package1.Defer1019_10())
	//fmt.Println(package1.Defer1019_11())
	//fmt.Println("Please visit http://127.0.0.1:12345/") //hello world 的革命
	//http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	//	s := fmt.Sprintf("你好, 世界! -- Time: %s", time.Now().String())
	//	fmt.Fprintf(w, "%v\n", s)
	//	log.Printf("%v\n", s)
	//})
	//if err := http.ListenAndServe(":12345", nil); err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}

	//package1.Test1()
	//package2.Test2()
	//package2.Test3()
	//package2.Test4()
	//package2.Test5()
	//package2.Test6()
	//package2.Test1012_01()
	//package2.Test1012_02()
	//package2.Test8()
	//package1.Test101202()

	//defer func() {
	//	fmt.Println(recover())
	//}()
	//panic("I am gongyao")

	//for i := 0; i < 5; i++ {
	//	defer func(i int) {
	//		println(i)
	//	}(i)
	//}

	//QuitGoroutine()

	//mianshi.Test1014_06()
	//fmt.Println(mianshi.DeffCall_1())
	//fmt.Println(mianshi.DeffCall_2())
	//fmt.Println(mianshi.DeffCall_3())

	//mianshi.Test1014_02()
}

func worker(ctx context.Context, num int, wg *sync.WaitGroup)  {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("quit goroutine", num)
			return
		default:
			fmt.Println("goroutine", num)
		}
	}
}

func QuitGoroutine()  {
	fmt.Println("goroutine num1  :", runtime.NumGoroutine())
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup

	wg.Add(1)
	go worker(ctx, 1, &wg)

	wg.Add(1)
	go worker(ctx, 2, &wg)

	fmt.Println("goroutine num2  :", runtime.NumGoroutine())

	time.Sleep(1 * time.Second)
	cancel()
	wg.Wait()

	fmt.Println("goroutine num3  :", runtime.NumGoroutine())
}

func prin(i, a, b int) int {
	sum := a + b
	fmt.Printf("%d) a = %d, b = %d, a + b = %d\n", i, a, b, sum)
	return sum
}

func main1()  {
	a := 1
	b := 2
	defer func() {
		prin(1, a, b)
	}()

	a = 3
	b = 4
	defer func() {
		prin(2, a, b)
	}()
	return
}