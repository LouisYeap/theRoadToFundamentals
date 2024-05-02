package main

import "fmt"

func main() {
	// map是一个key，value的键值对（无序的集合），主要是查询方便。切片和数组查询值的时候不得不去遍历性能比较低，而map直接可以拿到value通过key

	//初始化方式一：

	var courseMap = map[string]string{

		"go":     "hello",
		"java":   "hi",
		"python": "hi",
	}

	fmt.Println(courseMap)
	//取值
	fmt.Println(courseMap["go"])
	//放值
	courseMap["mysql"] = "原理"
	fmt.Println(courseMap)

	//	注意map的坑点：如果定义map结构体没有初始化的时候,map类型要设置值的时候，必须要初始化，一定要初始化加上打括号！！！！！！！！！
	//panic: assignment to entry in nil map
	//var courseMap1 map[string]string
	//courseMap1["math"] = "2"
	//fmt.Println(courseMap1)

	var courseMap1 = map[string]string{}
	courseMap1["math"] = "2"
	fmt.Println(courseMap1)

	//	方式二：make内置函数,必须初始化才能使用
	var courseMap2 = make(map[string]string, 3)
	courseMap2["mysql"] = "mysql原理"
	fmt.Println(courseMap2)

	var m []string
	if m == nil {
		fmt.Println("ok")
	}

	m = append(m, "a")

}
