package _0valid_parentheses

// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// An input string is valid if:
//
//    Open brackets must be closed by the same type of brackets.
//    Open brackets must be closed in the correct order.

// Constraints:
//
//    1 <= s.length <= 104
//    s consists of parentheses only '()[]{}'.
func isValid(s string) bool {
	return step3(s)
}

// [One Time Kill]
// Runtime: 3 ms, faster than 32.96% of Go online submissions for Valid Parentheses.
// Memory Usage: 1.9 MB, less than 86.41% of Go online submissions for Valid Parentheses.
// 执行性能确实不行，应该是自己对解题方法哪里理解的不够好（好像受到了早期解最少补充的括号数来使括号匹配的题思路的影响）
func step1(s string) bool {
	stack := make([]byte, len(s))

	pos := 0

	var p1, p2 int

	for _, v := range []byte(s) {
		stack[pos] = v
		pos++

		// 将栈中可以抵消的括号，进行抵消
		p1, p2 = pos-1, pos-2
		for p2 >= 0 {
			if stack[p1] == ')' && stack[p2] == '(' || stack[p1] == ']' && stack[p2] == '[' || stack[p1] == '}' && stack[p2] == '{' {
				p1 -= 2
				p2 -= 2
				pos -= 2
			} else {
				break
			}
		}
	}

	return pos == 0
}

// 参考：https://leetcode.com/problems/valid-parentheses/discuss/310836/Go
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
// Memory Usage: 2 MB, less than 86.41% of Go online submissions for Valid Parentheses.
func step2(s string) bool {
	stack := make([]byte, len(s))

	pos := 0

	for _, v := range []byte(s) {
		switch v {
		case ')', ']', '}':
			if pos == 0 {
				return false
			}
			if v == ')' && stack[pos-1] != '(' {
				return false
			}
			if v == ']' && stack[pos-1] != '[' {
				return false
			}
			if v == '}' && stack[pos-1] != '{' {
				return false
			}
			pos--
		case '(', '[', '{':
			stack[pos] = v
			pos++
		}
	}
	return pos == 0
}

// 将固定的空间改成自动扩容的形式（操作 stack 使用 append 的形式）
// 空间 99 了，时间又变长了
// 将为 stack 指定 cap，时间和空间都 有 88+ 了，可以接受
func step3(s string) bool {
	stack := make([]byte, 0)

	for _, v := range []byte(s) {
		switch v {
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			if v == ')' && stack[len(stack)-1] != '(' {
				return false
			}
			if v == ']' && stack[len(stack)-1] != '[' {
				return false
			}
			if v == '}' && stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '(', '[', '{':
			stack = append(stack, v)
		}
	}
	return len(stack) == 0
}
