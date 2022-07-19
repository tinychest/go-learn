package _5reverse_nodes_in_k_group

import . "go-learn/special/algorithm/leetcode/0common/list_node"

// Given the head of a linked list, reverse the nodes of the list k at a time, and return the modified list.
//
// k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes, in the end, should remain as it is.
//
// You may not alter the values in the list's nodes, only nodes themselves may be changed.
//
//
// Constraints:
//
//    The number of nodes in the list is n.
//    1 <= k <= n <= 5000
//    0 <= Node.val <= 1000
//
//
// Follow-up: Can you solve the problem in O(1) extra memory space?

func reverseKGroup(head *ListNode, k int) *ListNode {
	// 先用递归的思路试一下（实际找工作上机，做题时间也是很重要的指标）
	return step1(head, k)
}

// Runtime: 8 ms, faster than 41.60% of Go online submissions for Reverse Nodes in k-Group.
// Memory Usage: 3.8 MB, less than 12.82% of Go online submissions for Reverse Nodes in k-Group.
// 不是很理想
func step1(head *ListNode, k int) *ListNode {
	// 先用递归的思路试一下（实际找工作上机，做题时间也是很重要的指标）
	container := make([]*ListNode, 0, k)
	return step1Recursive(head, k, container)
}

func step1Recursive(head *ListNode, k int, container []*ListNode) *ListNode {
	container = container[:0]
	p := head

	for i := 0; i < k; i++ {
		if p == nil {
			return head
		}
		container = append(container, p)
		p = p.Next
	}

	res := container[k-1]
	behindHead := container[k-1].Next
	for i := k - 1; i >= 1; i-- {
		container[i].Next = container[i-1]
	}
	behind := step1Recursive(behindHead, k, container)
	head.Next = behind
	return res
}

// 想不借助辅助长度为 k 的数组来实现这个，则需要另外一个递归
