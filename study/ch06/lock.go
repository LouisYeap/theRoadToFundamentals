package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total int32
var wg sync.WaitGroup
var lock sync.Mutex

//锁能复制吗？是可以的，但是不要去复制，复制就失去了锁的效果，锁也是结构体，state int32 sema  uint32 通过改变这两个属性实现锁的原理
//第二种方法，原子操作方法,lock可以锁很多行代码，可以锁代码段，atomic不行
/*
对于共享资源是有竞争关系的
golang启动！
*/
func add() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		//lock.Lock()
		//total += 1
		//lock.Unlock()
		atomic.AddInt32(&total, +1)
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		//lock.Lock()
		//total -= 1
		//lock.Unlock()
		atomic.AddInt32(&total, -1)
	}
}
func main() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(total)

}
