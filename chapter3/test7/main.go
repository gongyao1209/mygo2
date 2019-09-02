package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main()  { // 多个使用者，多个接受者。可能都关闭信道
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)

	const MaxRandomNumber = 100000

	const NumSenders = 10 // 使用者数量
	const NumReceivers = 100 // 接受者数量

	//wg := sync.WaitGroup{}
	//wg.Add(NumReceivers)

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	for i := 0; i < NumSenders; i++ { // 信息发送者
		go func() {
			for true {
				select {
				case <-stopCh: // 信道关闭
					return
				default:
					value := rand.Intn(MaxRandomNumber)
					if value == 0 {
						fmt.Println("close1")
						close(stopCh)
						return
					}
					dataCh <- value
				}
			}
		}()
	}

	wg := sync.WaitGroup{}

	wg.Add(NumReceivers)

	for j := 0; j < NumReceivers; j++ {
		go func() {
			defer wg.Done()
			
			for true {
				select {
				case <-stopCh: // 信道关闭
					return
				case d :=<- dataCh:
					fmt.Println(d)
				default:
					value := rand.Intn(MaxRandomNumber)
					if value == 0 {
						fmt.Println("close2")
						close(stopCh)
						return
					}
				}
			}
		}()
	}

	wg.Wait()
}