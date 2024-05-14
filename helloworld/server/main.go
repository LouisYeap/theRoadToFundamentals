package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {

	*reply = "hello" + request
	return nil

}
func main() {
	/*	1.实例化server
		2.注册处理逻辑
		3.启动服务
	*/
	listen, err := net.Listen("tcp", "1234")
	if err != nil {
		return
	}

	//实例化名字和对象，注册rpc
	err = rpc.RegisterName("rpc_name", &HelloService{})
	if err != nil {
		return
	}
	//当一个新的连接进来时候
	accept, err := listen.Accept()
	if err != nil {
		return
	}
	rpc.ServeConn(accept)

}
