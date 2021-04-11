package basic

import (
	"testing"
)

// 1、map 值的 unaddressable
// 2、方法返回值 的 unaddressable
// 有说法：大多数匿名实例都是不可寻址的
func TestAddressable(t *testing.T) {
	interfaceSlice := make([]interface{}, 0)

	var v []interface{}
	interfaceSlice = append(interfaceSlice, &v)
	// interfaceSlice = append(interfaceSlice, &getInterfaceSliceType()) // 编译不通过
}

func getInterfaceSliceType() []interface{} {
	return nil
}
