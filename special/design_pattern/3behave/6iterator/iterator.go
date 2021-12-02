package _iterator

// Iterator 迭代器定义
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

// Iterable 可迭代定义
type Iterable interface {
	 Iterator() Iterator
}

// IntSlice int 集合定义
type IntSlice []int

func (i IntSlice) Iterator() Iterator {
	return &intIterator{
		slice: i,
		index: 0,
	}
}

// intIterator int 集合的迭代器定义
type intIterator struct {
	slice IntSlice
	index int
}

func (i *intIterator) HasNext() bool {
	return i.index < len(i.slice)
}

func (i *intIterator) Next() interface{} {
	if !i.HasNext() {
		panic("no more")
	}

	res := i.slice[i.index]
	i.index++
	return res
}