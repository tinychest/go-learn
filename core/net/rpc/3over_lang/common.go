package crpc

type HelloService interface {
	Hello(request string, reply *string) error
}
