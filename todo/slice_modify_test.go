package todo

import (
	"go-learn/util"
	"testing"
)

type Book struct {
	Name string
}

type Books []Book

func (bs Books) Modify() {
	// 对属主参数的间接部分的修改将反映到方法之外。
	bs[0].Name = "book"
}

func (bs Books) Modify2() {
	util.PrintSlice(bs)

	// 对属主参数的直接部分的修改不会反映到方法之外。
	bs = append(bs, Book{"book02"})
}

// 修改成功，违反了值属主类型的方法不能实际修改值么，NO，切片实际就是底层数组的一个地址值，这个地址值从未改变
func TestModify(t *testing.T) {
	var books = Books{{"book01"}}
	books.Modify()
	t.Log(books)
}

// TODO 添加元素失败（迷惑）
func TestModify2(t *testing.T) {
	var books = Books(make([]Book, 0, 4))
	books = append(books, Book{"book"})

	books.Modify2()
	util.PrintSlice(books)
}
