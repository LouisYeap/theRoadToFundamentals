package main

import "fmt"

func main() {

	var courseMap = map[string]string{

		"go":     "hello",
		"java":   "hi",
		"python": "hi",
	}

	fmt.Println(courseMap)

	//for key, value := range courseMap {
	//	fmt.Println(key, ":", value)
	//}
	for value2 := range courseMap {
		fmt.Println(value2)
	}
	//	map是无序的，不能保证每次打印都是有序的顺序

	//第一种 获取map中的元素，如果想要知道在不在里面,如果没有返回的是一个空的值""
	fmt.Println(courseMap["scala"])
	//第二种
	//value, ok := courseMap["scala"]

	//if !ok {
	//	fmt.Println("没有这个key")
	//} else {
	//	fmt.Println(value)
	//}

	//	可以缩写上面代码
	if value, ok := courseMap["scala"]; !ok {
		fmt.Println("not in")
	} else {
		fmt.Println(value)
	}

	//	删除一个元素,内置函数
	delete(courseMap, "python")
	fmt.Println(courseMap)
	//map 不是线程安全的，如果使用协程对一个map操作是要报错的！！！！！！！！！如果要使用的话使用sync.map

	//	go中集合了类型：
	//1、数组：不同长度的数组类型不一样、2、切片->动态数组，用起来方便，性能高3、map、4、list用的少
}
