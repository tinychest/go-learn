package _9remove_nth_node_from_end_of_list

// Given the head of a linked list, remove the nth node from the end of the list and return its head.
//
// Constraints:
//
//    The number of nodes in the list is sz.
//    1 <= sz <= 30
//    0 <= Node.val <= 100
//    1 <= n <= sz
//
//
//
// Follow up: Could you do this in one pass?

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node := &ListNode{}
	node.Next = head
	head = node

	xxx(node, n)
	return head.Next
}

// 时间复杂度：100%
func xxx(head *ListNode, n int) int {
	var res int
	if head.Next == nil {
		res = 1
	} else {
		res = xxx(head.Next, n) + 1
	}

	if res == n+1 {
		head.Next = head.Next.Next
	}
	return res
}
