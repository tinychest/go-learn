package slice

import (
	"go-learn/util"
	"testing"
)

func TestDelElem(t *testing.T) {
	// 删除开头的元素
	// deleteFirstTest()
	// deleteHeadNTest()
	// 删除中间元素（从第 i 个元素开始，删除 N 个元素）
	// deleteMiddleNTest()
	// 删除末尾的元素（这个很简单，取尾部的切片，样例不给了（任性））
}

func deleteFirstTest() {
	originArray := []int{1, 2, 3, 4, 5}

	// 方式 1
	// 之前：length:5 capacity:5
	deleteHeadNTest(&originArray, 1)
	// 之后：length:4 capacity:4
	util.PrintSlice(originArray)

	// 方式 2
	originArray = originArray[1:]
	util.PrintSlice(originArray)

	// 方式 3
	originArray = originArray[:copy(originArray, originArray[1:])]
}

// array 一定的是指针类型
func deleteHeadNTest(array *[]int, n int) {
	if n <= 0 {
		panic("请不要搞事情，瞧你的 n 输入了啥")
	}

	// 如果这里写成对 array 直接操作，那说明值传递的概念，又忘了
	*array = (*array)[n:]
}

func deleteMiddleNTest() {
	// 方式一（append）
	// a = append((a)[:i], (a)[i+n:]...)
	// a = (a)[:len(a)-n]
	// 方式二（copy）
	// a = (a)[:i+copy((a)[i:], (a)[i+n:])]
}
