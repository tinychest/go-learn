package _89rotate_array

// Given an array, rotate the array to the right by k steps, where k is non-negative.
//
// Constraints:
//
//    1 <= nums.length <= 105
//    -231 <= nums[i] <= 231 - 1
//    0 <= k <= 105
//
// Follow up:
//
//    Try to come up with as many solutions as you can. There are at least three different ways to solve this problem.
//    Could you do it in-place with O(1) extra space?

func rotate(nums []int, k int) {
	// 1.跳过最沙雕的办法
	// 2.分析，应该能够做到一步移动到位的
	// - 超出数组长度的位移，直接取 %
	k = k % len(nums)
	// - 正好复原，就不移动了
	if k == 0 {
		return
	}
	// - 观察数值规律，发现规律，可以很简单知道每个位置的数最终移动的位置，发现想的太简单了，不是简单的移位，后续缝补就行 → ❌
	// - 标准做法是 O(n) 的直观做法，但是说做到 O(1)，那只能直接移动到目标位置，原数出来也直接找到目标位置，但是里边仍涉及到原来的数已经被换走了 → ❌
	// - 按照对称规律终于找到了一个

	l := len(nums)

	// 结尾向前的 k 个元素进行对称交换
	reverse(nums[l-k:])
	// 整个进行对称
	reverse(nums)
	// 第一步没有对称到的元素，最终被对称换到队尾，进行内部对称即可
	reverse(nums[k:])
}

func reverse(nums []int) {
	l := len(nums)

	// - 不使用 Go 的交换语法，能够减小内存分配
	var t int
	for i := 0; i < l/2; i++ {
		t = nums[i]
		nums[i] = nums[l-i-1]
		nums[l-i-1] = t
	}
}

// 感觉应该是有 O(n) 的做法，而不是当前的 O(2n)
