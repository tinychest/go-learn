package _38copy_list_with_random_pointer

// 题目就是深拷贝链表节点数据结构如下的链表
// （偷懒，test 不写了，若再遇到该数据结构的题，再考虑抽取出一套相关方法）

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 假如没有 random，这个复制很简单；所以重点是如何复制 random 关系，按照理所当然的思路整了一遍
//
// [参考]
// 可以将序号抽象成 id 去理解，只需要一个辅助 map 就够了，创建 node 的过程直接考虑 Random 属性从而创建节点和 map 中的记录
func copyRandomList(head *Node) *Node {
	// 遍历链表：构建 节点-下标 map
	nodeIdxM := make(map[*Node]int)

	idx := 0
	for p := head; p != nil; p = p.Next {
		nodeIdxM[p] = idx
		idx++
	}
	nodeIdxM[nil] = -1

	// 遍历链表：构建 节点下标-指向的随机下标 map、构建 结果链表 + 下标-节点 map
	idxToRIdxM := make(map[int]int)
	res := new(Node)
	resP := res
	resIdxNodeM := make(map[int]*Node)

	idx = 0
	for p := head; p != nil; p = p.Next {
		idxToRIdxM[idx] = nodeIdxM[p.Random]

		node := &Node{Val: p.Val}
		resIdxNodeM[idx] = node
		resP.Next = node
		resP = node

		idx++
	}

	// 遍历结果链表维护随机指向关系
	idx = 0
	for p := res.Next; p != nil; p = p.Next {
		if i := idxToRIdxM[idx]; i != -1 {
			p.Random = resIdxNodeM[i]
		}
		idx++
	}

	return res.Next
}
