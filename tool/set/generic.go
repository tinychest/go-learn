package set

type Data interface {
	int | int64 | float64 | string
}

type TheSet[T Data] interface {
	Len() int
	Contain(T) bool
	Add(...T)
	Delete(T) bool
	Map() map[T]struct{}
	ToSlice() []T
}

type Set[T Data] map[T]struct{}

func New[T Data](sizes ...int) TheSet[T] {
	size := append(sizes, 0)[0]
	return make(Set[T], size)
}

func NewBy[T Data](list []T) TheSet[T] {
	s := New[T](len(list))
	for _, v := range list {
		s.Add(v)
	}
	return s
}

func (x Set[T]) Len() int {
	return len(x)
}

func (x Set[T]) Contain(e T) bool {
	_, ok := x[e]
	return ok
}

func (x Set[T]) Add(e ...T) {
	for _, elem := range e {
		x[elem] = struct{}{}
	}
}

func (x Set[T]) Delete(e T) bool {
	ok := x.Contain(e)
	delete(x, e)
	return ok
}

func (x Set[T]) Map() map[T]struct{} {
	return x
}

func (x Set[T]) ToSlice() []T {
	slice := make([]T, 0, len(x))
	for elem := range x {
		slice = append(slice, elem)
	}
	return slice
}
