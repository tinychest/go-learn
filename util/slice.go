package util

import "fmt"

// PS：go 的 sort 包下也有如下各种类型的切片，不过主要目的是为了排序
type StrSlice []string
type IntSlice []int
type Int64Slice []int64

func (s StrSlice) Contains(theStr string) bool {
	for _, str := range s {
		if str == theStr {
			return true
		}
	}
	return false
}

func Delete(s *[]string, index int) {
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

func (s IntSlice) Contains(theI int) bool {
	for _, i := range s {
		if i == theI {
			return true
		}
	}
	return false
}

func (s Int64Slice) Contains(theI int64) bool {
	for _, i := range s {
		if i == theI {
			return true
		}
	}
	return false
}
