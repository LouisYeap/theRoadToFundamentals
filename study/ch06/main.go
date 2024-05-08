package main

import (
	"fmt"
	"time"
)

func asyncPrint() {

}

func main() {

	/*
		多线程，内存占用大，线程切换成本高，gvm操作操作系统，内存大，并发要求高，web2.0，解决代价高，用户级线程，协程
		协程切换快，内存占用小，2k一个go语言没有线程，go语言诞生之后只有协程，goroutine，协程生态好，非常方便
		主协程挂了，子协程也会挂了
		特性：主死随从
		for循环，i变量会被重用，当第二轮时候i改变了异步过程，不能等到你i改变再循环改变地址了，是异步的!!!!!!,所以可能不能输出完整的，解决办法，将i变量赋值复制，再输出，第二种办法，值传递复制
	*/

	for i := 0; i < 10; i++ {
		//time.Sleep(1 * time.Second)

		//tmp := i
		go func(i int) {
			fmt.Println("hello", i)
		}(i)
	}

	time.Sleep(1 * time.Second)

}
