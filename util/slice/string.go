package slice

import "fmt"

type StringSlice interface {
	Len() int
	Contain(string) int
	Append(string)
	Delete(string) bool
	DeleteByIndex(int)
}

type String []string

func NewString(sizes ...int) StringSlice {
	size := append(sizes, 0)[0]
	slice := make(String, 0, size)
	return &slice
}

func (s String) Len() int {
	return len(s)
}

func (s String) Contain(e string) int {
	for index, elem := range s {
		if elem == e {
			return index
		}
	}
	return -1
}

func (s *String) Append(e string) {
	*s = append(*s, e)
}

func (s *String) Delete(e string) bool {
	index := s.Contain(e)
	if index == -1 {
		return false
	}
	s.DeleteByIndex(index)
	return true
}

func (s *String) DeleteByIndex(index int) {
	l := len(*s)
	if index < 0 {
		panic("index must be non-negative")
	}
	if index > l -1 {
		panic(fmt.Errorf("index out of range[%d] with length %d", index, l))
	}
	copy((*s)[index:], (*s)[index+1:])
	*s = (*s)[:l-1]
}