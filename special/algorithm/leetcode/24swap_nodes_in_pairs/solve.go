package _4swap_nodes_in_pairs

import (
	. "go-learn/special/algorithm/leetcode/0common/list_node"
)

// Given a linked list, swap every two adjacent nodes and return its head.
// You must solve the problem without modifying the values in the list's nodes
// (i.e., only nodes themselves may be changed.)
// 提示：不能通过修改节点数值的方式来实现
//
//
// Constraints:
//
//    The number of nodes in the list is in the range [0, 100].
//    0 <= Node.val <= 100

func swapPairs(listNode *ListNode) *ListNode {
	// return solve(listNode)
	return recursive(listNode)
}

// 解法没刺挑，但是限定使用递归的方式去解想必会更符合出题人的目的
func solve(listNode *ListNode) *ListNode {
	head := new(ListNode)
	head.Next = listNode
	p := head

	var t, np1, np2 *ListNode

	for p.Next != nil {
		np1 = p.Next
		if np1.Next == nil {
			break
		}
		np2 = np1.Next

		p.Next = np2
		t = np2.Next
		np2.Next = np1
		np1.Next = t

		p = np1
	}
	return head.Next
}

// 递归方式解决
// 当前节点为空、下一个节点为空；修改第 1 个节点指向、修改第 2 个节点指向；返回第 2 节点
func recursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p := head.Next
	head.Next = recursive(p.Next)
	p.Next = head
	return p
}
