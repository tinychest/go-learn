package _9group_anagrams

import "sort"

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
//
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
//
//
// Constraints:
//
//    1 <= strs.length <= 104
//    0 <= strs[i].length <= 100
//    strs[i] consists of lowercase English letters.

// Runtime: 25 ms, faster than 88.53% of Go online submissions for Group Anagrams.
// Memory Usage: 8.4 MB, less than 51.70% of Go online submissions for Group Anagrams.
// [小结]
// 额，就是排个序，判断元素是否相同，并通过一个 map 来收集结果
// [参考]
// 用时最少的答案，只是把下面 map 的键类型改为 [26]int
func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string, 0)

	for _, v := range strs {
		bs := []byte(v)
		sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
		s := string(bs)
		m[s] = append(m[s], v)
	}
	res := make([][]string, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
