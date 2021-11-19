package _0regular_expression_matching

// Constraint
// 1 <= s.length <= 20
// 1 <= p.length <= 30
// s contains only lowercase English letters.
// p contains only lowercase English letters, '.', and '*'.
// It is guaranteed for each appearance of the character '*', there will be a previous valid character to match.


// 看清楚后，发现这个题的死办法就是遍历所有情况 → 然后，有类似优化算法可以极大的减少时间复杂度
// 做题的时候，没有这样想，但是可以往，使用指定的正则字符去得出指定路径的思路去想
func isMatch(s, p string) bool {
	return isMatchr([]rune(s), []rune(p))
}

// 没有考虑到 . 是必须表示不为空的字符的
// func isMatchr(s, p []rune) bool {
// 	if len(s) == 0 && len(p) == 0 {
// 		return true
// 	}
//
// 	firstMatch := safeIndex(s, 0) == safeIndex(p, 0) || safeIndex(p, 0) == '.'
// 	if firstMatch && isMatchr(safeRange(s, 1), safeRange(p, 1)) {
// 		return true
// 	} else if len(p) >= 2 && p[1] == '*' {
// 		return firstMatch && len(s) >= 1 && isMatchr(s[1:], p) || isMatchr(s, p[2:])
// 	}
//
// 	return false
// }

// 也就是说空串也是消耗的就行了
func isMatchr(s, p []rune) bool {
	if s == nil {
		return false
	}
	if len(s) == 0 && len(p) == 0 {
		return true
	}

	firstMatch := safeIndex(s, 0) == safeIndex(p, 0) || safeIndex(p, 0) == '.'
	if firstMatch && isMatchr(safeRange(s, 1), safeRange(p, 1)) {
		return true
	} else if len(p) >= 2 && p[1] == '*' {
		return firstMatch && len(s) >= 1 && isMatchr(s[1:], p) || isMatchr(s, p[2:])
	}

	return false
}

func safeIndex(rs []rune, i int) rune {
	if len(rs) == 0 {
		return 0
	}
	return rs[i]
}

func safeRange(rs []rune, start int) []rune {
	if len(rs) == 0 {
		return nil
	}
	return rs[start:]
}