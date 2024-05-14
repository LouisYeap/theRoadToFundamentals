package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	Name    string `json:"name"`
	Company string `json:"company"`
}

type PrintResult struct {
	Info string `json:"info"`
	Err  error  `json:"err"`
}

func RpcPrintLn(e Employee) {

	/*
		1.建立连接，基于tcp，http，udp等等
		2.将对象序列化成json串
		3.发送json字符串->调用成功后是二进制的数据，
		4.等待服务器结果
		5.反序列化

		1.监听端口
		2.读取二进制数据
		3.解码二进制协议数据
		4.处理业务，将对象序列化
		5.将数据返回
	*/
	//2.6 start time

}

func main() {

	employee := Employee{
		Name:    "John",
		Company: "John Company",
	}
	fmt.Println("hello world")

	fmt.Println(employee)

	emStr, _ := json.Marshal(employee)
	fmt.Println(string(emStr))

	json.Unmarshal([]byte(emStr), &employee)

	fmt.Println(employee)

}
