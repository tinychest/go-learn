package _add_two_numbers

// You are given two non-empty linked lists representing two non-negative integers.
// The digits are stored in reverse order, and each of their nodes contains a single digit.
// Add the two numbers and return the sum as a linked list.
//
// You may assume the two numbers do not contain any leading zero, except the number 0 itself.
//
//
// Constraints:
//
//    The number of nodes in each linked list is in the range [1, 100].
//    0 <= Node.val <= 9
//    It is guaranteed that the list represents a number that does not have leading zeros.

type ListNode struct {
	Val  int
	Next *ListNode
}

// 其实申请节点空间是比较好的做法，代码会简单一些，其他并无本质区别
// - 下面写的太复杂了，要是面试中用这种方式解决，大概率要 g；有一点可以降低思考负担，就是短的链表可以认为是相同长度，但是高位的数都是零

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	t := &ListNode{Next: l2}
	l2 = t
	t = &ListNode{Next: l1}
	l1 = t

	p := t

	jinwei := 0
	sum := 0
	for l1.Next != nil && l2.Next != nil {
		v1, v2 := 0, 0
		if l1.Next != nil {
			v1 = l1.Next.Val
		}
		if l2.Next != nil {
			v2 = l2.Next.Val
		}

		sum = v1 + v2 + jinwei
		jinwei = 0

		if sum >= 10 {
			jinwei = sum / 10
			sum = sum % 10
		}

		l1.Next.Val = sum

		l1 = l1.Next
		l2 = l2.Next

		p = p.Next
	}
	for l1.Next != nil {
		sum = jinwei + l1.Next.Val
		jinwei = 0
		if sum >= 10 {
			jinwei = sum / 10
			sum = sum % 10
		}
		l1.Next.Val = sum
		l1 = l1.Next

		p = p.Next
	}
	if l2.Next != nil {
		l1.Next = l2.Next
	}
	for l2.Next != nil {
		sum = jinwei + l2.Next.Val
		jinwei = 0
		if sum >= 10 {
			jinwei = sum / 10
			sum = sum % 10
		}
		l2.Next.Val = sum
		l2 = l2.Next

		p = p.Next
	}

	if jinwei > 0 {
		p.Next = &ListNode{Val: jinwei}
	}
	return t.Next
}
