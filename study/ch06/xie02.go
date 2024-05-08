package main

import (
	"fmt"
	"sync"
)

func main() {
	/*
		goroutine 如何通知到主线程知道子协程结束了，主的协程如何知道子的协程结束了
	*/

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("hello ,golang", i)
		}(i)

	}
	wg.Wait()

}
