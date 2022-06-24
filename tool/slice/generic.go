package slice

import "fmt"

type Data interface {
	int | int64 | float64 | string
}

type Slice[T Data] interface {
	Len() int
	Contain(T) int
	Append(T)
	DeleteByIndex(int)
}

type slice[T Data] []T

func New[T Data](sizes ...int) Slice[T] {
	size := append(sizes, 0)[0]
	list := make([]T, 0, size)
	res := slice[T](list)
	return &res
}

func (s slice[T]) Len() int {
	return len(s)
}

func (s slice[T]) Contain(e T) int {
	for index, elem := range s {
		if elem == e {
			return index
		}
	}
	return -1
}

func (s *slice[T]) Append(e T) {
	*s = append(*s, e)
}

func (s *slice[T]) DeleteByIndex(index int) {
	l := len(*s)
	if index < 0 {
		panic("index must be non-negative")
	}
	if index > l-1 {
		panic(fmt.Errorf("index out of range[%d] with length %d", index, l))
	}
	copy((*s)[index:], (*s)[index+1:])
	*s = (*s)[:l-1]
}
