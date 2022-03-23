package compatible

import "go-learn/special/bussiness/enum"

// 《参数》
func p() {
	// Enum 不能自动转换（编译不通过）
	// p1(enum.Male)
	// string 可以自动转换
	p2("")
}

func p1(p string) {
}

func p2(p enum.GenderEnum) {
}
