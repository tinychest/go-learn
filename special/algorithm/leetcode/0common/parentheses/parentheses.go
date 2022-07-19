package parentheses

// 参考 https://leetcode.com/problems/generate-parentheses/solution/
// 非常简单的判断单类型括号字符串是否匹配的方法
// 其实不用承担什么规律之类的思考负担，括号匹配的本质就是指定下标位置左边的 右括号数量 不能比 左括号数量 多
// 这也是在自己完成 32 题后明白的
func valid(s string) bool {
	balance := 0
	for _, c := range []byte(s) {
		if c == '(' {
			balance++
		} else {
			balance--
		}
		if balance < 0 {
			return false
		}
	}
	return balance == 0
}
