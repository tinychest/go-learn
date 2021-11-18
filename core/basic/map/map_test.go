package _map

import (
	"fmt"
	"testing"
)

// 关于 map 的缩容：https://mp.weixin.qq.com/s/Slvgl3KZax2jsy2xGDdFKw

func TestMap(t *testing.T) {
	// testZeroValue()
	testExist()
	// testLength()
	// testMapAddressConcept()
	// testStructKey()
	// TraversingMap()
	// testRemove()
}

// 遍历的随机性
func TraversingMap() {
	testMap := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5, 6: 6}
	for key, value := range testMap {
		println(key, value)
	}
}

// 零值
func testZeroValue() {
	var uninitializedMap map[int]int
	// "true"
	fmt.Println(uninitializedMap == nil)
	// "map[]"
	fmt.Println(uninitializedMap)
}

// 删除（并没有真实释放空间）
func testRemove() {
	var theMap = map[string]string{
		"name": "xm",
	}
	fmt.Printf("%v\n", theMap)

	delete(theMap, "name")
	fmt.Printf("%v\n", theMap)
}

// 判断 map 中是否有指定键对应的值，只能通过第 2 个返回值去判断，因为即使没有对应的键值，map 也会返回值类型对应的零值
func testExist() {
	m := map[string]struct{}{
		"123": {},
	}

	v1, ok1 := m["123"]
	v2, ok2 := m["456"]

	fmt.Println(v1, v2, v1 == v2)
	fmt.Println(ok1, ok2)
}

// len 的概念（map 没有 cap 的概念）
func testLength() {
	// len: 0
	var theMap map[int]int
	fmt.Printf("len: %d\n", len(theMap))

	// len: 0
	// make 创建map的第二个参数相当于设置 map 的容量，没错，但是你却不能通过 cap 查看 map 的容量（个人理解的是 map 的数据结构较为复杂，不能说用一个整型数字来表示 map 容量的概念）
	theMap = make(map[int]int, 2)
	fmt.Printf("len: %d\n", len(theMap))

	// len: 3
	// len 代表 map 中有多少个键值对（同切片，超过默认的容量，仍旧可以向里边添加，有默认的扩容行为）
	// 为 nil 的未初始化的 map 可无法添加数据
	theMap[1] = 1
	theMap[2] = 2
	theMap[3] = 3
	fmt.Printf("len: %d\n", len(theMap))

	// len: 0
	theMap = make(map[int]int) // make(map[int]int, 0)
	fmt.Printf("len: %d\n", len(theMap))
}

// 内存地址
func testAddress() {
	getMapFunc := func() map[string]string {
		theMap := map[string]string{"name": "小明", "gender": "男"}
		fmt.Printf("%p\n", theMap)

		return theMap
	}

	fmt.Printf("%p\n", getMapFunc())
	// 结论：同切片，作为方法参数传递，传递的是地址
}

// 复杂的键：测试结构体实例作为 map 的键
func testStructKey() {
	type box struct {
		height int
		color  string
	}

	boxes := []box{{1, "红"}, {2, "蓝"}, {3, "绿"}}
	length := len(boxes)

	boxMap := make(map[box]bool, length)
	for index := 0; index < length; index++ {
		boxMap[boxes[index]] = true
	}

	println(boxMap[box{1, "红"}])
}
