package _map

import "fmt"

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
