package compatible

import "go-learn/special/bussiness/enum"

// 《返回值》
// Enum 不能自动转换
func r1() string {
	return enum.Male
}

// string 可以自动转换
func r2() enum.GenderEnum {
	return ""
}