package _7phone_number

import "bytes"

// Constraints:
//
//    0 <= digits.length <= 4
//    digits[i] is a digit in the range ['2', '9'].

// 漂亮 空间和时间都是 100%，0ms、2.1 MB

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	m := map[int][]string{
		2: {"a", "b", "c"},
		3: {"d", "e", "f"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"},
		8: {"t", "u", "v"},
		9: {"w", "x", "y", "z"},
	}

	ints := make([]int, 0, len(digits))
	for _, v := range digits {
		ints = append(ints, int(v)-48)
	}

	if len(digits) == 1 {
		return m[ints[0]]
	}

	s := make([][]string, 0, len(ints))
	for _, i := range ints {
		s = append(s, m[i])
	}
	return compose(s)
}

func compose(s [][]string) []string {
	var buffer bytes.Buffer
	buffer.Grow(len(s))

	res := make([]string, 0)
	poi := make([]int, len(s))
	for {
		// 记录路径结果
		for i, v := range s {
			buffer.WriteString(v[poi[i]])
		}
		res = append(res, buffer.String())
		buffer.Reset()

		// 变动指针
		j := len(s) - 1
		changed := false
		for j >= 0 {
			if poi[j] == len(s[j]) - 1 {
				poi[j] = 0
			} else {
				poi[j]++
				changed = true
				break
			}
			j--
		}
		if !changed {
			break
		}
	}
	return res
}
