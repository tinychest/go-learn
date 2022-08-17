package tree_node

type TreeNodeGeneric[T int | string] struct {
	Val   T
	Left  *TreeNodeGeneric[T]
	Right *TreeNodeGeneric[T]
}

// GenericFromArr 根据层级遍历的结果构建出树
// 调用者保证，参数中没有的节点用 nil 表示，否则仅靠一个角度的遍历是无法反推树的结构的
// 调用者保证，参数中的数据是合理的
func GenericFromArr[T int | string](arr ...*T) *TreeNodeGeneric[T] {
	if len(arr) == 0 {
		return nil
	}

	pos := 0
	nodes := make([]*TreeNodeGeneric[T], 0, len(arr))
	nodes = append(nodes, &TreeNodeGeneric[T]{Val: *arr[0]})

	for i := 0; i < len(arr)-1; i += 2 {
		var left, right *TreeNodeGeneric[T]
		if arr[i+1] != nil {
			left = &TreeNodeGeneric[T]{Val: *arr[i+1]}
		}
		if arr[i+2] != nil {
			right = &TreeNodeGeneric[T]{Val: *arr[i+2]}
		}

		nodes[pos].Left = left
		nodes[pos].Right = right

		nodes = append(nodes, left, right)
		pos++
	}
	return nodes[0]
}

func (tree *TreeNodeGeneric[T]) ToArr() []*T {
	if tree == nil {
		return nil
	}

	nodes := make([]*TreeNodeGeneric[T], 0)
	nodes = append(nodes, tree)

	for pos := 0; pos < len(nodes); pos++ {
		if nodes[pos] == nil {
			continue
		}
		if nodes[pos].Left != nil || nodes[pos].Right != nil {
			nodes = append(nodes, nodes[pos].Left, nodes[pos].Right)
		}
	}

	res := make([]*T, 0)
	for _, v := range nodes {
		var t *T
		if v != nil {
			t = &v.Val
		}
		res = append(res, t)
	}
	return res
}
