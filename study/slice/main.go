package main

import "fmt"

func main() {

	//	切片相当于python的list:列表,go疯狂在动态语言和静态语言中摇摆,注意不能放长度
	//var courses []string
	//
	////fmt.Printf("%T", courses)
	////在 Go 语言的 append 函数签名中，Type 并不是一个具体的类型，而是一个占位符，表示你可以使用任何类型来替代它。这是为了说明 append 函数可以接收任意类型的切片（slice）以及任意数量的同类型元素作为参数。
	//
	//courses = append(courses, "hello", "hi")
	////fmt.Printf("%v", courses)
	//
	//for _, course := range courses {
	//	fmt.Printf("%v\n", course)
	//}
	//	切片初始化，3种方式，1、从数组创建，2、使用string{}，3、make
	//var allCourses = []string{"h1", "h2", "h3", "h4"}
	////左闭右开的[},python的用法
	//allCourses = allCourses[0:2]
	//for _, course := range allCourses {
	//	println(course)
	//}
	//创建预分配的切片，后续可以自动扩容,make创建的时候性能比较高
	strings := make([]string, 3)
	strings[0] = "hello"
	//println(strings)
	fmt.Println(strings)
	//for _, s := range strings {
	//	println(s)
	//}
	//	不能往没有定义长度的切片通过var a[]string ,然后a[0] ="hello",会报错，因为切片底层还是数组，找不到对应的空间，要使用append，或者make
	//	访问切片的	元素
	//strings[:]

	//	重要！！！！
	//切片中添加切片使用...将切片打散，切片展开或切片解包切片后面跟...
	a := []string{"go", "python"}
	//b := []string{"java", "python"}
	//c := append(a, b...)
	//fmt.Println(c)
	//println(a...)
	fmt.Println(a)
	//将变量赋值的话他们的地址没有发生改变，指向同一个地址
	e := a
	f := a
	fmt.Printf("%p,%p,%p\n", e, f, a)

	//	复制
	//copyCourse := a
	//fmt.Println(copyCourse)
	//注意 拷贝的对象初始一定有长度，可以用make关键字，不然拷贝是[]，拷贝地址是不一样的
	var c = make([]string, len(a))
	copy(c, a)
	fmt.Println(c)

	//1-9开始
}
