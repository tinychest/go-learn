package list_node

type ListNode struct {
	Val  int
	Next *ListNode
}

func FromArr(nums ...int) *ListNode {
	var head = new(ListNode)
	p := head
	for _, v := range nums {
		p.Next = &ListNode{Val: v}
		p = p.Next
	}
	return head.Next
}
