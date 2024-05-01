package main

func main() {
	//-------------------------数组-----------------------
	//go 集合类型1、数字、切片(slice)、map、list
	//var courses1 [3]string
	//var courses2 [4]string
	//这是两种不同的类型,应为数组的长度不同
	//[]string 和[3]string也是不同的两种数据类型，前者是切片，后者是数组，数组长度固定的性能高

	//courses1[0] = "go"
	//courses1[1] = "python"
	//courses1[2] = "java"

	//fmt.Printf("%T\r\n", courses1)
	//fmt.Printf("%T\r\n", courses2)
	//fmt.Println(courses1)

	//for index, value := range courses1 {
	//	fmt.Printf("这是第%v个元素，值为%v\n", index, value)
	//}

	//	数组的初始化
	//var courses1 = [3]string{"go", "python", "java"}
	//第二种方式
	//courses1 := [3]string{"go", "python", "java"}
	//
	//for index, value := range courses1 {
	//	fmt.Printf("这是第%v个元素，值为%v\n", index, value)
	//}
	//...代表变长
	//courses2 := [...]string{2: "go"}
	//
	//for index, value := range courses2 {
	//	fmt.Printf("这是第%v个元素，值为%v\n", index, value)
	//}
	//
	//for i := 0; i < len(courses2); i++ {
	//	fmt.Println(courses2[i])
	//}

	//var courses3 = [3]string{"go", "python", "java"}
	//
	//if courses3 == courses2 {
	//	//数组之间的比较长度要相等，并且顺序要相等
	//	fmt.Println("ok")
	//}
	//	多维数组定义
	//var courseInfo [3][4]string
	//
	//courseInfo[0] = [4]string{"go", "python", "java"}
	//courseInfo[1][0] = "hello"
	//for _, value := range courseInfo {
	//
	//	for _, column := range value {
	//		fmt.Println(column)
	//	}
	//
	//}
	//
	//for i := 0; i < len(courseInfo); i++ {
	//	for j := 0; j < len(courseInfo[i]); j++ {
	//		fmt.Println(courseInfo[i][j])
	//	}
	//}

	//---------------------------------切片-------------------------------

}
