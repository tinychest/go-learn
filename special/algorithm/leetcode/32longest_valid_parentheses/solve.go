package _2longest_valid_parentheses

// Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.
//
//
// Constraints:
//
//    0 <= s.length <= 3 * 104
//    s[i] is '(', or ')'.

// 出发：基于 common 中高效的单类型括号是否匹配算法，进行拓展调整
// 核心：这道题借助真实的栈去做应该达到最佳的时间复杂度，但是自己希望发现其中规律，并且借助额外的空间去做
//
// 正向遍历：能够准确知道什么括号是不匹配的，但是不希望借助栈来实现这个判断最大匹配的算法，从而将后续的串作为起始串重新执行算法
// 在过程中发现了一个不对的逆向规律，因此引出 2，但实际发现 2 是不对的，正向和逆向是一样的
// 重要的右括号，如果进行左遍历，左括号不就成为了右括号（起始在这里已经想到了，尝试去进行正向和逆向的遍历）
//
// 最后想到下面这种做法，真的做了很多尝试
//
// 结论：最终实现了不借助栈的写法
// - 参考答案（submission 统计信息）中进行 正向遍历 和 逆向遍历 可以得到一个结果，最节省空间
// - 自己发现的规律可能不透彻
// 详见下面的 refer 思想，真的很简单
func longestValidParentheses(s string) int {
	// res：最终最大连续匹配数量
	// left：当前段的左括号数
	// right：当前段的右括号数
	res, left, right := 0, 0, 0

	// 最后一个导致不匹配的右括号的位置
	lastWrongRightIdx := -1

	// 匹配段最大连续数统计
	n := len(s)

	for i := 0; i < n; i++ {

		if s[i] == '(' {
			left++
		} else {
			right++
		}

		// 加上这个右括号，串就肯定不匹配了；此时最大的括号匹配组数量应该是左括号和右括号的数量
		if right > left {
			if res < left {
				res = left
			}
			lastWrongRightIdx = i
			left = 0
			right = 0
		}
	}
	res *= 2

	// 对末段 [因缺少右括号导致不匹配] 或者 [完全匹配的括号段]（有可能是整个段）进行最大连续括号匹配数计算
	// 发现的易用的规律：使用的思想是，从末端开始假如有一个下标指针，额外存储着一个整型值，这个指针向前走
	// 如果是右括号就自增 1，左括号就自减 1，值不能为负数，为负数就从 0 重新开始
	// 这样走过的最长的路就是最长的括号匹配数

	// 字段含义变更，left 表示剩下（值），right 表示最大连续匹配数
	left, right = 0, 0
	for i := n - 1; i > lastWrongRightIdx; i-- {
		if s[i] == ')' {
			left++
		} else {
			left--
		}

		right++

		if left < 0 {
			left = 0
			right = 0
		}
		if res < right {
			res = right
		}
	}

	return res
}

// 空间使用最少的参考答案
func refer(s string) int {
	var longest int
	if len(s) < 1 {
		return 0
	}
	n := len(s)

	left, right := 0, 0
	// Scan from left to right
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left < right {
			left, right = 0, 0
		} else if left == right {
			count := left + right
			if count > longest {
				longest = count
			}
		}
	}

	left, right = 0, 0
	// Scan from right to left
	for i := n - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left > right {
			left, right = 0, 0
		} else if left == right {
			count := left + right
			if count > longest {
				longest = count
			}
		}
	}

	return longest
}
