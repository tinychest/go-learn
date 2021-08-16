package _string_to_integer

import (
	"math"
	"strings"
)

// Ques：converts a string to a 32-bit signed integer

// 太没劲了，写完后简单扫了一下 strconv.Itoa 的源码，发现有个小亮点是小于 100 的数，就是常量字符串取字符

func myAtoi(s string) int {
	var (
		result   = 0
		add      int
		positive = true
	)

	s = strings.TrimLeft(s, " ")
	if len(s) == 0 {
		return 0
	}

	if s[0] == '-' {
		positive = false
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}

	for _, value := range s {
		if value < '0' || value > '9' {
			break
		}

		if positive {
			add = int(value) - '0'
			if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && add > 7) {
				return math.MaxInt32
			}
		} else {
			add = '0' - int(value)
			if result < math.MinInt32/10 || (result == math.MinInt32/10 && add < -8) {
				return math.MinInt32
			}
		}
		result *= 10
		result += add
	}

	return result
}
