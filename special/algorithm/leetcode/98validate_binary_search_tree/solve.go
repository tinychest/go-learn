package _8validate_binary_search_tree

// 简单回顾 数据结构
// 二叉树（BTree）、完全二叉树、满二叉树
// 二叉搜索树（BST）
// 平衡二叉树（AVL Tree）
// 红黑树

// Given the root of a binary tree, determine if it is a valid binary search tree (BST).
//
// A valid BST is defined as follows:
//
//    The left subtree of a node contains only nodes with keys less than the node's key.
//    The right subtree of a node contains only nodes with keys greater than the node's key.
//    Both the left and right subtrees must also be binary search trees.
//
//
// Constraints:
//
//    The number of nodes in the tree is in the range [1, 104].
//    -2^31 <= Node.val <= 2^31 - 1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	// return step1(nil, nil, root)

	const MaxInt = int(^uint(0) >> 1)
	const MinInt = -MaxInt - 1
	return step2(MaxInt, MinInt, root)
}

// [结果]
// Runtime: 8 ms, faster than 63.64% of Go online submissions for Validate Binary Search Tree.
// Memory Usage: 5.2 MB, less than 33.36% of Go online submissions for Validate Binary Search Tree.
// [小结]
// 错了很多次，其实真正关键的是从根节点出发
// - 判定左孩子是否满足题意时，应当更新 [左向树的最小值] 作为参数传给该节点
// - 判定右孩子是否满足题意时，应当更新 [右向树的最大值] 作为参数传给该节点
// 很容易错，需要警惕
func step1(lval, rVal *int, root *TreeNode) bool {
	if root == nil {
		return true
	}
	// 树中每一个节点都需要小于左向节点中最小的、大于右向节点中最大的
	if lval != nil && root.Val >= *lval {
		return false
	}
	if rVal != nil && root.Val <= *rVal {
		return false
	}
	return step1(minP(lval, root.Val), rVal, root.Left) && step1(lval, maxP(rVal, root.Val), root.Right)
}

func minP(a *int, b int) *int {
	if a == nil {
		return &b
	}
	if *a < b {
		return a
	}
	return &b
}

func maxP(a *int, b int) *int {
	if a == nil {
		return &b
	}
	if *a > b {
		return a
	}
	return &b
}

// 应用参考答案中的一些点
// Runtime: 11 ms, faster than 34.69% of Go online submissions for Validate Binary Search Tree.
// Memory Usage: 5.2 MB, less than 33.36% of Go online submissions for Validate Binary Search Tree.
// 变搓了...
func step2(lval, rVal int, root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Val >= lval {
		return false
	}
	if root.Val <= rVal {
		return false
	}
	return step2(min(lval, root.Val), rVal, root.Left) && step2(lval, max(rVal, root.Val), root.Right)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
