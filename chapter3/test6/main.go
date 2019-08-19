package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main()  { // 多个发送者，一个接受者。接受者关闭信道
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const MaxRandomNumber = 100000
	const NumSenders = 1000

	dataCh := make(chan int, 100) // 数据信道
	stopCh := make(chan struct{}) // 关闭信道

	for i := 0; i < NumSenders; i++ {
		go func() {
			for true {
				value := rand.Intn(MaxRandomNumber)
				select {
				case <-stopCh: // 说明关闭了
					return
				default:
					dataCh <- value
				}
			}
		}()
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for c := range dataCh {
			if c == MaxRandomNumber - 1 {
				close(stopCh)
				return
			}
			log.Println(c)
		}
	}()

	wg.Wait()
}