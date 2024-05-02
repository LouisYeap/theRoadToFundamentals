package main

import "fmt"

var age = 18
var ok bool

func main() {
	fmt.Println("我要成为go语言大师")
	println("我要成为go语言大师")

	//	变量,静态语言，和动态语言相比，差别比较大，静态语言需要先定义类型->
	/*
	   1、先定义，后使用
	   2、变量必须有类型
	   3、类型定下来之后不能改变
	*/

	var name string
	name = "louis"
	name1 := "baby"
	println(name)
	println(name1)
	println(age)
	println(ok)

	/*
		全局变量和局部变量
		变量定义必须使用
	*/
	//多变量定义,赋值和类型定义一致，变量名不能冲突，局部变量优先级大于全局变量，简洁变量不能用于全局变量，变量是有0值的
	var a int
	println(a)
	var b string
	println(b)
	var user1, user2 = "a", "2"
	println(user1, user2)

}
