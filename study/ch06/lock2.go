package main

import (
	"fmt"
	"sync"
)

/*
	我们有两组协程，其中一组负责协，另一个负责读，web大多是读多写少，
	虽然有多个协程，但是仔细分析我们会发现，读协程之间应该并发，读和写之间应该串行，读和读之间不应该并行
*/

func main() {

	var num int
	var rwlock sync.RWMutex

	wg.Add(2)
	//写的goroutine
	go func() {
		defer wg.Done()
		//加写锁，写锁时候会防止比的的锁获取，和读锁获取
		rwlock.Lock()
		defer rwlock.Unlock()
		num = 12
	}()
	go func() {
		defer wg.Done()
		//加读锁，读锁不会阻止别人读锁
		rwlock.RLock()
		defer rwlock.Unlock()
		fmt.Println(num)
	}()

	wg.Wait()

}
