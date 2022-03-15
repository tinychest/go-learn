package typ

import (
	"testing"
)

// 方法值的正规化
// 这是一个比较重要的概念（详见 Go101 方法/方法值的正规化）

type Book struct {
	Name string
}

func (b Book) Pages() string {
	return b.Name
}

func (b *Book) Pages2() string {
	// 效果一样
	// return b.Name
	return (*b).Pages()
}

func TestEvaluate(t *testing.T) {
	var b = Book{"book01"}
	var p = &b

	var f1 = b.Pages  // 已是正规化，在运行时刻对其进行估值时，属主实参 b 的一个副本将被存储下来
	var f2 = p.Pages  // 正规化为 (*p).Pages，同理，进行了估值
	var g1 = p.Pages2 // 已是正规化，属主实参 p 的一个副本将被存储下来，此副本的值为 b 值的地址；当 b 被修改后，此修改可以通过对此地址值解引用而反映出来
	var g2 = b.Pages2 // 正规化为 (&p).Pages2，同理

	b.Name = "book02"
	t.Log(f1()) // book01
	t.Log(f2()) // book01
	t.Log(g1()) // book02
	t.Log(g2()) // book02
}
