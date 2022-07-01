package sort

// 【希尔排序】
//
// 插入排序的最优时间复杂度居然可以达到 O(n)，插入排序达到这个时间复杂度，是数据基本有序的时候
// 所以，将目标数据划分成几个数据集，对每个数据集直接进行插入排序，这样可以让整个数据集趋向于有序，从而提高插入排序的效率；
//
// 注意，划分不是直接空间划分而是借助逻辑的间隔划分，划分成数据集的具体指（详见 shellSortPre）：
// 有一个间隙(gap)组序列，假如是 [3, 2, 1]（按照上面的描述，间隙组最后一个元素当然得是 1，且必不可少；之前的间隙，只是为最后一次间隙为 1 的插入排序做铺垫）
// 那么从 3 开始，那么就对目标数据进行 3 次指针跃进值为 3 的插入排序
// 随后是 2...再就是 1，也就是一个 插入排序
//
// [间隙组] 你会发现算法的效率肯定是和间隙组的有关的，应对不同的目标数据，肯定有不同的间隙组；
// 一般间隙组的第一个元素应当不小于目标数据个数的 1/3（最小值为 1），随后除以 3 得到下一个元素，以此类推，直到 1
//
// [是否稳定] false
//
// [空间复杂度] O(1)
//
// [时间复杂度]
// - 平均 O(nlogn)
// - 最优 O(nlog²n)
// - 最劣 O(nlog²n)

// shellSortPre
func shellSortPre(arr []int) {
	gaps := []int{4, 2, 1}

	for _, gap := range gaps {
		// 带有间隙的插入排序
		for i := gap; i < len(arr); i++ {
			for j := i; j-gap >= 0 && arr[j] < arr[j-gap]; j -= gap {
				arr[j], arr[j-gap] = arr[j-gap], arr[j]
			}
		}
	}
}

// ShellSort 应用稍好一些的插入排序 + 应用动态生成的间隙组规则
func ShellSort(arr []int) {
	length := len(arr)
	gap := 1
	for gap < length/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < length; i++ {
			temp := arr[i]

			// 前面有元素，大的向后
			j := i - gap
			for j >= 0 && arr[j] > temp {
				arr[j+gap] = arr[j]
				j -= gap
			}
			arr[j+gap] = temp
		}
		gap = gap / 3
	}
}
