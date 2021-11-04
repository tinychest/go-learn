package set

type Int64Set interface {
	Len() int
	Contain(int64) bool
	Add(...int64)
	Delete(int64) bool
	ToSlice() []int64
}

type Int64 map[int64]struct{}

func NewInt64(sizes ...int) Int64Set {
	size := append(sizes, 0)[0]
	return make(Int64, size)
}

func (s Int64) Len() int {
	return len(s)
}

func (s Int64) Contain(e int64) bool {
	_, ok := s[e]
	return ok
}

func (s Int64) Add(e ...int64) {
	for _, elem := range e {
		s[elem] = struct{}{}
	}
}

func (s Int64) Delete(e int64) bool {
	ok := s.Contain(e)
	delete(s, e)
	return ok
}

func (s Int64) ToSlice() []int64 {
	slice := make([]int64, 0, len(s))
	for elem := range s {
		slice = append(slice, elem)
	}
	return slice
}
