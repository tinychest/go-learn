package slice

import (
	"fmt"
	"go-learn/tool"
	"testing"
)

func TestInsertElem(t *testing.T) {
	// 在循环体中，插入元素
	// insertTest()
	// 在循环中记录要插入的多个元素及其对应的位置，循环结束后一次性插入（提高效率）
	finalInsertTest()
}

func insert(dst *[]int, index int, value int) {
	length := len(*dst)

	if index < 0 || index > length {
		panic("请不要搞事情，瞧你的 n 输入了啥")
	}

	result := make([]int, length+1)

	copy(result, (*dst)[:index])
	result[index] = value
	copy(result[index+1:], (*dst)[index:])

	*dst = result
}

func insertTest() {
	// 循环中插入元素（正序）
	dst := []int{1, 2, 3, 4, 5}
	index := 0
	for i, value := range dst {
		insert(&dst, i+1+index, value*10)
		index++
	}
	tool.PrintSlice(dst)

	// 循环中插入元素（逆序）
	dst = []int{1, 2, 3, 4, 5}
	for i := len(dst) - 1; i >= 0; i-- {
		insert(&dst, i+1, dst[i]*10)
	}
	tool.PrintSlice(dst)
}

func finalInsertTest() {
	dst := []int{1, 2, 3, 4, 5}
	src := [][2]int{{1, 10}, {2, 20}, {3, 30}, {4, 40}, {5, 50}}

	tool.PrintSlice(dst)
	finalInsert(&dst, src)
	tool.PrintSlice(dst)
}

// finalInsert 向 slice 下标为 map.key(int) 的位置插入 map.value(slice)；先确定位置，最后一次性插入，所以叫 final
func finalInsert(dstPtr *[]int, src [][2]int) {
	dst := *dstPtr
	dstLen := len(dst)
	srcLen := len(src)

	// 检测要插入的下标是否合法
	for _, item := range src {
		if item[0] < 0 || item[0] > dstLen {
			panic(fmt.Sprintf("非法的下标: %d", item[0]))
		}
	}

	i1, i2 := 0, 0
	ki, vi := 0, 1
	result := make([]int, dstLen+srcLen)

	for i := range result {
		// fmt.Printf("键：%d 值：%d\n", key, value)
		if i2 < srcLen && i == src[i2][ki]+i2 {
			result[i] = src[i2][vi]
			i2++
			continue
		}
		result[i] = (dst)[i1]
		i1++
	}

	*dstPtr = result
}
