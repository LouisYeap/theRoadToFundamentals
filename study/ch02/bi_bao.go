package main

import "fmt"

func autoincrement() func() int {

	local := 0
	return func() int {
		local += 1
		return local
	}
}

// 闭包，函数内可以调用这个local参数，用完会自动清空
func main() {

	f := autoincrement()

	for i := 0; i < 5; i++ {
		println(f())
	}

	defer func() {
		fmt.Println("hello world")
	}()

	//	如果有多个defer，执行顺序是栈的概念，先进后出，所以先后进的defer先出

}
