package main

import "strings"

func main() {

	//匿名变量
	var _s string = "1"
	println(_s)
	//	特殊常量：可以被编译器的修改的常量
	const (
		ERROR  = iota + 1
		ERROR2 = iota + 2
		ERROR3 = iota
	)

	type IT int

	var a IT
	println(a)

	//go的stringbuilder

	var vvv strings.Builder
	vvv.WriteString("111")
	s := vvv.String()
	println(s)

}
