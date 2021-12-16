package util

// TODO list 包下有提供默认实现

type Queue struct {
	Head *Node
	Tail *Node
	Size int
}

type Node struct {
	Data interface{}
	Next *Node
}

func NewQueue() *Queue {
	return new(Queue)
}

func (q *Queue) In(data interface{}) *Queue {
	node := &Node{
		Data: data,
	}

	if q.IsEmpty() {
		q.Head = node
		q.Tail = node
	} else {
		q.Tail.Next = node
		q.Tail = node
	}
	q.Size++
	return q
}

func (q *Queue) Out() (result interface{}, ok bool) {
	if q.IsEmpty() {
		return nil, false
	}
	result = q.Head.Data

	q.Head = q.Head.Next
	q.Size--

	if q.IsEmpty() {
		q.Tail = nil
	}

	return result, !q.IsEmpty()
}

func (q *Queue) IsEmpty() bool {
	return q.Size == 0
}
