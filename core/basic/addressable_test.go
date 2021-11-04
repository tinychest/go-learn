package basic

import (
	"testing"
)

// unaddressable（不可寻址）：map 值、方法返回值、大多数匿名实例
func TestAddressable(t *testing.T) {
	iSlice := make([]interface{}, 0)

	var v []interface{}
	iSlice = append(iSlice, &v)
	// iSlice = append(iSlice, &getInterfaceSliceType()) // 编译不通过
}

func getInterfaceSliceType() []interface{} {
	return nil
}
