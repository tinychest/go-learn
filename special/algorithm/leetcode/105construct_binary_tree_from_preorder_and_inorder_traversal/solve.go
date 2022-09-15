package _05construct_binary_tree_from_preorder_and_inorder_traversal

import . "go-learn/special/algorithm/leetcode/0common/tree_node"

// Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.
//
//
//
// Constraints:
//
//    1 <= preorder.length <= 3000
//    inorder.length == preorder.length
//    -3000 <= preorder[i], inorder[i] <= 3000
//    preorder and inorder consist of unique values.
//    Each value of inorder also appears in preorder.
//    preorder is guaranteed to be the preorder traversal of the tree.
//    inorder is guaranteed to be the inorder traversal of the tree.

// 用两个特定顺序的遍历序列来构建二叉树，关键在于交替利用两个树中确定位置的节点，如先序遍历序列的第一个元素就是树根
// 这个题在学数据结构的时候，自己就编码写过，但是细想，序列中的值必须是独一无二的，否则，无法实现
// 中序遍历序列特点：任意节点的左孩子在遍历序列中也分布在左边，右孩子同理
// 先序和中序协同规律：对于树上任意一颗子树的遍历序列，其末端节点，两个顺序遍历的结果，它都是在最后
func buildTree(preorder []int, inorder []int) *TreeNode {
	return step01(preorder, inorder)
	// return step02(preorder, inorder)
}

// 提交效果不佳
func step01(preorder, inorder []int) *TreeNode {
	root := &TreeNode{Val: preorder[0]}
	var inorderRootIdx int
	for i, v := range inorder {
		if v == root.Val {
			inorderRootIdx = i
			break
		}
	}

	if inorderRootIdx != 0 {
		root.Left = step01(preorder[1:inorderRootIdx+1], inorder[0:inorderRootIdx])
	}
	if inorderRootIdx != len(inorder)-1 {
		root.Right = step01(preorder[inorderRootIdx+1:], inorder[inorderRootIdx+1:])
	}
	return root
}

// 使用了切片裁剪的特性无法或者很难复用 inorder 元素下标的 map
// 额，提交效果还是一样，这里的做法和参考答案的核心思想一样了
func step02(preorder, inorder []int) *TreeNode {
	inorderIdxM := make(map[int]int, len(inorder))
	for i, v := range inorder {
		inorderIdxM[v] = i
	}
	return help(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1, inorderIdxM)
}

func help(preorder []int, ps, pe int, inorder []int, is, ie int, inorderIdxM map[int]int) *TreeNode {
	rootVal := preorder[ps]
	root := &TreeNode{Val: rootVal}
	// 根节点在中序遍历序列中的下标
	inorderRootIdx := inorderIdxM[rootVal]
	// 左子树长度
	leftLen := inorderRootIdx - is

	if leftLen != 0 {
		root.Left = help(
			preorder, ps+1, ps+leftLen,
			inorder, is, inorderRootIdx-1,
			inorderIdxM,
		)
	}
	if inorderRootIdx != ie {
		root.Right = help(
			preorder, ps+leftLen+1, pe,
			inorder, inorderRootIdx+1, ie,
			inorderIdxM,
		)
	}
	return root
}
