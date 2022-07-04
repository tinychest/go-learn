package _01symmetric_tree

// Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
//
// Constraints:
//
//    The number of nodes in the tree is in the range [1, 1000].
//    -100 <= Node.val <= 100

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	// return step1(root)
	return refer(root)
}

// 做题时间超时，提交效果还不佳
// 时间：50%
// 空间：80%
func step1(root *TreeNode) bool {
	if root == nil {
		return true
	}

	const Unavailable = 101

	s1 := make([]*TreeNode, 0)
	s2 := make([]*TreeNode, 0)

	l, r := root, root
	for l != nil && r != nil || len(s1) != 0 && len(s2) != 0 {
		lv, rv := Unavailable, Unavailable
		// 获取 先左序遍历 节点的值
		if l != nil {
			lv = l.Val
			s1 = append(s1, l)
			l = l.Left
		} else {
			// 出栈（准备下一次遍历的节点）
			if len(s1) != 0 {
				l = s1[len(s1)-1]
				s1 = s1[:len(s1)-1]
				l = l.Right
			}
		}
		// 获取 先右序遍历 节点的值
		if r != nil {
			rv = r.Val
			s2 = append(s2, r)
			r = r.Right
		} else {
			// 出栈（准备下一次遍历的节点）
			if len(s2) != 0 {
				r = s2[len(s2)-1]
				s2 = s2[:len(s2)-1]
				r = r.Left
			}
		}

		if lv != rv {
			return false
		}
	}
	return true
}

// 向着递归方向思考解决该题的标准解如下
// Runtime: 3 ms, faster than 64.78% of Go online submissions for Symmetric Tree.
// Memory Usage: 2.9 MB, less than 18.21% of Go online submissions for Symmetric Tree.
// （效果没有想象中的好）
func refer(root *TreeNode) bool {
	return referHelper(root.Left, root.Right)
}

func referHelper(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && referHelper(left.Right, right.Left) && referHelper(left.Left, right.Right)
}
