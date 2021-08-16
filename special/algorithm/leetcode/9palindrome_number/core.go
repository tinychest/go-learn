package _palindrome_number

import "strconv"

// Ques：是否是回文数

// 提交后，表现并，但是就算法的时间复杂度来说，是没什么问题的
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	s := strconv.Itoa(x)
	var i, j int
	if len(s)%2 == 0 {
		i, j = len(s)/2-1, len(s)/2
	} else {
		i, j = len(s)/2-1, len(s)/2+1
	}
	for ; i >= 0 && j < len(s); {
		if s[i] != s[j] {
			return false
		}
		i--
		j++
	}
	return true
}

// 不用说一定要对称着去比较，直接取该数的回文数即可
func aisPalindrome2(x int) bool {
	if x < 0 {
		return false
	}

	var (
		reverted int
		temp = x
	)
	for ; temp > 0; {
		reverted = 10*reverted + temp%10
		temp /= 10
	}
	return reverted == x
}

// 还可以进一步优化：就是不用完整回文，只需要回一半，如果知道达到一半 - 原数和回文数的大小
func aisPalindrome3(x int) bool {
	if x < 0 || (x % 10 == 0 && x != 0) {
		return false
	}

	var revertedNumber = 0
	for ; x > revertedNumber; {
		revertedNumber = revertedNumber * 10 + x % 10
		x /= 10
	}

	// When the length is an odd number, we can get rid of the middle digit by revertedNumber/10
	// For example when the input is 12321, at the end of the while loop we get x = 12, revertedNumber = 123,
	// since the middle digit doesn't matter in palidrome(it will always equal to itself), we can simply get rid of it.
	return x == revertedNumber || x == revertedNumber/10
}