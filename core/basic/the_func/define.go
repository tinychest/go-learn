package the_func

type S struct {
	Name string
}

func (s *S) ptr() {
	s.Name = "ptr"
}

// 编译器自动声明
// func (*S).ptr(s *S) {
// 	s.Name = "ptr"
// }
// 原方法变为
// func (s *S) ptr() {
// 	(*S).ptr(s)
// }

func (s S) nor() {
	s.Name = "nor"
}

// 编译器自动声明
// func S.nor(s S) {
// 	s.Name = "nor"
// }
// 原方法变为
// func (s S) nor() {
// 	S.nor(s)
// }

// 为 Receiver 定义的方法，编译器会自动隐式声明一个相同的 *Receiver 的方法（为什么 *s 可以直接调用为 s 定义的方法）
// func (s *S) nor() {
// 	s.Name = "nor"
// }
// 上面只是为了好理解，实际是
// （不要有疑惑，方法声明 Receiver 带 *，实际调用的是不带 * 的 nor 方法，就是为了不能实际去改变值，因为最初的 nor 方法是为 Receiver 类型定义的，所以不能改变值）
// func (s *S) nor() {
// 	S.nor(*s)
// }

/*
针对 “编译器自动声明” 的说明，上面的自动声明都属于 aType.MethodName，不能显示命名为这样，因为这不属于合法标识符，
这样的函数只能由编译器隐式声明，但我们可以在代码中调用这些隐式声明的函数。（为什么有特殊的方法调用形式）

针对 “原方法变为” 的说明，事实上，在隐式声明上述两个函数的同时，编译器也将改写这两个函数对应的显式方法（至少，我们可以这样认为），
让这两个方法在体内直接调用这两个隐式函数。
*/

var (
	s    = S{}
	sPtr = &S{}
)
