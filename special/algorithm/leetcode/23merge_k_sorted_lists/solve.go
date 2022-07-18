package _3merge_k_sorted_lists

import . "go-learn/special/algorithm/leetcode/0common/list_node"

// You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.
//
// Merge all the linked-lists into one sorted linked-list and return it.
//
//
// Constraints:
//
//    k == lists.length
//    0 <= k <= 104
//    0 <= lists[i].length <= 500
//    -104 <= lists[i][j] <= 104
//    lists[i] is sorted in ascending order.
//    The sum of lists[i].length will not exceed 104.

func mergeKLists(lists []*ListNode) *ListNode {
	// return step1(lists)
	return step2(lists)
}

// Runtime: 325 ms, faster than 16.78% of Go online submissions for Merge k Sorted Lists.
// Memory Usage: 5.3 MB, less than 74.62% of Go online submissions for Merge k Sorted Lists.
// 全部一起上的性能很差，空间也不咋地
func step1(lists []*ListNode) *ListNode {
	fakeHead := new(ListNode)
	fakePos := fakeHead

	var turnMinIdx int
	var turnMinVal int
	var turnNotZero int

	for {
		turnMinIdx = -1
		turnMinVal = 105
		turnNotZero = 0

		for idx, list := range lists {
			if list == nil {
				continue
			}
			turnNotZero++
			if turnMinVal > list.Val {
				turnMinIdx = idx
				turnMinVal = list.Val
			}
		}

		if turnNotZero < 2 {
			break
		}
		fakePos.Next = lists[turnMinIdx]
		fakePos = lists[turnMinIdx]
		lists[turnMinIdx] = lists[turnMinIdx].Next
	}
	if turnMinIdx != -1 {
		fakePos.Next = lists[turnMinIdx]
	}
	return fakeHead.Next
}

// 全部一起上性能不行，那就是分而治之，参考 21
// Runtime: 125 ms, faster than 32.46% of Go online submissions for Merge k Sorted Lists.
// Memory Usage: 5.3 MB, less than 74.62% of Go online submissions for Merge k Sorted Lists.
// 性能提升了好几倍
func step2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	res := lists[0]
	for i := 1; i < len(lists); i++ {
		res = mergeTwoLists(res, lists[i])
	}
	return res
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p := &ListNode{}
	h := p
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
		p = p.Next
	}
	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}
	return h.Next
}
