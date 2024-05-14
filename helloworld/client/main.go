package main

import (
	"fmt"
	"net/rpc"
)

func main() {

	/*
		1、建立连接
		2、
	*/
	dial, err := rpc.Dial("tcp", "8080")
	if err != nil {
		return
	}
	var reply *string = new(string)
	err = dial.Call("HelloService.Hello", "bobby", reply)
	if err != nil {
		return
	}
	fmt.Println(*reply)

}
