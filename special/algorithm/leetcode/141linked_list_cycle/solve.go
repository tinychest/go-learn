package _41linked_list_cycle

import . "go-learn/special/algorithm/leetcode/0common/list_node"

// Given head, the head of a linked list, determine if the linked list has a cycle in it.
//
// There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.
//
// Return true if there is a cycle in the linked list. Otherwise, return false.
//
// Constraints:
//
//    The number of the nodes in the list is in the range [0, 104].
//    -105 <= Node.val <= 105
//    pos is -1 or a valid index in the linked-list.
//
//
//
// Follow up: Can you solve it using O(1) (i.e. constant) memory?

// 这个其实并不符合题目空间复杂度的要求，但是也提交通过了
// 这个也短是比较经典的题目了，在一些语言底层、垃圾回收中的对象互相引用中涉及到该问题
//
// 参考，得到 快慢指针法，很巧妙，其原理，为什么快慢指针一定能相遇，可以这样去理解：
// - 相差一格
// 快 []
// [] 慢
// 下一次会相遇
// - 相差两格
// 快 [] []
// [] [] 慢
// 下一次会变成第一种情况
// - 相差三格
// 快 [] [] []
// [] [] [] 慢
// 下一次会变成第二种情况
// 以此类推，会发现，终将相遇，在环形中，前面就是后面，后边就是前面
func hasCycle(head *ListNode) bool {
	nodeAddrM := make(map[*ListNode]bool, 0)

	for head != nil {
		if nodeAddrM[head] {
			return true
		}
		nodeAddrM[head] = true
		head = head.Next
	}
	return false
}
