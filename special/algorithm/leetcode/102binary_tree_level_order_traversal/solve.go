package _02binary_tree_level_order_traversal

import . "go-learn/special/algorithm/leetcode/0common/tree_node"

func NewQueue() *Queue {
	return &Queue{}
}

type Queue struct {
	head *Node
	tail *Node
}

type Node struct {
	node *TreeNode
	next *Node
}

func (q *Queue) IsEmpty() bool {
	return q.head == nil
}

// Pop 从队首弹出一个元素
func (q *Queue) Pop() *TreeNode {
	if q.head == nil {
		return nil
	}
	if q.head == q.tail {
		q.tail = nil
	}
	t := q.head
	q.head = q.head.next
	return t.node
}

// Head 获取队首的元素
func (q *Queue) Head() *TreeNode {
	if q.head == nil {
		return nil
	}
	return q.head.node
}

// Tail 获取队尾的元素
func (q *Queue) Tail() *TreeNode {
	if q.tail == nil {
		return nil
	}
	return q.tail.node
}

func (q *Queue) Push(e *TreeNode) {
	n := &Node{node: e}
	if q.head == nil {
		q.head = n
	} else {
		q.tail.next = n
	}
	q.tail = n
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	q := NewQueue()
	q.Push(root)
	c := make([][]int, 0)
	for !q.IsEmpty() {
		c = append(c, traverseLevel(q, q.Tail()))
	}
	return c
}

func traverseLevel(q *Queue, endNode *TreeNode) []int {
	res := make([]int, 0)

	for n := q.Pop(); n != nil; n = q.Pop() {
		if n.Left != nil {
			q.Push(n.Left)
		}
		if n.Right != nil {
			q.Push(n.Right)
		}
		res = append(res, n.Val)

		if n == endNode {
			return res
		}
	}
	return res
}
