package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 共享变量,尽量不是用共享变量，使用消息
//var stop bool
//var stop = make(chan struct{})

// func cpuInfo(stop <-chan struct{}) {
func cpuInfo(ctx context.Context) {
	// 这里能请求到请求的id，withValue存放一些与业务无关的参数
	defer wg.Done()
	for {

		select {
		case <-ctx.Done():
			fmt.Println("退出cpu监控")
			//	记住return
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("cpu info")
		}

		//相当于是一个机关
		//if stop {
		//	break
		//}

	}

}

func main() {

	//	context详解
	/*
			为什么要使用context？
			示例：有一个goroutine监控cpu信息，
			context	提供了三种函数，withCancel，withTimeout，withValue,
		如果你的goruntine，函数中，希望被控制，超时，传递值，但是我不希望影响我原来的接口信息的时候，函数参数中第一个参数就尽量带上
	*/
	//var stop = make(chan struct{})
	//	渐进式的方式
	wg.Add(1)
	//这是父的ctx
	//ctx, cancel := context.WithCancel(context.Background())
	////可以传递，同时生效cancel，主动取消
	//ctx2, _ := context.WithCancel(ctx)
	//go cpuInfo(ctx2)
	////defer wg.Done()
	//time.Sleep(6 * time.Second)
	//stop <- struct{}{}
	//cancel()
	//主动超时
	//ctx, _ := context.WithTimeout(context.Background(), 6*time.Second)
	//cpuInfo(ctx)

	//3.withDeadLine 	在时间点上cancel
	//4.withValue  	在时间点上cancel
	//ctx := context.WithValue(ctx, "mykey", "myvalue")
	//go cpuInfo(ctx)
	//wg.Wait()
	//fmt.Println("监控完成！")

}
