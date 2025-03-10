package _9divide_two_integers

import (
	"math"
)

// Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.
//
// The integer division should truncate toward zero, which means losing its fractional part. For example, 8.345 would be truncated to 8, and -2.7335 would be truncated to -2.
//
// Return the quotient after dividing dividend by divisor.
//
// Note: Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−231, 231 − 1]. For this problem,
//  if the quotient is strictly greater than 2^31 - 1, then return 2^31 - 1,
//  and if the quotient is strictly less than -2^31, then return -2^31.
//
// Example 1:
//
// Input: dividend = 10, divisor = 3
// Output: 3
// Explanation: 10/3 = 3.33333.. which is truncated to 3.
//
// Example 2:
//
// Input: dividend = 7, divisor = -3
// Output: -2
// Explanation: 7/-3 = -2.33333.. which is truncated to -2.
//
//
//
// Constraints:
//
//    -2^31 <= dividend, divisor <= 2^31 - 1
//    divisor != 0

// 参考思路有些是什么鬼，用 long 存储，还有直接用除法
// 但也是有收获的，就是肯定了下面通过加法翻倍提高效率、还有没必要就要转成正数，相关应该转成负数
// 除数翻倍 + 递归 就是最核心的解法
//
// 优化：除数翻倍的加法做法完全可以使用位移运算符去做
//
// 补错：注意题目要求 - 大于特定值以特定值返回
//
// 提交后优化：递归 可以 借助数组 优化内存空间

func divide(dividend int, divisor int) int {
	signOne, signTwo := dividend > 0, divisor > 0

	if signOne {
		dividend = 0 - dividend
	}
	if signTwo {
		divisor = 0 - divisor
	}

	res, _ := core(dividend, divisor, 1)
	if signOne != signTwo {
		res = 0 - res
	}
	if res > math.MaxInt32 {
		return math.MaxInt32
	}
	return res
}

// 因为除数和被除数都是负数，所以被除数小占大多数情况
func core(dividend, divisor, ratio int) (quotients, remainder int) {
	if divisor == -1 {
		return 0 - dividend, 0
	}

	var res int

	// 如果除数可以被两倍的除数除，就不除先进行除法，先进行除数翻倍的除法，将得到结果再进行除法（可以除的话）
	if dividend < (divisor << 1) {
		v1, v2 := core(dividend, divisor<<1, ratio<<1)
		res += v1
		dividend = v2
	}

	i := 0
	for i = 0; dividend <= divisor; i++ {
		dividend -= divisor
	}
	res += i * ratio
	return res, dividend
}
