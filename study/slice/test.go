package main

import "fmt"

func main() {
	slice := []int{1, 2, 3}
	fmt.Println(slice) // 输出: [1 2 3]

	// 如果你有一个指向切片的指针
	slicePtr := &slice
	fmt.Println(slicePtr) // 输出: &[1 2 3]（这可能是一个内存地址的十六进制表示，但它被fmt.Println解析为切片的内容）
	// 要明确打印指针的地址，你可以使用%p格式化动词
	fmt.Printf("%p\n", slicePtr) // 输出: 0x...（这里会是切片指针的实际内存地址）
}
