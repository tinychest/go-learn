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

func minDistance(word1 string, word2 string) int {
	// 分析：这个问题应该是求两个串的最长公共子串
	//
	// 方向推测：应该是想考查 KMP 算法的变种
	//
	// 需求：KMP 算法的编写
}
