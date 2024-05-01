package main

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
	for _, s := range strings {
		println(s)
	}
	//	不能往没有定义场地的切片通过var a[]string ,然后a[0] ="hello",会报错，因为切片底层还是数组，找不到对应的空间，要使用append，或者make

}
