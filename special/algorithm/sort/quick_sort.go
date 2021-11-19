package sort

/*
【快速排序】
定义一个数组的头指针和一个数组的尾指针，在整个数组中随便选一个数（一般取第一个），先用一个临时变量将其记录下来

假定挖出数组第一个元素，记作 a；挖出来的同时，意味着产生了一个坑位.
从尾指针开始，向前找一个比 a 小或者相等 的数，找到了就将该数挖出来填到坑位里边，这样，前面的坑填上了，全局还是有一个坑.

然后从头指针开始，向后找一个比挖出来（注意还是第一次挖出来的数） 大或相等 的数，同样，挖出来，填到全局的坑.

上面两个步骤中的找寻条件，除了找到了符合条件的数，还有两个指针相遇了，这时寻找也应该停下
当两个指针相遇后，此时两个指针指向的地方就是一个坑，此时将最开始被挖出的数填到这个地方
（最初被挖出来的数，在经过上面流程后，剩下的坑就是这个数最终的正确位置，理解起来简单也迷惑，最终坑位左边的数都比这个数小或者相等，最终坑位右边的数都比这个数大或者相等）
（在上面不停挖坑，不停填埋的过程中，头指针或尾指针始终是指向当前坑的位置的）

【总结】
上面的流程时快速排序算法的核心，你也会发现，算法执行完了一轮，就会确定一个数的最终位置
接下来我们复用上面的流程，将第一轮确定的数的左边所有的数拉去快速排序，右边的数也拉去快速排序

没错这就是经典的快速排序的递归形式，当然，光是这样并没有结束，递归总是讲究局部原理的重复，以及原理必须具备的完整性，即这里还缺少终止条件
稍微想想，也不难发现，当分组分到最后，也就是一组里面就一个元素的时候，那么此时就无法按照上面的流程重复了，因为左右两边都没有元素了，那么此时，就应该终止了

【脑洞】
左边挖一下，右边挖一下，好麻烦啊，不能一下把最终坑左边的数确定么？你可以一直找比初次挖出数小的，那么这样你就要遍历到底，并且遍历完，最终的是你并不知道临界点在哪。从这里边走出来，你会发现这个就是选择排序了...
所以这个一定是交替进行的
*/

// 从大到小排序
func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	casual := arr[0]
	start, end := 0, len(arr)-1

	for start < end {
		// 注意等号问题：两个内循环不能同时去掉等号，即，当头或者尾指针指向的元素和目标元素相等时，不应该跳出循环移动元素
		// 不是略微提高算法性能的问题 ，假如包含重复元素，会导致死循环，尾元素移到头，头元素移到尾
		for arr[end] <= casual && start < end {
			end--
		}
		// 如果找到了才应该做填
		if start < end {
			arr[start] = arr[end]
		}
		for arr[start] >= casual && start < end {
			start++
		}
		if start < end {
			arr[end] = arr[start]
		}
	}
	arr[start] = casual

	QuickSort(arr[:start])
	// NOTE 因为是先从尾指针向前回溯开始（而不是头指针向后），所以不会出现越界的情况
	QuickSort(arr[start+1:])
}

/*
【补充】
1.想把别的数作为第一个被挖坑的数呢，那需要做一下简单处理（将第一个数和你想指定的数换一个位置），否则会有问题
2.像上面的快速排序代码实现，完全可以将主逻辑体单独抽取到一个函数中，这样快速排序的递归逻辑就可以看的更清晰了
*/