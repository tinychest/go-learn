package basic

import "testing"

/*
  Goland 没有红线提示，实际无法通过编译
  result parameter res not in scope at return
  result parameter err not in scope at return
*/
func TestResultParameter(t *testing.T) {
	res, err := ResultParameter()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(res)
}

func ResultParameter() (res interface{}, err error) {
	// 分析：在 Go 中对于函数语法 result parameter 的定义，
	// 意味着方法的返回值应当是定义的 result parameter 的值。
	// 但是，在下面的 if 代码块中，result parameter 的同名变量被申明了，
	// 这意味这在这个代码块中，无法访问到方法上声明的 result parameter，这样违反了 go 的语法和设计原则。
	//
	// 总结：就像编译器给出的错误提示，我们在开发时，应该遵从这样的原则，不在声明了 result parameter 的方法中，
	// 再定义任何同名的局部变量，不仅是原则上的问题，同名变量的重复定义，验中降低代码的可读性。
	// if true {
	// 	var res interface{}
	// 	var err error
	// 	_ = res
	// 	_ = err
	// 	return
	// }
	return
}
