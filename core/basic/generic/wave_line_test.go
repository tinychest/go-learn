package generic

import "testing"

// 波浪线 ~ 符号的作用

type Uint interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Uints[T Uint] []T

type MyUint uint

// 尽管 MyUInt 的 底层类型就是 uint，Uint 定义的泛型也包含 uint，但是实例化 Units 类型时，并不能以 MyUint 作为示例化的泛型类型
func TestSpecialUse(t *testing.T) {
	var _ Uints[uint]
	// var _ Uints[MyUInt] // 编译不通过 MyUInt does not implement Uint (possibly missing ~ for uint in constraint Uint)
}

type UintPref interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type UintPrefs[T UintPref] []T

// 泛型类型定义时，在基础类型前加上 ~，就行了
func TestSpecialUsePref(t *testing.T) {
	var _ UintPrefs[uint]
	var _ UintPrefs[MyUint]
}

// - 注意 ~ 后边不能是接口类型，只能是基础类型
