package _2generate_parentheses

// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
//
//
// Constraints:
//
//    1 <= n <= 8

func generateParenthesis(n int) []string {
	// g，并没有发现规律和方法，感觉像是从来没有涉及过的题 → 查看参考答案
	// > 这种毫无头绪的题，尽早去查看答案是很重要的
	// 好像并不会有什么规律可循，所以自己卡死了
	// 突破点：排列组合，会在过程发现匹配问题

	// n 个 ( 和 n 个 ) 的排列组合问题
	// - 最死的办法，随机组合，然后用一个验证括号是否匹配的方法去检查，匹配就加到结果集
	// - 想到一个可能好的解题模型，4 个左括号，拿 4 个右括号插入到左括号之间达成不同解（在题目的背景之下，只有一个要求，就是插入位置的右括号，其位置左边的左括号的数量不能少于右括号）
	//   实际发现这个思路的递归解法，就是讨论中参考的答案的极简写法

	resHelp := [][]int{make([]int, 1)}

	// 循环 n 次，每次得出坑位可以赋的值
	for i := 1; i <= n; i++ {
		// 基于 resHelp 中的前置数情况列举出所有新的情况
		newResHelp := make([][]int, 0)
		for _, res := range resHelp {
			if i == n {
				copyRes := copySlice(res)
				newResHelp = append(newResHelp, append(copyRes, n-copyRes[0]))
				continue
			}

			canTo := i - res[0]
			for j := 0; j <= canTo; j++ {
				copyRes := copySlice(res)
				copyRes[0] += j
				newResHelp = append(newResHelp, append(copyRes, j))
			}
		}
		resHelp = newResHelp
	}

	// 根据所有答案的组合情况得出最终的解
	finalRes := make([]string, 0)
	for _, vs := range resHelp {
		var str string
		for i := 1; i <= n; i++ {
			str += "(" + repeat(vs[i])
		}
		finalRes = append(finalRes, str)
	}

	return finalRes
}

func repeat(n int) string {
	switch n {
	case 0:
		return ""
	case 1:
		return ")"
	case 2:
		return "))"
	case 3:
		return ")))"
	case 4:
		return "))))"
	case 5:
		return ")))))"
	case 6:
		return "))))))"
	case 7:
		return ")))))))"
	case 8:
		return "))))))))"
	}
	panic("impossible! you are a liar!")
}

func copySlice(arr []int) []int {
	res := make([]int, len(arr), len(arr)+1)
	for i := 0; i < len(arr); i++ {
		res[i] = arr[i]
	}
	return res
}
