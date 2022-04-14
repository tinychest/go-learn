package generic

import "testing"

type TheInt interface {
	int | int8 | int16 | int32 | int64
}

type TheUint interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type TheFloat interface {
	float32 | float64
}

type TheNormal interface {
	Number | string
}

// [并集]
type AllInt interface {
	TheInt | TheUint
}

type AllNumber interface {
	AllInt | TheFloat
}

// [交集]
type Intersection interface {
	TheNormal
	string
}

type IS[T Intersection] []T

func TestIntersection(t *testing.T) {
	var _ IS[string]
	// var _ IS[int] // 编译不通过
}

// [空集合] 合乎语法，但并没有什么意义
type Empty interface {
	int
	string
}

// [空接口] 本来这块不应该放在当前文件中，但是接着说，上下文语意通畅
// interface{} 空接口等同于上面的空集合么？答案是否定的，空接口依旧代表所有、任意类型，泛型中也体现出这个意思
