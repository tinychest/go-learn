package create

// Go 没有构造函数，也没有方法重载，所以一般会用 Option 模式来设计一个 NewXxx 方法。
// 偶尔也会考虑 Builder 模式

type xxx struct {
	Field1 string
	Field2 int
}

/* 定义 Option 类型就是接受当前类型的方法 */

type XxxOption func(*xxx)

/* 在创建方法中，增加不定参数 opts ...Option。内部遍历 opts，应用在实例上 */

func NewXxx(opts ...XxxOption) *xxx {
	res := new(xxx)
	for _, opt := range opts {
		opt(res)
	}
	return res
}

/*
WithXxx 方法定义返回各种 Options
WithXxx 的 With 开头的方法常用于表示通过设定 Xxx 字段或者属性来创建指定的实例，如 Go 标准库中的 context，就有如 WithCancel、WithDeadline 等方法
*/

func WithField1Option(field1 string) XxxOption {
	return func(x *xxx) {
		x.Field1 = field1
	}
}
