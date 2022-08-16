package _48sort_list

import (
	. "go-learn/special/algorithm/leetcode/0common/list_node"
	"sort"
)

// Given the head of a linked list, return the list after sorting it in ascending order.
//
//
// Constraints:
//
//    The number of nodes in the list is in the range [0, 5 * 104].
//    -105 <= Node.val <= 105

func sortList(head *ListNode) *ListNode {
	// return fail01(head)
	// return step02(head)
	return step03(head)
}

// 通过翻阅相关排序算法，得到最适合的排序算法，应该是 归并排序，但是归并排序的标准实现包含+通过下标取元素，所以不好实现
// [结果] Time Limit Exceeded
func fail01(head *ListNode) *ListNode {
	var l int
	for p := head; p != nil; p = p.Next {
		l++
	}

	for i := 0; i < l-1; i++ {
		p := head
		for j := 0; j < l-i-1; j++ {
			if p.Val > p.Next.Val {
				p.Val, p.Next.Val = p.Next.Val, p.Val
			}
			p = p.Next
		}
	}
	return head
}

// 不借助额外的空间，整不来
func step02(head *ListNode) *ListNode {
	var l int
	for p := head; p != nil; p = p.Next {
		l++
	}

	arr := make([]int, 0, l)
	for p := head; p != nil; p = p.Next {
		arr = append(arr, p.Val)
	}

	sort.Ints(arr)

	p := head
	for i := 0; i < l; i++ {
		p.Val = arr[i]

		p = p.Next
	}
	return head
}

// 参考了一下提交，果然执行时间最短的，就是利用 merge sort
// 想知道下标，就遍历至目标下标就行
// 就这样的做法，多次提交的结果差异过大，runtime 的参考意义不大
func step03(head *ListNode) *ListNode {
	var l int
	for p := head; p != nil; p = p.Next {
		l++
	}
	return mergeSort(head, l)
}

func mergeSort(head *ListNode, l int) *ListNode {
	if l < 2 {
		return head
	}

	mid := head
	for i := 0; i < l/2-1; i++ {
		mid = mid.Next
	}
	// 打断
	t := mid.Next
	mid.Next = nil
	mid = t

	head = mergeSort(head, l/2)
	mid = mergeSort(mid, l-l/2)
	return merge(head, mid)
}

func merge(a, b *ListNode) *ListNode {
	head := new(ListNode)
	t := head
	for a != nil && b != nil {
		if a.Val < b.Val {
			t.Next = a
			a = a.Next
		} else {
			t.Next = b
			b = b.Next
		}
		t = t.Next
	}
	if a != nil {
		t.Next = a
	}
	if b != nil {
		t.Next = b
	}
	return head.Next
}
