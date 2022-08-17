package _55min_stack

import "math"

// Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
//
// Implement the MinStack class:
//
//    MinStack() initializes the stack object.
//    void push(int val) pushes the element val onto the stack.
//    void pop() removes the element on the top of the stack.
//    int top() gets the top element of the stack.
//    int getMin() retrieves the minimum element in the stack.
//
// You must implement a solution with O(1) time complexity for each function.
//
//
// Constraints:
//
// -231 <= val <= 231 - 1
// Methods pop, top and getMin operations will always be called on non-empty stacks.
// At most 3 * 104 calls will be made to push, pop, top, and getMin.

// 下面的解法自然没有问题，但是性能不咋地
// [参考]
// > 这里不贴具体的实现代码
// 维护最小值实际上有技巧，肯定是要维护一个用于获取最小值的数组或者链表
// 这里边并不容易想到，并不需求将所有元素添加进来；假如，1 个数不是最小值，那么他将不再维护最小值关系的数组中，
// 那么从主栈中删除这个元素，实际上最小关系的数组中并不需要做任何调整

type Node struct {
	Val  int
	Next *Node
}

type MinStack struct {
	Head *Node
	Min  int
}

func Constructor() MinStack {
	return MinStack{Min: math.MaxInt32}
}

func (s *MinStack) Push(val int) {
	if s.Min > val {
		s.Min = val
	}

	n := &Node{Val: val, Next: s.Head}
	s.Head = n
}

func (s *MinStack) Pop() {
	if s.Head == nil {
		return
	}
	hv := s.Head.Val
	s.Head = s.Head.Next

	// 维护最小值
	if s.Min == hv {
		s.Min = math.MaxInt32
		for p := s.Head; p != nil; p = p.Next {
			if s.Min > p.Val {
				s.Min = p.Val
			}
		}
	}
}

func (s *MinStack) Top() int {
	return s.Head.Val
}

func (s *MinStack) GetMin() int {
	return s.Min
}
