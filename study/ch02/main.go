package main

import (
	"fmt"
	"time"
)

// 函数参数传递的时候，值传递，go语言中全部都是值传递
func add(a *int, b int, c float64) (sum int, err error) {
	*a = *a + 100
	return *a + b + int(c), nil
}

func runForever() {
	for {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("doing")
	}

}
func add2(a ...int) (sum int, err error) {
	for _, i := range a {
		sum += i
	}
	return sum, err
}

func cal(op string, items ...int) func() {
	switch op {
	case "+":
		return func() {
			fmt.Println("+法")
		}
	case "-":

		return func() {
			fmt.Println("-法")
		}
	default:
		return func() {
			fmt.Println("不是+也不是-")
		}

	}

}
func main() {
	//go函数支持普通函数，匿名函数和闭包
	//go中函数是一等公民，1、函数本身可以当作变量，2、匿名函数，闭包，3、函数可以满足接口

	//a := 1
	//b := 2
	//i, err := add(&a, b, 123.9)
	//if err != nil {
	//	return
	//}
	//fmt.Println(i)
	//fmt.Println(a)
	////runForever()
	//a, err = add2(a, b)
	//fmt.Println(a, err)

	//获取函数变量，不能加（）
	//
	//funVar := add
	//sum, _ := funVar(&a, 2, 2.1)
	//fmt.Println(sum)

	//返回是函数，所以说需要再加()运行
	cal("+", 1, 2, 3)()
	//	明确计划和目标，python满足日常工作需求就可以，go作为长期发展语言

}
