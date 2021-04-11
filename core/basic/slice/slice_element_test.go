package slice

import (
	"fmt"
	"go-learn/util"
	"testing"
)

func TestSlice(t *testing.T) {
	// 删除开头的元素
	// deleteFirstTest()
	// 删除中间元素（从第 i 个元素开始，删除 N 个元素）
	// deleteMiddle()
	// 删除末尾的元素（这个很简单，取尾部的切片，样例代码不给了（任性））

	// 在循环体中，插入元素
	// insertTest()
	// 在循环中记录要插入的多个元素及其对应的位置，循环结束后一次性插入（提高效率）
	// insertFinalTest()
}

func insert(array *[]int, index int, value int) {
	length := len(*array)

	if index < 0 || index > length {
		panic("请不要搞事情")
	}

	result := make([]int, length+1)

	// 不包括结尾的元素
	copy(result, (*array)[:index])
	result[index] = value
	copy(result[index+1:], (*array)[index:])

	*array = result
}

func deleteFirstTest() {
	originArray := []int{1, 2, 3, 4, 5}

	// 方式1
	deleteFirstNElementOne(originArray, 1)
	util.PrintSliceInfo(originArray) // 失败

	// 方式2
	// 之前：length:5 capacity:5
	deleteFirstNElementTwo(&originArray, 1)
	// 之后：length:4 capacity:4
	util.PrintSliceInfo(originArray) // 成功

	// 方式3
	originArray = originArray[1:]
	util.PrintSliceInfo(originArray) // 成功

	// 方式4（copy 删除段后边的数据复制到开头 并返回复制的长度，然后切片表达式限制大小）
	originArray = originArray[:copy(originArray, originArray[1:])]
}

func deleteFirstNElementOne(array []int, n int) {
	if n <= 0 {
		panic("请不要搞事情，瞧你的 n 输入了啥")
	}
	array = array[n:]
}

func deleteFirstNElementTwo(array *[]int, n int) {
	if n <= 0 {
		panic("请不要搞事情，瞧你的 n 输入了啥")
	}

	// 错误示例 且 这样写无法通过编译
	// array = &(*array)[n:]

	// 错误示例
	// a1 := (*array)[n:]
	// array = &a1

	*array = (*array)[n:]
}

func deleteMiddle() {
	// 方式一
	// a = append(a[:i], a[i+N:]...)
	// 方式二（能随便写出这种复杂的表达式，那可就牛逼了）
	// a = a[:i+copy(a[i:], a[i+N:])]
}

func insertTest() {
	// 循环中插入元素（正序遍历）
	array := []int{1, 2, 3, 4, 5}
	indexIncr := 0
	for index, value := range array {
		insert(&array, index+1+indexIncr, value*10)
		indexIncr++
	}
	println(array)

	// 循环中插入元素（逆序遍历）
	array = []int{1, 2, 3, 4, 5}
	for index := len(array) - 1; index >= 0; index-- {
		insert(&array, index+1, array[index]*10)
	}
	println(array)
}

func insertFinalTest() {
	slice := []int{1, 2, 3, 4, 5}
	insertElements := [][2]int{{1, 10}, {2, 20}, {3, 30}, {4, 40}, {5, 50}}

	println("insert before：", slice)
	insertFinal(&slice, insertElements)
	println("insert after：", slice)
}

// 向 slice 下标为 map.key 插入 map.value
// 为什么方法名是 insertFinal，和 final 的关系：只是说这是在确定要插入的元素后，最后一次性插入，多次单元素插入的多次移位的效率不高
// 要求1：保证 insertElements 不要向指定的下标插入多个元素 - 数组越界异常
// 要求2：要插入的元素列表在数组的一定要按照从大到小的顺序
func insertFinal(slice *[]int, insertElements [][2]int) {
	sliceLength := len(*slice)
	elementLength := len(insertElements)

	// 检测 map 参数的下标是否合法
	for key, _ := range insertElements {
		if key < 0 || key > sliceLength {
			panic(fmt.Sprintf("非法的下标: %d", key))
		}
	}

	resultSlice := make([]int, sliceLength+elementLength)

	sliceIndexPtr, elementIndexPtr := 0, 0
	keyIndex, valueIndex := 0, 1
	for index, _ := range resultSlice {
		// fmt.Printf("键：%d 值：%d\n", key, value)
		if elementIndexPtr < elementLength && index == insertElements[elementIndexPtr][keyIndex]+elementIndexPtr {
			resultSlice[index] = insertElements[elementIndexPtr][valueIndex]
			elementIndexPtr++
			continue
		}
		resultSlice[index] = (*slice)[sliceIndexPtr]
		sliceIndexPtr++
	}

	*slice = resultSlice
}
