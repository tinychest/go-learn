package _42linked_list_cycle_II

import . "go-learn/special/algorithm/leetcode/0common/list_node"

// Given the head of a linked list, return the node where the cycle begins. If there is no cycle, return null.
//
// There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to (0-indexed). It is -1 if there is no cycle. Note that pos is not passed as a parameter.
//
// Do not modify the linked list.
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

// 经过第一题，相信也知道快慢指针法是解决这种题最好的方法
// 这里，进一步提出要求，希望直到发生环的位置，也就是遍历过程中第一次遍历到的重复节点
// 唯一可以利用的点就是，快慢指针法相遇的节点有什么讲究？
// 慢 走到一半时，快 已经到五环链表的结尾了；当 慢 走到关键环节点时，无论，此时快节点在环的哪个位置，总是能在慢节点走到无环终点前或者恰好到达无环终点时，相遇
// 简单画图分析一下：起点--------环点?-----相遇点------无环终点，把 3 段距离分别视作 a b c
// 那么根据快走的距离是慢走的距离的两倍，则有：2(a + b) = a + (n + 1)b + nc
// 变换一下：a + b = n(b + c)
// 假如我们让两个慢指针，分别从起点 和 相遇点同时出发，当第一个慢指针一个格子一个格子走到相遇点时，也就是第一个指针走了 a + b，第二个指针走了 n(b + c) 的距离，
// 也就是跑了 n 个圈回到了原点，也就是相遇点，也就是说两个点正好在相遇点会遇到，emmm，还没反应过来么，两个指针速度一样，emm，还没反应过来么，他们第一次的相遇点正好就是环点
// 解法也就出来
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	// 求相遇点
	oneP, twoP := head, head
	for twoP != nil && twoP.Next != nil {
		oneP = oneP.Next
		twoP = twoP.Next.Next
		if oneP == twoP {
			break
		}
	}
	// 无环
	if oneP != twoP {
		return nil
	}

	// 求环点
	oneP = head
	for oneP != twoP {
		oneP = oneP.Next
		twoP = twoP.Next
	}
	return oneP
}
