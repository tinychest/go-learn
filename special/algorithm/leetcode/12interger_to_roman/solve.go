package _2interger_to_roman

// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
//
// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
//
// For example, 2 is written as II in Roman numeral, just two one's added together.
// 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
//
// Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV.
// Because the one is before the five we subtract it making four.
// The same principle applies to the number nine, which is written as IX.
// There are six instances where subtraction is used:
//
//    I can be placed before V (5) and X (10) to make 4 and 9.
//    X can be placed before L (50) and C (100) to make 40 and 90.
//    C can be placed before D (500) and M (1000) to make 400 and 900.
//
// Given an integer, convert it to a roman numeral.
//
//
//
// Example 1:
//
// Input: num = 3
// Output: "III"
// Explanation: 3 is represented as 3 ones.
//
// Example 2:
//
// Input: num = 58
// Output: "LVIII"
// Explanation: L = 50, V = 5, III = 3.
//
// Example 3:
//
// Input: num = 1994
// Output: "MCMXCIV"
// Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
//
//
//
// Constraints:
//
//    1 <= num <= 3999

// NOTE 优化方向
//  1.使用 strings.Builder
//  2. 可以直接把特殊情况直接纳入 ”进制单位“
func intToRoman(num int) string {
	var res string

	const ArrSum, SpecialSum = 7, 6

	// 所有能用来 “进制单位” 的数
	arr := [ArrSum]int{1000, 500, 100, 50, 10, 5, 1}
	// “进制单位” 数对应的字母
	arrMapping := [ArrSum]string{"M", "D", "C", "L", "X", "V", "I"}
	// 要考虑的特殊情况
	special := [SpecialSum]int{900, 400, 90, 40, 9, 4}
	// 特殊情况对应字母
	specialMapping := [SpecialSum]string{"CM", "CD", "XC", "XL", "IX", "IV"}

	var sumQuotient, numRemainder int
	for i := 0; i < ArrSum; i++ {
		sumQuotient = num / arr[i]
		numRemainder = num % arr[i]

		// 拼接叠加字符
		for j := 0; j < sumQuotient; j++ {
			res += arrMapping[i]
		}
		// 处理特殊情况（要处理的情况有限）
		if i != SpecialSum && numRemainder >= special[i] {
			res += specialMapping[i]
			numRemainder -= special[i]
		}

		num = numRemainder
	}

	return res
}
