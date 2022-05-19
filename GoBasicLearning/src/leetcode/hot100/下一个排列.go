package hot100

/*
 整数数组的一个排列就是将其所有成员以序列或线性顺序排列。
*/

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	// 从后往前，先找到第一个升序对[i, i+1] (那么说明由[i+1, len-1]必定为降序)
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		// 从后往前找到第一个大于nums[i]的元素nums[j]
		for j >= 0 && nums[i] >= nums[j] {
			j--
		}
		// 交换nums[i]和nums[j]
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 此时[i+1, len-1]依然为降序
	// 反转区间[i+1, len-1]，使其变为升序
	reverse(nums[i+1:])
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
