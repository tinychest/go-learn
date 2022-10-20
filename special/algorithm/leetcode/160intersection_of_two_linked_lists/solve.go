package _60intersection_of_two_linked_lists

import (
	. "go-learn/special/algorithm/leetcode/0common/list_node"
)

// Given the heads of two singly linked-lists headA and headB, return the node at which the two lists intersect. If the two linked lists have no intersection at all, return null.
//
// Constraints:
//
//    The number of nodes of listA is in the m.
//    The number of nodes of listB is in the n.
//    1 <= m, n <= 3 * 104
//    1 <= Node.val <= 105
//    0 <= skipA < m
//    0 <= skipB < n
//    intersectVal is 0 if listA and listB do not intersect.
//    intersectVal == listA[skipA] == listB[skipB] if listA and listB intersect.
//
//
// Follow up: Could you write a solution that runs in O(m + n) time and use only O(1) memory?

// 说人话，就是返回参数两个链表发生交集的节点
//
// 要想达成题目要求的时间复杂度
//
// - 思路1：遍历留下痕迹（使用 map 记录 或 在节点本身留下记号），A 先遍历，留下记号，再遍历 B，遍历过程中判断是否有记号就行了，但是借助了额外的空间，不行】
// - 思路2：如果存在相交节点的两个链表，将这两个链表进行尾部对齐，短链表开头对齐到长链表的节点位置要是能够知道，这个题就简单了
// 还有，有印象之前一道判断链表中是否有环的题目中有使用双指针法，有感觉这道题的做法应该和那个差不多
//
// > 为了表述清晰，假如 A 是长链表，B 是短链表
// 假如把两个链表头尾相连？两个指针 a、b 分别位于两个链表的头部，同时走，b 走完 B 时，a 走到了 B 的长度位置（这个位置，a 到 A 的终点还有 A - B 的距离，注意这个并不是我们要的）
//
// 自己的思路到这里就死了，实际上思路是对的，只是这样思考正好阻挡了正确答案，看了答案之后：
//
// 不应该以头尾相连的建模理解，事实上，题目也不让你改，链表结尾的 nil 是重要信息，应当以这样的逻辑 当 a 遍历到 nil，换到 B 的头，b 遍历到 nil，换到 A 的头
//
// 当 a 走完 A 时，B 因为走了 A 的距离，因此此时正好位于 A - B 的位置，这个正是我们需要的
// 先别急，继续看，a 走到了 B 的开头，b 走到了离 A 终点还差 B 距离的地方，继续走的话，a 会走到 B 的终点，b 会走到 A 的终点
// 因为是同时，细想，假如两个链表有相交的地方，那么两个指针一定会在同时遍历到 nil，相等前，就提前相等了
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pa, pb := headA, headB
	for pa != pb {
		if pa != nil {
			pa = pa.Next
		} else {
			pa = headB
		}

		if pb != nil {
			pb = pb.Next
		} else {
			pb = headA
		}
	}
	return pa
}
