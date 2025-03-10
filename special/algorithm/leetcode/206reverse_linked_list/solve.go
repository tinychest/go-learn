package _06reverse_linked_list

import (
	. "go-learn/special/algorithm/leetcode/0common/list_node"
)

// Given the head of a singly linked list, reverse the list, and return the reversed list.
//
//
// Example 1:
//
// Input: head = [1,2,3,4,5]
// Output: [5,4,3,2,1]
//
// Example 2:
//
// Input: head = [1,2]
// Output: [2,1]
//
// Example 3:
//
// Input: head = []
// Output: []
//
//
//
// Constraints:
//
//    The number of nodes in the list is the range [0, 5000].
//    -5000 <= Node.val <= 5000
//
//
//
// Follow up: A linked list can be reversed either iteratively or recursively. Could you implement both?

func reverseList(head *ListNode) *ListNode {
	// return solve(head)
	// return pref01(head)
	return pref02(head)
}

// 优化方向，一个方法的单递归就可以解决
func pref01(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	newHead := pref01(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// 其实不需要递归，这里边最大的思想屏障是，你不需要遍历到链表的根，再开始逆转
// pref 不需要从一个实际的节点开始，为 nil 最好
func pref02(head *ListNode) *ListNode {
	var pref, next *ListNode
	cur := head

	for cur != nil {
		// 记录
		next = cur.Next
		// 变更
		cur.Next = pref
		// 位移
		pref = cur
		cur = next
	}

	return pref
}

// 先来个递归（写的并不好，认为只有遍历到结尾才知道逆转后的头，应该从头开始逆转，逆转到最后就是头更好一些）
func solve(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	_, res := solve1core(head)
	head.Next = nil
	return res
}

func solve1core(head *ListNode) (cur, end *ListNode) {
	var t *ListNode

	if head.Next != nil {
		t, end = solve1core(head.Next)
		t.Next = head
	} else {
		end = head
	}

	return head, end
}
