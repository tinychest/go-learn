package _reverse_integer

import (
	"strconv"
	"strings"
)

// 速度快（时间复杂度 O(log(x))），但是占用的空间大

// 做题时没有想到，%10 而是一直在想 strings.Reverse 实际根本没有这个方法

// 反省：不能一直错误提交啊，一定要好好看清楚题目

// 参考答案：
// 使用数值操作 % / 10，就一个注意点，就是整数超过能表示的最大最小数问题，贴出代码
// （7 和 -8 分别是 math.MaxInt32 和 math.MinInt32 的个位数值）
// if rev > math.MaxInt32/10 || (rev == math.MaxInt32 / 10 && pop > 7) return 0
// if rev < math.MinInt32/10 || (rev == math.MinInt32 / 10 && pop < -8) return 0

func reverse(x int) int {
	if x == 0 {
		return 0
	}

	s := strconv.Itoa(x)

	b := strings.Builder{}
	b.Grow(len(s))

	if strings.Contains(s, "-") {
		b.WriteRune('-')
		s = s[1:]
	}

	rs := []rune(s)
	for l := len(s) - 1; l >= 0; l-- {
		b.WriteRune(rs[l])
	}

	result, err := strconv.Atoi(b.String())
	if err != nil {
		panic(err)
	}
	if result < -2<<30 || result > 2<<30-1 {
		return 0
	}
	return result
}
