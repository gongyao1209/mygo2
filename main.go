package main

import "mygo2/package2"

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
	package2.Test1010_03()
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
