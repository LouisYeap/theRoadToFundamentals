package main

import "fmt"

type Person struct {
	name string
}

// 取指和取值
func changeName(p *Person) {
	p.name = "hello"
}
func main() {

	var a Person
	fmt.Println(a)
	//changeName(&p)

	//初始化两个关键字，map，channel，slice推荐使用make方法
	//指针初始化推荐使用new方法
	var b = new(Person)
	fmt.Println(b)

	var ps []Person
	if ps == nil {
		fmt.Println("this is a nil")
	}

	//这边是make初始化了所以不是nil
	people := make([]Person, 3)

	if people == nil {
		fmt.Println("this is aaaaa nil")
	}

}
