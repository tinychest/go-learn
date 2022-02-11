package helloworld

type IHello interface {
	Hello(*Args, *Reply) error
}
