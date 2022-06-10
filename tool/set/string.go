package set

type StringSet interface {
	Len() int
	Contain(string) bool
	Add(...string)
	Delete(string) bool
	ToSlice() []string
}

type String map[string]struct{}

func NewString(sizes ...int) StringSet {
	size := append(sizes, 0)[0]
	return make(String, size)
}

func (s String) Len() int {
	return len(s)
}

func (s String) Contain(e string) bool {
	_, ok := s[e]
	return ok
}

func (s String) Add(e ...string) {
	for _, elem := range e {
		s[elem] = struct{}{}
	}
}

func (s String) Delete(e string) bool {
	ok := s.Contain(e)
	delete(s, e)
	return ok
}

func (s String) ToSlice() []string {
	slice := make([]string, 0, len(s))
	for elem := range s {
		slice = append(slice, elem)
	}
	return slice
}
