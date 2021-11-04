package _map

import (
	"fmt"
	"go-learn/util/set"
	"testing"
)

func TestBus(t *testing.T) {
	checkDuplicateTest()
	// emptyValueTest()
}

func checkDuplicateTest() {
	source := []int64{1, 2, 2, 2, 3, 3, 4}

	s := set.NewInt64(len(source))
	dupElemMap := make(map[int64]int, len(source))
	for _, elem := range source {
		if s.Contain(elem) {
			dupElemMap[elem]++
		}
		s.Add(elem)
	}

	if len(dupElemMap) != 0 {
		fmt.Printf("参数非法：%v\n", dupElemMap)
		return
	}
	println("校验通过，参数没有重复元素")
}

// 无法借助 空结构体 类型
func emptyValueTest() {
	theMap := make(map[string]struct{}, 0)
	// 假如要向里边存储这样含义的值
	// "1" 有
	// "2" 没有
	// "3" 有
	// 就没法存了 → 值为 2 时就不存进去呗，这样循环遍历就不会遍历到（没错是这样）
	// 但是实际场景是，判断键是否有对应的值呢 → 也没问题，不要取值，要判断第二个 bool 类型的返回参数
	if value, ok := theMap["2"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("没有")
	}
}