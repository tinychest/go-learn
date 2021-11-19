package _4longest_common_prefix

// Constraints:
//
//    1 <= strs.length <= 200
//    0 <= strs[i].length <= 200
//    strs[i] consists of only lower-case English letters.

// 太简单了，没什么好说，时间 空间 100%
func longestCommonPrefix(strs []string) string {
	var (
		result string
		j      = 0
		c      uint8
	)
	for {
		for i := 0; i < len(strs); i++ {
			if len(strs[i]) == j {
				return result
			}
			if i == 0 {
				c = strs[i][j]
			}
			if strs[i][j] != c {
				return result
			}
		}
		result += string(c)
		j++
	}
}
