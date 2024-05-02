package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func updateSlice(a []string) {
	a[0] = "hello" //对原来数据发生改变

	//对原来数据没有发生改变
	for i := 0; i < 10; i++ {
		a = append(a, strconv.Itoa(i))
	}
}

//slice底层是结构体

type slice struct {
	array unsafe.Pointer //用来存放世纪数据的数组指针，指向一块连续的内存
	len   int            //切片中元素的数量
	cap   int            //array数组的长度
}

func main() {

	//	go的slice在函数参数传递的时候，是值传递还是引用传递，效果又呈现出引用传递的效果，但不完全是,其实是值传递

	//ms := slice{
	//	len: 3,
	//	cap: 3,
	//}

	//a := []string{"hi", "h2"} //底层是结构体赋值，这是语法糖写法
	//updateSlice(a)
	//fmt.Println(a)
	//	在底层中，在函数参数传递时候，会拷贝一份slice，但是注意，他们的头指针指向同一个数组，导致引用指向同一个数据结构，这是注意的点
	//扩容会重新分配，指向不同的数据结构，再改变数据，就不会对原来数据结构改变，这也就是为什么append接收一个返回的参数值，append超级有可能引发扩容，地址会发生改变

	//	下面举一个简单的例子

	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := data[1:6]
	s2 := data[2:7]
	fmt.Println(len(s1), cap(s1))
	fmt.Println(len(s2), cap(s2))
	//没有扩容时候指向同一个数据结构，改变的时候一起改变

	//当发生append导致扩容时候，会重新分配内存，指向不同的数据结构
	s2 = append(s2, 1, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3)
	s2[0] = 12211

	fmt.Println(s1)
	fmt.Println(s2)

}
