package main

import (
	"errors"
	"fmt"
)

func A() (int, error) {

	//panic("this is a panic")
	fmt.Println("错误")
	return 0, errors.New("this is a error")
}

func main() {
	//	开发函数的人需要又一个返回值去告诉是否调用成功，go设计者我们必须要处理这个err，代码中大量会出现if err != nil ，go设计被认为必须要处理这个error，防御编程。健壮性很好
	//
	//_, err := A()
	//if err != nil {
	//	fmt.Println(err)
	//}

	//	panic：会导致程序退出，程序就会挂掉，在go语言中不推荐随便使用panic，开发中很少会用到，一般我们服务启动的过程中，比如我的服务想要必须有些依赖服务，比如日志，mysql文件没有
	//	问题，如果服务启动过程中发现任何一个不满足，就主动调用panic，但是你的程序一旦启动了，这个时候你的某行代码中，写了panic，那整个程序都会挂了，就会造成重大事故
	//	但是架不住，但是有些地方代码写的不小心，被动触发panic

	//比如以下代码，map没有初始化，被动造成panic，这个时候recover函数出现了，去捕获panic
	defer func() {
		//捕获异常信息
		if re := recover(); re != nil {
			//这边就可以对函数造成的异常捕获然后输出打印了,捕获异常时候要写在前面，不然先panic就直接报错了
			fmt.Println("recover捕获了异常:", re)
		}
	}()

	var names map[string]string

	names["go"] = "go开发工程师"
	fmt.Println("hello")

}
