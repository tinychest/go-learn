package todo

import (
	"go-learn/tool"
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
	tool.PrintSlice(bs)

	// 对属主参数的直接部分的修改不会反映到方法之外。
	bs = append(bs, Book{"book02"})
}

// 修改成功，违反了值属主类型的方法不能实际修改值么，NO，切片实际就是底层数组的一个地址值，这个地址值从未改变
func TestModify(t *testing.T) {
	var books = Books{{"book01"}}
	books.Modify()
	t.Log(books)
}

// 注意，这里很容易产生误解；好好想想定义 切片（数组内存地址 + len + cap）
// 方法里的操作确实对数组内存空间的元素产生了影响，但是切片本身没有任何变化，还是原来的 len 和 cap
// 容易产生误解的原因就是任务切片就是底层的数组
func TestModify2(t *testing.T) {
	var books = Books(make([]Book, 0, 4))
	books = append(books, Book{"book"})

	books.Modify2()
	tool.PrintSlice(books)
}
