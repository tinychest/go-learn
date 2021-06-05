package _map

import (
	"fmt"
	"testing"
)

func checkDuplicate() {
	idInt64Slice := []int64{1, 2, 2, 2, 3, 3, 4}

	idDuplicateSaveMap := make(map[int64]bool, len(idInt64Slice))
	idDuplicateCheckMap := make(map[int64]bool, len(idInt64Slice))
	for _, id := range idInt64Slice {
		_, ok := idDuplicateCheckMap[id]
		if ok {
			idDuplicateSaveMap[id] = true
		} else {
			idDuplicateCheckMap[id] = true
		}
	}
	if len(idDuplicateSaveMap) != 0 {
		duplicateIds := make([]int64, 0, len(idInt64Slice))
		for id := range idDuplicateSaveMap {
			duplicateIds = append(duplicateIds, id)
		}
		fmt.Printf("参数非法，重复的id：%v\n", duplicateIds)
		return
	}

	println("校验通过，元素没有重复内容")
}

// 拓展
// 无法借助 空结构体 类型
func TestEmptyValueMap(t *testing.T) {
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