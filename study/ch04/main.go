package main

import "fmt"

// 结构定义
type Duck interface {
	// Gaga 方法的申请
	Gaga()
	Walk()
	Swimming()
}

// 具体实现
type pskDuck struct {

	/*可达鸭有没有腿无所谓，照样可以实现走路的方法*/
	//legs int
}

func (pd *pskDuck) Gaga() {
	fmt.Println("嘎嘎嘎嘎")

}
func (pd *pskDuck) Walk() {
	fmt.Println("走路喽")

}
func (pd *pskDuck) Swimming() {
	fmt.Println("游泳喽")

}

func main() {

	//	接口练习，理解鸭子类型，php，python
	//	go语言中，处处都是interface，到处都是鸭子类型
	//	当看到一只鸟，走起来像鸭子，游泳起来像鸭子，叫起来像鸭子，那么这只鸟就是鸭子
	//动词、方法、具备某些方法：鸭子类型强调事物的外部行为，暴露出来的方法，而不是内部的结构

	//重点，声明一个Duck类型变量，将可达鸭复制给他
	//interface是结构的定义也就是方法集，只有定义结构体，将方法绑定了,然后就可以以下赋值使用了，结构体方法的绑定相当于Java中的implement
	var duck Duck = &pskDuck{}
	duck.Walk()
	duck.Gaga()
	duck.Swimming()
}
