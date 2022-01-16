package error

import (
	"errors"
	"testing"
)

type MyError1 struct {
	s string
}

func (e MyError1) Error() string {
	return "error 01"
}

type MyError2 struct {
	error
}

func (e MyError2) Error() string {
	return e.error.Error()
}

func (e MyError2) Unwrap() error {
	return e.error
}

func TestErrIs(t *testing.T) {
	var e1 error = MyError1{}
	var e2 error = MyError2{e1}

	t.Log(e1 == e2)          // 呃，当然为 false
	t.Log(errors.Is(e2, e1)) // false

	// 问：有什么办法可以让上面从逻辑上说 “相同” 的错误，进行比较，返回 true
	// 答：让 MyError2 实现 Wrapper 接口
	// 注意：参数的顺序有要求的

	// 后续（为 MyError2 实现了 Wrapper 接口后）
	t.Log(e1 == e2.(MyError2).Unwrap()) // true
	t.Log(e1 == errors.New("error 01")) // false

	re := e2.(MyError2).Unwrap()
	// 假如有面试题，地址不相同的两个实例，通过 == 比较，一定为 false 么，如果，不是这样的，举个例子
	t.Logf("%p %p %t\n", &e1, &re, e1 == re)
	// 其实这个本质是
	t.Log(MyError1{} == MyError1{})
}

func TestErrIs2(t *testing.T) {
	// 再仔细想想，这个题，问的有问题
	var i1 = 1
	var i2 = 1
	t.Logf("%p %p %t\n", &i1, &i2, i1 == i2)

	// 再想想，好像还有不对的地方，下面是实际的例子
	var e1 error = errors.New("123")
	var e2 error = MyError2{e1}

	t.Log(e1 == e2)
	t.Log(errors.Is(e2, e1))

	re := e2.(MyError2).Unwrap()
	t.Logf("%p %p %t\n", &e1, &re, e1 == re)

	// 接着测试
	t.Log(errors.New("123") == errors.New("123")) // false

	var e = errors.New("123")
	var ee = e
	t.Log(e == ee) // true

	// 见 errors.New 源码，点1 - 摸得到的实例，点2 -地址
}

func TestErrIs3(t *testing.T) {
	// 继续引出（本以为是不相等的，没事了）
	t.Log(MyError1{"123"} == MyError1{"123"})

	// 其实，想引出的是，error.New 返回的是地址，实现 Error 方法类型的结构体类型是带指针的
	// TODO 不带指针也没有问题
}
