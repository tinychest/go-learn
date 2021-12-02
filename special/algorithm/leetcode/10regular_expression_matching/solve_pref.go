package _0regular_expression_matching

// 好好理解一下这个算法
func isMatchPre(text, pattern string) bool {
	var l1, l2 = len(text), len(pattern)

	var dp = make([][]bool, l1+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, l2+1)
	}
	dp[l1][l2] = true

	for i := l1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			firstMatch := i < l1 && (pattern[j] == text[i] || pattern[j] == '.')

			if j+1 < l2 && pattern[j+1] == '*' {
				dp[i][j] = dp[i][j+2] || firstMatch && dp[i+1][j]
			} else {
				dp[i][j] = firstMatch && dp[i+1][j+1]
			}

		}
	}
	return dp[0][0]
}
