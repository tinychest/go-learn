package _longest_substring_without_repeating_characters

// Given a string s, find the length of the longest substring without repeating characters.
//
//
// Constraints:
//
//    0 <= s.length <= 5 * 104
//    s consists of English letters, digits, symbols and spaces. 26 * 2 + 10 + ? + 1

func lengthOfLongestSubstring(s string) int {
	// - 这道题以时间为主，暂时先不考虑空间消耗
	// - 直接感觉是不是和 KMP 有关，但实际没有直接关系，但是想做到高效肯定是要借助 KMP 中那复用之前遍历过的结果
	// - 并且很容易也很自然的想到，使用 map 去做是否已经存在的操作

	m := make(map[byte]int, min(100, len(s)))

	var start int

	var res int
	for i, v := range []byte(s) {
		idx, ok := m[v]
		m[v] = i
		if !ok {
			res = max(res, len(m))
			continue
		}
		// 从发现元素重复的位置向后，这是这里实现了算法核心，达到了避免完全回溯的效果
		for j := start; j < idx; j++ {
			delete(m, s[j])
		}
		start = idx + 1
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
