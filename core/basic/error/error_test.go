package error

import (
	"errors"
	"fmt"
	"testing"
)

// https://studygolang.com/articles/23346?fr=sidebar

// Go 1.13 提出的 Error wrapping 概念
// wrap   指 fmt.Errorf（要求只能有一个 %w，否则，Errorf call has more than one error-wrapping directive %w）
// unwrap 指 errors.Unwrap（也就是调用 error 实例的 Unwrap 方法）

var (
	err11 = errors.New("1")
	err12 = errors.New("2 " + err11.Error())
	err13 = errors.New("3 " + err12.Error())
	err21 = errors.New("1")
	err22 = fmt.Errorf("2 %w", err21)
	err23 = fmt.Errorf("3 %w", err22)
)

func TestErr(t *testing.T) {
	// testUnwrap(t)
	// testIs(t)
	testAs(t)
}

// errors.Unwrap
func testUnwrap(t *testing.T) {
	t.Log(errors.Unwrap(err13)) // <nil>
	t.Log(errors.Unwrap(err23)) // 2 1
}

// errors.Is
func testIs(t *testing.T) {
	t.Log(err13 == err11)          // false
	t.Log(err23 == err21)          // false
	t.Log(errors.Is(err13, err11)) // false
	t.Log(errors.Is(err23, err21)) // true
}

// errors.As
type MyError struct {
	msg string
}

func (e MyError) Error() string {
	return e.msg
}

func (e MyError) Name() string {
	return e.msg + "123"
}

// errors.Unwrap 是解一层，并返回，没有实现 Unwrap 体系，则返回 nil
// errors.As 和 errors.Is 不断 errors.Unwrap 直到找到指定的容器（参数2）类型的层，并赋值给容器
func testAs(t *testing.T) {
	myErr := MyError{msg: "error1"}
	wrapErr := fmt.Errorf("errorf %w", myErr)

	// 没有执行
	if e, ok := wrapErr.(MyError); ok {
		t.Log("1 " + e.Name())
	}

	// 执行
	var e MyError
	if ok := errors.As(wrapErr, &e); ok {
		t.Log("2 " + e.Name())
	}

	// 只看类型，下面这样都 ok
	t.Log(err11)
	t.Log(errors.As(err13, &err11))
	t.Log(err11)
}