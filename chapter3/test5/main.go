package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main()  {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const MaxRandomNumber = 100000
	const NumReceivers = 100

	wg := sync.WaitGroup{}
	wg.Add(NumReceivers)

	ch := make(chan int, 100)

	// 写入
	go func() {
		for true {
			if value := rand.Intn(MaxRandomNumber); value == 0 {
				close(ch)
				return
			} else {
				ch <- value
			}
		}
	}()

	// 读出
	for i := 0; i < NumReceivers; i++ {
		go func() {
			defer wg.Done()

			for va := range ch {
				log.Println(va)
			}
		}()
	}

	wg.Wait()
}