package tree_node

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// FromArr 根据层级遍历的结果构建出树
// 调用者保证，参数中没有的节点用 nil 表示，否则仅靠一个角度的遍历是无法反推树的结构的
// 调用者保证，参数中的数据是合理的
func FromArr(arr ...*int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	pos := 0
	nodes := make([]*TreeNode, 0, len(arr))
	nodes = append(nodes, &TreeNode{Val: *arr[0]})

	for i := 0; i < len(arr)-1; i += 2 {
		var left, right *TreeNode
		if arr[i+1] != nil {
			left = &TreeNode{Val: *arr[i+1]}
		}
		if arr[i+2] != nil {
			right = &TreeNode{Val: *arr[i+2]}
		}

		nodes[pos].Left = left
		nodes[pos].Right = right

		nodes = append(nodes, left, right)
		pos++
	}
	return nodes[0]
}

func (tree *TreeNode) ToArr() []*int {
	if tree == nil {
		return nil
	}

	nodes := make([]*TreeNode, 0)
	nodes = append(nodes, tree)

	for pos := 0; pos < len(nodes); pos++ {
		if nodes[pos] == nil {
			continue
		}
		if nodes[pos].Left != nil || nodes[pos].Right != nil {
			nodes = append(nodes, nodes[pos].Left, nodes[pos].Right)
		}
	}

	res := make([]*int, 0)
	for _, v := range nodes {
		var t *int
		if v != nil {
			t = &v.Val
		}
		res = append(res, t)
	}
	return res
}
