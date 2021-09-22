package simple

import (
	"math/rand"
	"time"
)

// 进行性能测试的方法1：计算第 N 个菲波那切数
func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	sum1 := fib(n - 1)
	sum2 := fib(n - 2)
	return sum1 + sum2
}

// 进行性能测试的方法2：向数组添加元素（数组未设置初始容量大小）
func sliceCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

// 进行性能测试的方法3：向数组添加元素（数组设置了初始容量大小）
func sliceAppropriateCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}
