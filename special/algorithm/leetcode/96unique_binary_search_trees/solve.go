package _6unique_binary_search_trees

// Given an integer n, return the number of structurally unique BST's (binary search trees) which has exactly n nodes of unique values from 1 to n.
//
// Constraints:
//
//    1 <= n <= 19
//
// 给定指定的数量的节点（节点的值就是从），能够构成多少颗不同的二叉搜索树
// 分析：感觉也是动态规划、可以用递归的意思去做，如 n 个节点得到的可能数和 n - 1 个节点得到的可能数的关系
//
// 对上了这题的做法，一次搞定

func numTrees(n int) int {
	// [1 个节点]
	// 1 作为根 → 1
	// = 1
	//
	// [2 个节点]
	// 1 作为根 → 1
	// 2 作为根 → 1
	// = 2
	//
	// [3 个节点]
	// 1 作为根 → 2 先、3 先 → 2
	// 2 作为根 → 1
	// 3 作为根 → [2 个节点]
	// = 5
	//
	// [4 个节点]
	// 1 作为根 → 3-1 + 3-2 + 3-3 → [3 个节点]
	// 2 作为根 → [1 个节点] * [2 个节点]
	// 3 作为根 → [2 个节点] * [1 个节点]
	// 4 作为根 → [3 个节点]
	// = 5 + 5 + 2 + 2 = 14

	arr := make([]int, n+1)
	arr[0] = 1
	arr[1] = 1
	for i := 2; i <= n; i++ {
		sum := 0
		for j := 0; j < i; j++ {
			sum += arr[j] * arr[i-j-1]
		}
		arr[i] = sum
	}
	return arr[n]
}
