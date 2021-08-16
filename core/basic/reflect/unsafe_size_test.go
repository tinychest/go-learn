package _reflect

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func TestUnsafeSize(t *testing.T) {
	// 1、int；8
	fmt.Println(unsafe.Sizeof(reflect.Int))
	// 2、rune = Unint8：8
	fmt.Println(unsafe.Sizeof(reflect.Uint8))
	// 3、string：虽然支持该类型，但是并得不到具体的类型大小
	size := unsafe.Sizeof("012345678901234567890123123123123123123")
	fmt.Println(size)
	// 4、切片（无论 len cap 是多少，这里返回的都是常量）
	fmt.Println(unsafe.Sizeof(make([]string, 2, 4)))
}
