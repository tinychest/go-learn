package helloworld

import "fmt"

type helloService struct{}

func NewHelloService() IHello {
	return new(helloService)
}

func (s *helloService) Hello(args *Args, reply *Reply) error {
	fmt.Println("a new call:", args.Value)
	reply.Value = "Hello! I'm server!"
	return nil
}
