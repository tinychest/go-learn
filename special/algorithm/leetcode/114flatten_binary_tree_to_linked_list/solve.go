package _14flatten_binary_tree_to_linked_list

import . "go-learn/special/algorithm/leetcode/0common/tree_node"

// Given the root of a binary tree, flatten the tree into a "linked list":
//
//    The "linked list" should use the same TreeNode class where the right child pointer points to the next node in the list and the left child pointer is always null.
//    The "linked list" should be in the same order as a pre-order traversal of the binary tree.
//
//
// Constraints:
//
//    The number of nodes in the tree is in the range [0, 2000].
//    -100 <= Node.val <= 100
//
//
// Follow up: Can you flatten the tree in-place (with O(1) extra space)?

// 其实按照要求和提示去做，题目并不难
// - 不用额外的空间，那就递归咯
// - 将先序进行连接，那就先写个先序遍历咯
// 希望实现目标效果，发现一来直接设置右节点不行，但是，左节点可以；所以，先左节点练成一线，随后调整为右节点
// [结果]
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Flatten Binary Tree to Linked List.
// Memory Usage: 3 MB, less than 25.86% of Go online submissions for Flatten Binary Tree to Linked List.
// [参考]
// 可以变形一下，递归方法中直接处理当前节点和两个子节点，而不是仅处理当前节点，两个子节点递归
// 还应该把最终的左树转成右树的操作放到递归过程里边
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	// 这里一定要定义一个额外的临时变量，否则，在递归方法中修改 p 底层对应的值，就是修改 root 的值，而 root 变量的值是不应该更改的
	tmp := root
	p := &tmp
	preOrder(root.Left, p)
	preOrder(root.Right, p)

	for t := root; t != nil; t = t.Right {
		t.Right = t.Left
		t.Left = nil
	}
}

func preOrder(root *TreeNode, last **TreeNode) {
	if root == nil {
		return
	}
	(*last).Left = root
	*last = root

	preOrder(root.Left, last)
	preOrder(root.Right, last)
}

// 感觉这个叫解决问题的核心规律，但是自己这个也不错
func refer(root *TreeNode) {
	rightBottom := func(node *TreeNode) *TreeNode {
		for node.Right != nil {
			node = node.Right
		}
		return node
	}

	if root == nil {
		return
	}

	if root.Left != nil {
		rbNode := rightBottom(root.Left)
		rbNode.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}

	flatten(root.Right)
}
