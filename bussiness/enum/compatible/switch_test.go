package compatible

import "go-learn/bussiness/enum"

// 《switch 语法》
// string 不能自动转换
func s1(p string) {
	switch p {
	case enum.Male:
	case enum.Female:
	}
}

// Enum 可以自动转换
func s2(p enum.GenderEnum) {
	switch p {
	case "MALE":
	case "FEMALE":
	}
}
