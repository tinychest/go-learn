package other

// StrMatch kmp（Knuth Morris Pratt） 是经典的不需要回溯高效串匹配算法
// 下面的方法将返回子串在目标串所有出现位置的下标
// - 构建 nextArr 数组 和 使用 nextArr 进行 kmp 串匹配的核心算法是一样
// - kmp 并不是在任何场景下都是高效，你会发现如果子串中的重复元素越多，每次发生不匹配时，就越有可能将子串中靠后的元素与主串下一个元素对齐进行新一轮的比较
//   如果对齐的是子串的第一个元素，相较于传统传比较，就会凸显出利用好之前匹配的过程，子串一下跳过了很多进行传统匹配的起始位置
//   换句话说，如果子串没有重复元素，那么当第 n + 1 个字符发生不匹配时，kmp 就能借助 nextArr 一下跳过更多的元素
func StrMatch(target, sub string) []int {
	// return basic(target, sub)
	return kmp(target, sub)
}

func basic(target, sub string) []int {
	res := make([]int, 0)
	l := len(sub)

	for i := 0; i < len(target); i++ {
		if len(target[i:]) < l {
			break
		}
		if target[i:i+l] == sub {
			res = append(res, i)
		}
	}
	return res
}

func kmp(target, sub string) []int {
	next := nextArr(sub)

	tl := len(target)
	sl := len(sub)

	res := make([]int, 0)
	j := -1
	for i := 0; i < tl; i++ {
		for j != -1 && target[i] != sub[j+1] {
			j = next[j]
		}
		if target[i] == sub[j+1] {
			j++
		}
		if j == sl-1 {
			res = append(res, i-j)
			j = next[j]
		}
	}
	return res
}

// sub：aaab
// next：0 0 1 2
func kmpExpect(target, sub string) []int {
	next := nextArr(sub)

	tl := len(target)
	sl := len(sub)

	res := make([]int, 0)
	j := 0
	for i := 0; i < tl; i++ {
		// 确定新一轮主串指针与子串的匹配
		for j != 0 && target[i] != sub[j] {
			j = next[j]
		}
		if target[i] == sub[j] {
			j++
		}
		if j == sl {
			res = append(res, i-j)
			j = next[j-1]
		}
	}
	return res
}
