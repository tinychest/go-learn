package slice

import (
	"fmt"
	"go-learn/util"
	"testing"
)

/*
copy 删除段后边的数据复制到开头 并返回复制的长度，然后切片表达式限制大小
*/

func TestSlice(t *testing.T) {
	// 删除开头的元素
	// deleteFirstTest()
	// 删除中间元素（从第 i 个元素开始，删除 N 个元素）
	// deleteMiddle()
	// 删除末尾的元素（这个很简单，取尾部的切片，样例代码不给了（任性））

	// 在循环体中，插入元素
	insertTest()
	// 在循环中记录要插入的多个元素及其对应的位置，循环结束后一次性插入（提高效率）
	// insertFinalTest()
}

func insert(array *[]int, index int, value int) {
	length := len(*array)

	if index < 0 || index > length {
		panic("请不要搞事情，瞧你的 n 输入了啥")
	}

	result := make([]int, length+1)

	copy(result, (*array)[:index])
	result[index] = value
	copy(result[index+1:], (*array)[index:])

	*array = result
}

func deleteFirstTest() {
	originArray := []int{1, 2, 3, 4, 5}

	// 方式 1
	// 之前：length:5 capacity:5
	deleteHeadN(&originArray, 1)
	// 之后：length:4 capacity:4
	util.PrintSliceInfo(originArray)

	// 方式 2
	originArray = originArray[1:]
	util.PrintSliceInfo(originArray)

	// 方式 3
	originArray = originArray[:copy(originArray, originArray[1:])]
}

// array 一定的是指针类型
func deleteHeadN(array *[]int, n int) {
	if n <= 0 {
		panic("请不要搞事情，瞧你的 n 输入了啥")
	}

	// 如果这里写成对 array 直接操作，那说明值传递的概念，又忘了
	*array = (*array)[n:]
}

func deleteMiddle() {
	// 方式一（append）
	// a = append((a)[:i], (a)[i+n:]...)
	// a = (a)[:len(a)-n]
	// 方式二（copy）
	// a = (a)[:i+copy((a)[i:], (a)[i+n:])]
}

func insertTest() {
	// 循环中插入元素（正序）
	array := []int{1, 2, 3, 4, 5}
	index := 0
	for i, value := range array {
		insert(&array, i+1+index, value*10)
		index++
	}
	util.PrintSliceInfo(array)

	// 循环中插入元素（逆序）
	array = []int{1, 2, 3, 4, 5}
	for i := len(array) - 1; i >= 0; i-- {
		insert(&array, i+1, array[i]*10)
	}
	util.PrintSliceInfo(array)
}

func insertFinalTest() {
	slice := []int{1, 2, 3, 4, 5}
	insertElements := [][2]int{{1, 10}, {2, 20}, {3, 30}, {4, 40}, {5, 50}}

	util.PrintSliceInfo(slice)
	insertFinal(&slice, insertElements)
	util.PrintSliceInfo(slice)
}

// 向 slice 下标为 map.key 插入 map.value
// 为什么方法名是 insertFinal，和 final 的关系：只是说这是在确定要插入的元素后，最后一次性插入，多次单元素插入的多次移位的效率不高
// 要求1：保证 insertElements 不要向指定的下标插入多个元素 - 数组越界异常
// 要求2：要插入的元素列表在数组的一定要按照从大到小的顺序
func insertFinal(slice *[]int, insertSlice [][2]int) {
	sliceLen := len(*slice)
	elemLen := len(insertSlice)

	// 检测 map 参数的下标是否合法
	for key := range insertSlice {
		if key < 0 || key > sliceLen {
			panic(fmt.Sprintf("非法的下标: %d", key))
		}
	}

	resultSlice := make([]int, sliceLen+elemLen)

	i1, i2 := 0, 0
	keyIndex, valueIndex := 0, 1
	for i := range resultSlice {
		// fmt.Printf("键：%d 值：%d\n", key, value)
		if i2 < elemLen && i == insertSlice[i2][keyIndex]+i2 {
			resultSlice[i] = insertSlice[i2][valueIndex]
			i2++
			continue
		}
		resultSlice[i] = (*slice)[i1]
		i1++
	}

	*slice = resultSlice
}
