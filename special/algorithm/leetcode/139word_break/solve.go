package _39word_break

import "strings"

// Given a string s and a dictionary of strings wordDict, return true if s can be segmented into a space-separated sequence of one or more dictionary words.
//
// Note that the same word in the dictionary may be reused multiple times in the segmentation.
//
//
// Constraints:
//
//    1 <= s.length <= 300
//    1 <= wordDict.length <= 1000
//    1 <= wordDict[i].length <= 20
//    s and wordDict[i] consist of only lowercase English letters.
//    All the strings of wordDict are unique.

func wordBreak(s string, wordDict []string) bool {
	// return step01(s, wordDict)

	// cannot := make(map[int]bool, 0)
	// return refer01(s, 0, wordDict, cannot)

	cannot := make([]bool, len(s))
	return refer02(s, 0, wordDict, cannot)
}

// 最直接的思路是说，将单词组按照首字母进行分组，然后，剩下的就直接按照匹配组匹配就是了
// - 自己也就一个暴力思路，暴力思路，不仅实现效率低，实现起来还费劲
// - 参考了一下，不应该思考从主串逐个字符遍历，而是，如果主串如果能由词组构成，则可以抵消完毕
//   怎么说呢，这种题见到就应该想到使用递归的方式去实现，目前这种题做得还是太少了
// （写这种字符串匹配相关的题，自己总是会被 KMP 影响，因为如果题目的解法是基于 KMP 做调整，那肯定是不好写的题目）
//
// [结果]
// 直接超时，超时样例当然是 a.....ab 和 [a, aa, aaa, aaaa, ......]
func step01(s string, wordDict []string) bool {
	if s == "" {
		return true
	}

	for _, v := range wordDict {
		if strings.HasPrefix(s, v) {
			return step01(s[len(v):], wordDict)
		}
	}
	return false
}

// 参考一下，优化点确实和 KMP 相似，就是想办法利用上之前过程的结果，有个名词叫 记忆递归（自己之前用过这个思路，优化汉诺塔问题）
// 我们分析上面非常耗时的样例，会发现当使用了很多个 a 进行匹配，得到不匹配的结果的结果后，
// 从末端开始，就开始使用 aa 进行匹配，最终又得到不匹配的结果，实际上会发现，上面核心的递归方法会重复很多次，从特定位置开始的核心算法过程。
// 应该将其记录下来
//
// [结果] 提交通过
func refer01(s string, index int, wordDict []string, cannot map[int]bool) bool {
	if index == len(s) {
		return true
	}
	if cannot[index] {
		return false
	}

	for _, v := range wordDict {
		if strings.HasPrefix(s[index:], v) {
			if refer01(s, index+len(v), wordDict, cannot) {
				return true
			} else {
				cannot[index] = true
			}
		}
	}

	return false
}

// 好像没必要用 map，所以改成了 []bool，0 ms 了
func refer02(s string, index int, wordDict []string, cannot []bool) bool {
	if index == len(s) {
		return true
	}
	if cannot[index] {
		return false
	}

	for _, v := range wordDict {
		if strings.HasPrefix(s[index:], v) {
			if refer02(s, index+len(v), wordDict, cannot) {
				return true
			} else {
				cannot[index] = true
			}
		}
	}

	return false
}
