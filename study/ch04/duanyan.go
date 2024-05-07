package main

import "fmt"

func add(a, b interface{}) interface{} {

	switch a.(type) {
	case int:
		ai, _ := a.(int)
		bi, _ := b.(int)
		return ai + bi

	case float64:
		ai, _ := a.(float64)
		bi, _ := b.(float64)
		return ai + bi
	default:
		return "什么都不是"
	}

}

func main() {
	//这里有个问题，如下，switch进入int float直接转换为默认值0了，所以还是用范形
	var a int = 1
	var b float64 = 1.1
	re := add(a, b)
	fmt.Println(re)
}
