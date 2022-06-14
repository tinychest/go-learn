package nil

import "testing"

// *XxxStruct(nil) 和 interface（不限 interface{}，泛指接口） 只要之间扯上关系了，那就不能直接和 nil 直接比较，而是应该和相应具体的 nil 类型进行比较
// 也就是，日常开发中，返回值为接口类型，对其进行 is nil 的判断时，需要多加思考

type MyError interface {
	error
}

type myError struct{}

func (e myError) Error() string {
	return ""
}

func Instance1() MyError {
	return nil
}

func Instance2() *myError {
	return nil
}

func Instance3() MyError {
	return (*myError)(nil)
}

func TestIsNil(t *testing.T) {
	t.Log(Instance1() == nil)
	t.Log((MyError)(nil) == nil)

	t.Log(Instance2() == nil)
	t.Log((*myError)(nil) == nil)

	t.Log(Instance3() == nil)
	t.Log(MyError((*myError)(nil)) == nil)

	// Instance3 如何才能得到 true 的比较结果
	t.Log(Instance3() == (*myError)(nil))
	// 从上面的样例可以得到结论，如果想知道原理，详见 interface/compare_test.go 的比较
}
