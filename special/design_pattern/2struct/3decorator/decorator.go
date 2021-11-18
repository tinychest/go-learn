package _decorator

// IDraw 初级接口
type IDraw interface {
	Draw() string
}

// Square 初级实现
type Square struct {}

func (s Square) Draw() string {
	return "this is a square"
}

// ColorSquare 装饰
type ColorSquare struct {
	Square IDraw
	color string
}

func NewColorSquare(square IDraw, color string) ColorSquare {
	return ColorSquare{
		Square: square,
		color:  color,
	}
}

func (c ColorSquare) Draw() string {
	return c.Square.Draw() + ", color is " + c.color
}
