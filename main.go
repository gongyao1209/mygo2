package main

import (
	"context"
	"fmt"
	"mygo2/freeWheel"
	"runtime"
	"sync"
	"time"
)

func main()  {
	freeWheel.TestReflect()
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


func twoSum(nums []int, target int) []int {
	temp := make(map[int]int)
	res := make([]int, 2)

	for k , v := range nums {
		if _, ok := temp[v]; !ok {
			temp[v] = k
		}
	}

	for k , v := range nums {
		t := target - v

		if a, ok := temp[t]; ok {
			if k == a {
				continue
			}

			res[0] = k
			res[1] = a

			return res
		}
	}

	return res
}

