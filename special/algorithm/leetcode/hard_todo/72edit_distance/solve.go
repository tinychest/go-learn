package _2edit_distance

// Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.
//
// You have the following three operations permitted on a word:
//
//    Insert a character
//    Delete a character
//    Replace a character
//
// Constraints:
//
//    0 <= word1.length, word2.length <= 500
//    word1 and word2 consist of lowercase English letters.

// 这道题真不愧为 hard 级别，不说用算法实现想法，这道题最直观的做法都不太好像，题目很抽象
func minDistance(word1 string, word2 string) int {
	res := 0
	l1, l2 := len(word1), len(word2)

	// 长度上的差异必定是结果中的一部分
	if l1 > l2 {
		res += l1 - l2
	} else {
		res += l2 - l1
	}

	// 观察字符构成

	return 0
}
