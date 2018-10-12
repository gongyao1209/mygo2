package main

import (
	"mygo2/package2"
)

func main()  {
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
	package2.Test1011_01()
	//package2.Test8()

	//defer func() {
	//	fmt.Println(recover())
	//}()
	//panic("I am gongyao")

	//for i := 0; i < 5; i++ {
	//	defer func(i int) {
	//		println(i)
	//	}(i)
	//}
}

////生产者
//func Producer(ch chan int, ctx context.Context)  {
//	for i := 1; ; i++ {
//		select {
//		case <-ctx.Done():
//			return
//		case ch <- i:
//		}
//	}
//}
//
////消费者
//func Consumer(ch chan int, ctx context.Context)  {
//	for value := range ch {
//		select {
//		case <-ctx.Done():
//			return
//		default:
//			fmt.Println(value)
//		}
//	}
//}
//
//func main()  {
//	ch := make(chan int, 60)
//	ctx, cancel := context.WithCancel(context.Background()) //通过 context 来退出goroutine
//
//	go Producer(ch, ctx)
//	go Consumer(ch, ctx)
//
//	time.Sleep(2 * time.Second)
//	cancel()
//}