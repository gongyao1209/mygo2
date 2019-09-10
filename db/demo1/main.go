package main

import (
	"fmt"
	"gocode/mygo2/db"
	"sync"
)

func main()  {

	var wg sync.WaitGroup

	i := 0
	Xunhuan:
	for true {
		if i == 11 {
			break Xunhuan
		}
		wg.Add(1)
		go func() {
			defer wg.Done()

			s := db.GetData(1)
			fmt.Println(s)
		}()


		i++
	}
	wg.Wait()
	//s := db.GetData(1)
	//s2 := db.GetData(3)
	//fmt.Println(s, s2)
}