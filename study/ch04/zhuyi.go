package main

import "fmt"

func mPrint(datas ...interface{}) {
	for _, data := range datas {
		fmt.Println(data)
	}
}

func main() {

	//var data = []interface{}{
	//	"hi", "hello", "你好",
	//}

	//mPrint(data)

	//	如果是这样的是可以的，但是如果以下情况就不行了

	var data = []string{
		"hi", "hello", "你好",
	}

	mPrint(data)

}
