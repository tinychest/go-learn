package _24binary_tree_maximum_path_sum

import . "go-learn/special/algorithm/leetcode/0common/tree_node"

// A path in a binary tree is a sequence of nodes where each pair of adjacent nodes in the sequence has an edge connecting them.
// A node can only appear in the sequence at most once.
// Note that the path does not need to pass through the root.
//
// The path sum of a path is the sum of the node's values in the path.
//
// Given the root of a binary tree, return the maximum path sum of any non-empty path.
//
//
//
// Constraints:
//
//    The number of nodes in the tree is in the range [1, 3 * 104].
//    -1000 <= Node.val <= 1000

// 题目的大概意思是说，在二叉树上找到一条路径，使得路径上节点的值之和最大
//
// [分析]
// 二叉树上任意两个节点之间都有一条路径（好像并不止一条），那么就可以得到一个路径上节点值的和
// 树的题，不把一些相关信息总结到父节点或者路径上，都说不过去
// 于是，便画图思考，应当按照这样的思路解题，我们应当算出每条边上能到达的最远距离，这样遍历整棵树就可以得出答案了
// 我们把上面说的这个值用一个 map 去存储
func maxPathSum(root *TreeNode) int {
	// return step01(root, make(map[*TreeNode]int))
	// return wrong02(root)
	// res, _ := step03(root)
	// return res

	res := -1001
	step04(root, &res)
	return res
}

// 提交效果不佳
func step01(root *TreeNode, pathM map[*TreeNode]int) int {
	if root == nil {
		return -1001
	}
	leftMax := step01(root.Left, pathM)
	rightMax := step01(root.Right, pathM)

	// 经过指定节点能得到的最大路径值延伸
	pathM[root] = max(root.Val+max(pathM[root.Left], pathM[root.Right]), root.Val)

	// 返回经过当前节点，子树最大节点值和路径的具体值和
	// - 第一部分：根节点 + 左子树(可选) + 右子树(可选)
	// - 第二部分：左子树
	// - 第三部分：右子树
	return max(root.Val+max(pathM[root.Left], 0)+max(pathM[root.Right], 0), leftMax, rightMax)
}

// 感觉 map 实可以去掉的
// 不对，下面这个逻辑是应用于根节点的，但不适用子节点，因为 要求用到节点的最大值 和 当前节点给出的最大值 不是一个概念
func wrong02(root *TreeNode) int {
	if root == nil {
		return -1001
	}
	leftMax := wrong02(root.Left)
	rightMax := wrong02(root.Right)

	// 返回经过当前节点，子树最大节点值和路径的具体值和
	// - 第一部分：根节点 + 左子树(可选) + 右子树(可选)
	// - 第二部分：左子树
	// - 第三部分：右子树
	return max(root.Val+max(leftMax, 0)+max(rightMax, 0), leftMax, rightMax)
}

// 返回：题目要求的值, 必须使用到当前节点的值
// 提交效果还是不佳
func step03(root *TreeNode) (int, int) {
	if root == nil {
		return -1001, -1001
	}
	leftMax, useLeftMax := step03(root.Left)
	rightMax, useRightMax := step03(root.Right)

	// 返回部分1：使用当前节点的值 + 尽可能使用上左节点和右节点 | 看左子树的了（当前节点的值太小了） | 看右子树的了（当前节点的值太小了）
	// 返回部分2：因为是能够向上传递的，使用了当前节点值的最大值，因此左子树和右子树只能取其一
	return max(root.Val+max(useLeftMax, 0)+max(useRightMax, 0), leftMax, rightMax), root.Val + max(useLeftMax, useRightMax, 0)
}

// 参考：和标准答案很接近了，参考答案很精简，和自己的思路最大的不同就是，自己的函数没必要弄个第二个返回值，只要用一个变量记录一下过程中出现的最大值即可
// 还是不行，个人猜测是函数内联的太多了，这个题就到这里了
func step04(root *TreeNode, res *int) int {
	if root == nil {
		return -1001
	}
	useLeftMax := step04(root.Left, res)
	useRightMax := step04(root.Right, res)

	*res = max(*res, root.Val+max(useLeftMax, 0)+max(useRightMax, 0))
	return root.Val + max(useLeftMax, useRightMax, 0)
}

func max(arr ...int) int {
	res := arr[0]
	for i := 1; i < len(arr); i++ {
		if res < arr[i] {
			res = arr[i]
		}
	}
	return res
}
