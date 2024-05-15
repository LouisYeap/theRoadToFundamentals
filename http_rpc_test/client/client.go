package main

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {

	*reply = "hello " + request
	return nil
}

func main() {

}
