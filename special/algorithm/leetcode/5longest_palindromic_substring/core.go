package _longest_palindromic_substring

// 题目简述：最长回文

// 就死办法去写，其中重试了不下 10 次，缝缝补补

// 可以改进代码逻辑（将相同的逻辑抽取出来）

// 改进算法1 O(n²)：求 目标串 与 目标回文串 求最长的公共子串（果然还是 KMP 有关系）
// 改进算法2 O(n)：特殊 Manacher's Algorithm

func longestPalindrome(s string) string {
	var (
		l        = len(s)
		maxLen   = 1
		maxStart = 0
		maxEnd   = 0
	)

	if len(s) == 1 {
		return s
	}

	for i := 1; i < l-maxLen/2+1; i++ {
		theLen, j1, j2 := 1, i-1, i+1
		for ; j1 >= 0 && j2 < l; {
			if s[j1] != s[j2] {
				break
			}
			theLen+=2
			j1--
			j2++
		}
		if theLen > maxLen {
			maxLen = theLen
			if j1 >= 0 && j2 < l && s[j1] == s[j2] {
				maxStart = j1
				maxEnd = j2
			} else {
				maxStart = j1 + 1
				maxEnd = j2 - 1
			}
		}

		theLen, j1, j2 = 0, i-1, i
		for ; j1 >= 0 && j2 < l; {
			if s[j1] != s[j2] {
				break
			}
			theLen+=2
			j1--
			j2++
		}
		if theLen > maxLen {
			maxLen = theLen
			if j1 >= 0 && j2 < l && s[j1] == s[j2] {
				maxStart = j1
				maxEnd = j2
			} else {
				maxStart = j1 + 1
				maxEnd = j2 - 1
			}
		}
	}

	return s[maxStart:maxEnd+1]
}
