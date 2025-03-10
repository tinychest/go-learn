package other

// 每个特定的字符串都可以构建得到一个特征数组，特征数组可以用于在进行串匹配，当发现不匹配时，子串需要回溯的低效问题
func nextArr(s string) []int {
	// - s 作为子串和主串进行匹配，如果第一个字符就不匹配，主串指针向后走一格，子串当然依旧从第一个元素进行比较，所以是 0
	// - 构建 nextArr 数组的时候，如果发生字符不匹配，并不是去找寻可以跳转的位置，相反，你需要去决定这个值
	next := make([]int, len(s))
	next[0] = -1

	j := -1
	for i := 1; i < len(s); i++ {
		for j != -1 && s[i] != s[j+1] {
			j = next[j]
		}
		if s[i] == s[j+1] {
			j++
		}
		next[i] = j
	}
	return next
}
