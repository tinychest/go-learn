package helloworld

type helloService struct{}

func NewHelloService() IHello {
	return new(helloService)
}

func (s *helloService) Hello(args *Args, reply *Reply) error {
	reply.Value = "hello:" + args.Value
	return nil
}
