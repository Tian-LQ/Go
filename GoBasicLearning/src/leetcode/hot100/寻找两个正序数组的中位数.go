package hot100

import "math"

/*
 给定两个大小分别为m和n的正序（从小到大）数组nums1和nums2
 请你找出并返回这两个正序数组的 中位数
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	leftIndex := (l1 + l2 - 1) / 2
	rightIndex := (l1 + l2) / 2
	var leftNum, rightNum float64 = 0, 0
	for i, j, count := 0, 0, 0; count <= rightIndex; count++ {
		var val1, val2 int
		if i < l1 {
			val1 = nums1[i]
		} else {
			val1 = math.MaxInt
		}
		if j < l2 {
			val2 = nums2[j]
		} else {
			val2 = math.MaxInt
		}
		if val1 < val2 {
			i++
		} else {
			j++
		}
		if count == leftIndex {
			leftNum = min(val1, val2)
		}
		if count == rightIndex {
			rightNum = min(val1, val2)
		}
	}
	return (leftNum + rightNum) / 2
}

func min(val1, val2 int) float64 {
	if val1 < val2 {
		return float64(val1)
	} else {
		return float64(val2)
	}
}
