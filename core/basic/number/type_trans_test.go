package number

import "testing"

func TestTransType(t *testing.T) {
	a := 100 * 0.02

	// 将 100 和 0.02 拆分就不行（mismatched types int and float64）
	// a1 := 100
	// a2 := 0.02
	// a = a1 * a2

	// 猜测：像上面直接写数字，go 有帮忙做默认的类型转化 100（int） → 100（float64），但是你自己写定类型，就不行了
	println(a)
}
