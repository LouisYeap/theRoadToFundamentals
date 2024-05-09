package main

import (
	"fmt"
	"time"
)

func main() {

	//var msg chan string
	//这样会死锁堵塞，
	//msg := make(chan string, 0)
	////因为变量塞进去了没有别的协程读，就堵塞了，不会进行到下一步
	//msg <- "hello world"

	//data := msg
	//
	//fmt.Println(data)

	//解决办法
	//go 有一种head-before机制
	msg := make(chan string, 0)
	go func(msg chan string) {

		data := <-msg
		fmt.Println(data)
	}(msg)

	//因为变量塞进去了没有别的协程读，就堵塞了，不会进行到下一步
	msg <- "hello world"
	//waitgroup如果少了done调用，容易出现deadlock，无缓冲的channel也容易出现死锁，需要单独起一个协程预先消费
	time.Sleep(time.Second * 10)

}
